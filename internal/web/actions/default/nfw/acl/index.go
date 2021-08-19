package acl

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	req_acl "github.com/1uLang/zhiannet-api/opnsense/request/acl"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/acl"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Keyword  string
	NodeId   uint64
	PageNum  int
	PageSize int
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		//this.ErrorPage(err)
		//return
		node = make([]*subassemblynode.Subassemblynode, 0)
	}
	// 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	list, err := acl.GetAclList(params.NodeId)
	if err != nil || list == nil {
		//this.ErrorPage(err)
		//return
		list = make([]*req_acl.AclListResp, 0)
	}

	this.Data["tableData"] = list
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId
	this.Show()
}
