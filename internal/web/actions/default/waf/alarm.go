package waf

import (
	req_ips "github.com/1uLang/zhiannet-api/opnsense/request/ips"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/ips"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type AlarmAction struct {

	actionutils.ParentAction
}

func (this *AlarmAction) Init() {
	this.Nav("", "", "")
}

func (this *AlarmAction) RunGet(params struct {
	NodeId   uint64
	Keyword  string
	FileId   string
	PageSize int
	Page     int
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		//node = make([]*subassemblynode.Subassemblynode, 0)
		this.ErrorPage(err)
		return
	}
	// 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}

	if params.PageSize == 0 {
		params.PageSize = 20
	}
	if params.Page == 0 {
		params.Page = 1
	}

	list, err := ips.GetIpsAlarmList(&ips.IpsAlarmReq{
		IpsReq: ips.IpsReq{
			NodeId:   params.NodeId,
			PageSize: params.PageSize,
			PageNum:  params.Page,
		},
		FileId: params.FileId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := list.Total
	page := this.NewPage(int64(count))
	this.Data["page"] = page.AsHTML()
	if list == nil {
		list = &req_ips.IpsAlarmListResp{}
	}
	if len(list.Rows) > 0 {
		this.Data["tableData"] = list.Rows
	} else {
		this.Data["tableData"] = make([]interface{}, 0)
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
}