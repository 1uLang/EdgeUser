package servers

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/utils/domainutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"net"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct{}) {
	clusterIdResp, err := this.RPC().UserRPC().FindUserNodeClusterId(this.UserContext(), &pb.FindUserNodeClusterIdRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["clusterId"] = clusterIdResp.NodeClusterId

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	DomainNames     []string
	Protocols       []string
	CertIdsJSON     []byte
	OriginsJSON     []byte
	RequestHostType int32
	RequestHost     string
	CacheCondsJSON  []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	// 检查用户所在集群
	clusterIdResp, err := this.RPC().UserRPC().FindUserNodeClusterId(this.UserContext(), &pb.FindUserNodeClusterIdRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	clusterId := clusterIdResp.NodeClusterId

	if len(params.DomainNames) == 0 {
		this.Fail("请添加要加速的域名")
	}

	serverNames := []*serverconfigs.ServerNameConfig{}
	for _, domainName := range params.DomainNames {
		if !domainutils.ValidateDomainFormat(domainName) {
			this.Fail("域名'" + domainName + "'输入错误")
		}
		serverNames = append(serverNames, &serverconfigs.ServerNameConfig{
			Name:     domainName,
			Type:     "",
			SubNames: nil,
		})
	}
	serverNamesJSON, err := json.Marshal(serverNames)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 协议
	if len(params.Protocols) == 0 {
		this.Fail("请选择至少一个协议")
	}

	httpConfig := &serverconfigs.HTTPProtocolConfig{}
	httpsConfig := &serverconfigs.HTTPSProtocolConfig{}

	// HTTP
	if lists.ContainsString(params.Protocols, "http") {
		httpConfig.IsOn = true
		httpConfig.Listen = []*serverconfigs.NetworkAddressConfig{
			{
				Protocol:  serverconfigs.ProtocolHTTP,
				Host:      "",
				PortRange: "80",
			},
		}
	}

	// HTTPS
	if lists.ContainsString(params.Protocols, "https") {
		httpsConfig.IsOn = true
		httpsConfig.Listen = []*serverconfigs.NetworkAddressConfig{
			{
				Protocol:  serverconfigs.ProtocolHTTPS,
				Host:      "",
				PortRange: "443",
			},
		}

		if len(params.CertIdsJSON) == 0 {
			this.Fail("请选择或者上传HTTPS证书")
		}
		certIds := []int64{}
		err := json.Unmarshal(params.CertIdsJSON, &certIds)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if len(certIds) == 0 {
			this.Fail("请选择或者上传HTTPS证书")
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
		httpsConfig.SSLPolicyRef = &sslconfigs.SSLPolicyRef{
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
		if !domainutils.ValidateDomainFormat(host) {
			this.Fail("源站地址'" + host + "'格式错误")
		}

		if scheme != "http" && scheme != "https" {
			this.Fail("错误的源站协议")
		}

		addrHost, addrPort, err := net.SplitHostPort(host)
		if err != nil {
			addrHost = host
			if scheme == "http" {
				addrPort = "80"
			} else if scheme == "https" {
				addrPort = "443"
			}
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
	if params.RequestHostType > 0 {
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxy(this.UserContext(), &pb.UpdateReverseProxyRequest{
			ReverseProxyId:  reverseProxyId,
			RequestHostType: params.RequestHostType,
			RequestHost:     params.RequestHost,
			RequestURI:      "",
			StripPrefix:     "",
			AutoFlush:       false,
		})
		if err != nil {
			this.ErrorPage(err)
		}
	}

	// 缓存设置
	cacheCondMaps := []maps.Map{}
	if len(params.CacheCondsJSON) == 0 {
		this.Fail("请添加缓存设置")
	}
	err = json.Unmarshal(params.CacheCondsJSON, &cacheCondMaps)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(cacheCondMaps) == 0 {
		this.Fail("请添加缓存设置")
	}

	cacheRefs := []*serverconfigs.HTTPCacheRef{}
	if len(cacheCondMaps) > 0 {
		for _, condMap := range cacheCondMaps {
			durationMap := condMap.GetMap("duration")
			if durationMap == nil {
				continue
			}

			duration := &shared.TimeDuration{
				Count: durationMap.GetInt64("count"),
				Unit:  durationMap.GetString("unit"),
			}

			value := ""
			param := ""
			operator := ""
			condType := condMap.GetString("type")
			switch condType {
			case "url-extension":
				param = "${requestPathExtension}"
				operator = shared.RequestCondOperatorIn
				result := []string{}
				v := condMap.GetString("value")
				if len(v) > 0 {
					err = json.Unmarshal([]byte(v), &result)
					if err != nil {
						this.ErrorPage(err)
						return
					}
					value = v
				}
			case "url-prefix":
				param = "${requestPath}"
				operator = shared.RequestCondOperatorHasPrefix
				value = condMap.GetString("value")
			default:
				continue
			}

			conds := &shared.HTTPRequestCondsConfig{
				IsOn:      true,
				Connector: "or",
				Groups: []*shared.HTTPRequestCondGroup{
					{
						IsOn:      true,
						Connector: "or",
						Conds: []*shared.HTTPRequestCond{
							{
								IsRequest: true,
								Type:      condType,
								Param:     param,
								Operator:  operator,
								Value:     value,
							},
						},
					},
				},
			}

			cacheRef := &serverconfigs.HTTPCacheRef{
				IsOn:          true,
				CachePolicyId: 0,
				Key:           "${scheme}://${host}${requestURI}",
				Life:          duration,
				Status:        []int{200},
				MaxSize: &shared.SizeCapacity{
					Count: 128,
					Unit:  shared.SizeCapacityUnitMB,
				},
				SkipResponseCacheControlValues: nil,
				SkipResponseSetCookie:          true,
				EnableRequestCachePragma:       false,
				Conds:                          conds,
				CachePolicy:                    nil,
			}
			cacheRefs = append(cacheRefs, cacheRef)
		}
	}

	cacheConfig := &serverconfigs.HTTPCacheConfig{
		IsPrior:   false,
		IsOn:      true,
		CacheRefs: cacheRefs,
	}
	cacheConfigJSON, err := cacheConfig.AsJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 开始保存
	httpJSON, err := httpConfig.AsJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	httpsJSON, err := httpsConfig.AsJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	createResp, err := this.RPC().ServerRPC().CreateServer(this.UserContext(), &pb.CreateServerRequest{
		UserId:           this.UserId(),
		AdminId:          0,
		Type:             serverconfigs.ServerTypeHTTPProxy,
		Name:             serverNames[0].Name,
		Description:      "",
		ServerNamesJON:   serverNamesJSON,
		HttpJSON:         httpJSON,
		HttpsJSON:        httpsJSON,
		TcpJSON:          nil,
		TlsJSON:          nil,
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

	defer this.CreateLogInfo("创建CDN加速服务 %d", serverId)

	// 保存缓存设置
	webIdResp, err := this.RPC().HTTPWebRPC().CreateHTTPWeb(this.UserContext(), &pb.CreateHTTPWebRequest{RootJSON: nil})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	webId := webIdResp.WebId
	_, err = this.RPC().ServerRPC().UpdateServerWeb(this.UserContext(), &pb.UpdateServerWebRequest{
		ServerId: serverId,
		WebId:    webId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebCache(this.UserContext(), &pb.UpdateHTTPWebCacheRequest{
		WebId:     webId,
		CacheJSON: cacheConfigJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
