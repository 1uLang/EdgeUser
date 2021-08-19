package ddos

import (
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type ShieldAction struct {
	actionutils.ParentAction
}

func (this *ShieldAction) Init() {
	this.Nav("", "", "")
}

func (this *ShieldAction) RunGet(params struct {
	Address   string
	NodeId uint64
}) {
	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	if len(ddos) == 0 {
		this.Data["errorMessage"] = "未配置DDoS防火墙节点"
		return
	}
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}
	var all bool
	if len(params.Address) == 0 {
		all = true
	}
	list := &host_status.StatusFblink{}
	if all {
		hosts, _, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
			NodeId:   params.NodeId,
			Addr:     params.Address,
			PageSize: 999,
			PageNum:  1,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		for _,host := range hosts {
			hostShield, err := host_status_server.GetHostShieldList(&host_status_server.ShieldReq{
				Addr:   host.Addr,
				NodeId: params.NodeId,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			list.Fblink = append(list.Fblink, hostShield.Fblink...)
		}
	}else{
		list, err = host_status_server.GetHostShieldList(&host_status_server.ShieldReq{
			Addr:   params.Address,
			NodeId: params.NodeId,
		})
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list.Fblink
	this.Data["total"] = len(list.Fblink)
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address

	this.Show()
}
