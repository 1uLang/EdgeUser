package ddos

import (
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type LinkAction struct {
	actionutils.ParentAction
}

func (this *LinkAction) Init() {
	this.Nav("", "", "")
}

func (this *LinkAction) RunGet(params struct {
	Address  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	this.Data["ddos"] = "[]"
	this.Data["nodeId"] = ""
	this.Data["address"] = ""
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
	all := false
	if len(params.Address) == 0 {
		all = true
	}
	list := &host_status.StatusLink{}
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
		for _, host := range hosts {
			hostLink, err := host_status_server.GetLinkList(&host_status_server.LinkReq{
				NodeId: params.NodeId,
				Addr:   host.Addr,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			list.Link = append(list.Link, hostLink.Link...)
		}
	} else {
		list, err = host_status_server.GetLinkList(&host_status_server.LinkReq{
			NodeId: params.NodeId,
			Addr:   params.Address,
		})
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list.Link
	this.Data["total"] = len(list.Link)
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address
	this.Show()
}
