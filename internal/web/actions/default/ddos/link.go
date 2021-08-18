package host

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type LinkAction struct {
	actionutils.ParentAction
}

func (this *LinkAction) Init() {
	this.Nav("", "", "")
}

func (this *LinkAction) RunGet(params struct {
	Addr     string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	if params.NodeId == 0 {
		this.ErrorPage(fmt.Errorf("请选择节点"))
		return
	}
	list, err := host_status_server.GetLinkList(&host_status_server.LinkReq{
		NodeId: params.NodeId,
		Addr:   params.Addr,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list.Link
	this.Data["total"] = list.Total
	this.Success()
}
