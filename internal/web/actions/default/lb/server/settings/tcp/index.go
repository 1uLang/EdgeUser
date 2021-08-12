package tcp

import (
	"encoding/json"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/types"
)

// TCP设置
type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("tcp")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	this.Data["canSpecifyPort"] = this.ValidateFeature("lb-tcp.port")

	server, err := dao.SharedServerDAO.FindEnabledServer(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if server == nil {
		this.NotFound("server", params.ServerId)
		return
	}
	tcpConfig := &serverconfigs.TCPProtocolConfig{}
	if len(server.TcpJSON) > 0 {
		err := json.Unmarshal(server.TcpJSON, tcpConfig)
		if err != nil {
			this.ErrorPage(err)
		}
	} else {
		tcpConfig.IsOn = true
	}

	this.Data["serverType"] = server.Type
	this.Data["tcpConfig"] = tcpConfig

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId   int64
	ServerType string
	Addresses  string

	Must *actions.Must
}) {
	defer this.CreateLogInfo("修改代理服务 %d TCP设置", params.ServerId)

	canSpecifyPort := this.ValidateFeature("lb-tcp.port")

	server, err := dao.SharedServerDAO.FindEnabledServer(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if server == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	addresses := []*serverconfigs.NetworkAddressConfig{}
	err = json.Unmarshal([]byte(params.Addresses), &addresses)
	if err != nil {
		this.Fail("端口地址解析失败：" + err.Error())
	}

	// 检查端口是否被使用
	clusterIdResp, err := this.RPC().UserRPC().FindUserNodeClusterId(this.UserContext(), &pb.FindUserNodeClusterIdRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	clusterId := clusterIdResp.NodeClusterId
	if clusterId == 0 {
		this.Fail("当前用户没有指定集群，不能使用此服务")
	}
	for _, address := range addresses {
		port := types.Int32(address.PortRange)
		if port < 1024 || port > 65534 {
			this.Fail("'" + address.PortRange + "' 端口范围错误")
		}
		resp, err := this.RPC().NodeClusterRPC().CheckPortIsUsingInNodeCluster(this.UserContext(), &pb.CheckPortIsUsingInNodeClusterRequest{
			Port:            port,
			NodeClusterId:   clusterId,
			ExcludeServerId: params.ServerId,
			ExcludeProtocol: "tcp",
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if resp.IsUsing {
			this.Fail("端口 '" + fmt.Sprintf("%d", port) + "' 正在被别的服务或者同服务其他网络协议使用，请换一个")
		}
	}

	tcpConfig := &serverconfigs.TCPProtocolConfig{}
	if len(server.TcpJSON) > 0 {
		err := json.Unmarshal(server.TcpJSON, tcpConfig)
		if err != nil {
			this.ErrorPage(err)
		}
	} else {
		tcpConfig.IsOn = true
	}
	if canSpecifyPort {
		tcpConfig.Listen = addresses
	}

	configData, err := json.Marshal(tcpConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().ServerRPC().UpdateServerTCP(this.UserContext(), &pb.UpdateServerTCPRequest{
		ServerId: params.ServerId,
		TcpJSON:  configData,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
