package ddos

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Address  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {

	defer this.Show()

	this.Data["list"] = "[]"
	this.Data["total"] = 0
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
	list, total, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{
		NodeId:   params.NodeId,
		Addr:     params.Address,
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
	})
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取DDoS防火墙主机状态列表失败：%v", err.Error())
		return
	}
	this.Data["list"] = list
	this.Data["total"] = total
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId
	this.Data["address"] = params.Address
}
