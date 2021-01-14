package lb

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"net"
	"strconv"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct{}) {
	if !this.ValidateFeature("server.tcp") {
		return
	}

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Name        string
	Protocols   []string
	CertIdsJSON []byte
	OriginsJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	if !this.ValidateFeature("server.tcp") {
		this.Fail("你没有权限使用此功能")
	}

	params.Must.
		Field("name", params.Name).
		Require("请输入服务名称")

	// 协议
	if len(params.Protocols) == 0 {
		this.Fail("请选择至少一个协议")
	}

	// 检查用户所在集群
	clusterIdResp, err := this.RPC().UserRPC().FindUserNodeClusterId(this.UserContext(), &pb.FindUserNodeClusterIdRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	clusterId := clusterIdResp.NodeClusterId

	// 先加锁
	lockerKey := "create_tcp_server"
	lockResp, err := this.RPC().SysLockerRPC().SysLockerLock(this.UserContext(), &pb.SysLockerLockRequest{
		Key:            lockerKey,
		TimeoutSeconds: 30,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if !lockResp.Ok {
		this.Fail("操作繁忙，请稍后再试")
	}
	defer func() {
		_, err := this.RPC().SysLockerRPC().SysLockerUnlock(this.UserContext(), &pb.SysLockerUnlockRequest{Key: lockerKey})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}()

	tcpConfig := &serverconfigs.TCPProtocolConfig{}
	tlsConfig := &serverconfigs.TLSProtocolConfig{}

	// TCP
	ports := []int{}
	if lists.ContainsString(params.Protocols, "tcp") {
		// 获取随机端口
		portResp, err := this.RPC().NodeClusterRPC().FindFreePortInNodeCluster(this.UserContext(), &pb.FindFreePortInNodeClusterRequest{NodeClusterId: clusterId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		port := int(portResp.Port)
		ports = append(ports, port)

		tcpConfig.IsOn = true
		tcpConfig.Listen = []*serverconfigs.NetworkAddressConfig{
			{
				Protocol:  serverconfigs.ProtocolTCP,
				Host:      "",
				PortRange: strconv.Itoa(port),
			},
		}
	}

	// TLS
	if lists.ContainsString(params.Protocols, "tls") {
		var port int

		// 尝试N次
		for i := 0; i < 5; i++ {
			// 获取随机端口
			portResp, err := this.RPC().NodeClusterRPC().FindFreePortInNodeCluster(this.UserContext(), &pb.FindFreePortInNodeClusterRequest{NodeClusterId: clusterId})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			p := int(portResp.Port)
			if !lists.ContainsInt(ports, p) {
				port = p
				break
			}
		}
		if port == 0 {
			this.Fail("无法找到可用的端口，请稍后重试")
		}

		tlsConfig.IsOn = true
		tlsConfig.Listen = []*serverconfigs.NetworkAddressConfig{
			{
				Protocol:  serverconfigs.ProtocolTLS,
				Host:      "",
				PortRange: strconv.Itoa(port),
			},
		}

		if len(params.CertIdsJSON) == 0 {
			this.Fail("请选择或者上传TLS证书")
		}
		certIds := []int64{}
		err := json.Unmarshal(params.CertIdsJSON, &certIds)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if len(certIds) == 0 {
			this.Fail("请选择或者上传TLS证书")
		}

		certRefs := []*sslconfigs.SSLCertRef{}
		for _, certId := range certIds {
			certRefs = append(certRefs, &sslconfigs.SSLCertRef{
				IsOn:   true,
				CertId: certId,
			})
		}
		certRefsJSON, err := json.Marshal(certRefs)
		if err != nil {
			this.ErrorPage(err)
			return
		}

		// 创建策略
		sslPolicyIdResp, err := this.RPC().SSLPolicyRPC().CreateSSLPolicy(this.UserContext(), &pb.CreateSSLPolicyRequest{
			Http2Enabled:      false,
			MinVersion:        "TLS 1.1",
			SslCertsJSON:      certRefsJSON,
			HstsJSON:          nil,
			ClientAuthType:    0,
			ClientCACertsJSON: nil,
			CipherSuites:      nil,
			CipherSuitesIsOn:  false,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		tlsConfig.SSLPolicyRef = &sslconfigs.SSLPolicyRef{
			IsOn:        true,
			SSLPolicyId: sslPolicyIdResp.SslPolicyId,
		}
	}

	// 源站信息
	originMaps := []maps.Map{}
	if len(params.OriginsJSON) == 0 {
		this.Fail("请输入源站信息")
	}
	err = json.Unmarshal(params.OriginsJSON, &originMaps)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(originMaps) == 0 {
		this.Fail("请输入源站信息")
	}
	primaryOriginRefs := []*serverconfigs.OriginRef{}
	backupOriginRefs := []*serverconfigs.OriginRef{}
	for _, originMap := range originMaps {
		host := originMap.GetString("host")
		isPrimary := originMap.GetBool("isPrimary")
		scheme := originMap.GetString("scheme")

		if len(host) == 0 {
			this.Fail("源站地址不能为空")
		}
		addrHost, addrPort, err := net.SplitHostPort(host)
		if err != nil {
			this.Fail("源站地址'" + host + "'格式错误")
		}

		if scheme != "tcp" && scheme != "tls" {
			this.Fail("错误的源站协议")
		}

		originIdResp, err := this.RPC().OriginRPC().CreateOrigin(this.UserContext(), &pb.CreateOriginRequest{
			Name: "",
			Addr: &pb.NetworkAddress{
				Protocol:  scheme,
				Host:      addrHost,
				PortRange: addrPort,
			},
			Description: "",
			Weight:      10,
			IsOn:        true,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if isPrimary {
			primaryOriginRefs = append(primaryOriginRefs, &serverconfigs.OriginRef{
				IsOn:     true,
				OriginId: originIdResp.OriginId,
			})
		} else {
			backupOriginRefs = append(backupOriginRefs, &serverconfigs.OriginRef{
				IsOn:     true,
				OriginId: originIdResp.OriginId,
			})
		}
	}
	primaryOriginsJSON, err := json.Marshal(primaryOriginRefs)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	backupOriginsJSON, err := json.Marshal(backupOriginRefs)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	scheduling := &serverconfigs.SchedulingConfig{
		Code:    "random",
		Options: nil,
	}
	schedulingJSON, err := json.Marshal(scheduling)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 反向代理
	reverseProxyResp, err := this.RPC().ReverseProxyRPC().CreateReverseProxy(this.UserContext(), &pb.CreateReverseProxyRequest{
		SchedulingJSON:     schedulingJSON,
		PrimaryOriginsJSON: primaryOriginsJSON,
		BackupOriginsJSON:  backupOriginsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	reverseProxyId := reverseProxyResp.ReverseProxyId
	reverseProxyRef := &serverconfigs.ReverseProxyRef{
		IsPrior:        false,
		IsOn:           true,
		ReverseProxyId: reverseProxyId,
	}
	reverseProxyRefJSON, err := json.Marshal(reverseProxyRef)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 开始保存
	tcpJSON, err := tcpConfig.AsJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	tlsJSON, err := tlsConfig.AsJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	createResp, err := this.RPC().ServerRPC().CreateServer(this.UserContext(), &pb.CreateServerRequest{
		UserId:           this.UserId(),
		AdminId:          0,
		Type:             serverconfigs.ServerTypeTCPProxy,
		Name:             params.Name,
		Description:      "",
		ServerNamesJON:   []byte("[]"),
		HttpJSON:         nil,
		HttpsJSON:        nil,
		TcpJSON:          tcpJSON,
		TlsJSON:          tlsJSON,
		UnixJSON:         nil,
		UdpJSON:          nil,
		WebId:            0,
		ReverseProxyJSON: reverseProxyRefJSON,
		GroupIds:         nil,
		NodeClusterId:    clusterId,
		IncludeNodesJSON: nil,
		ExcludeNodesJSON: nil,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverId := createResp.ServerId

	defer this.CreateLogInfo("创建TCP负载均衡服务 %d", serverId)

	this.Success()
}
