package feature

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	opnsense_server "github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/1uLang/zhiannet-api/opnsense/server/ips"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId uint64
}) {
	node, _, err := opnsense_server.GetOpnsenseNodeList()
	if err != nil || node == nil {
		node = make([]*subassemblynode.Subassemblynode, 0)
		//this.ErrorPage(err)
		//return
	}
	//nat 规则列表
	if params.NodeId == 0 && len(node) > 0 {
		params.NodeId = node[0].Id
	}
	version, err := ips.GetRuleInfo(&ips.IpsReq{
		NodeId: params.NodeId,
	})
	if err != nil || version == nil {
		//this.Show()
		//this.ErrorPage(err)
		//return
		this.Data["version"] = maps.Map{
			"update_time":  time.Now().Format("2006-01-02 15:04"),
			"version":      "版本",
			"all_total":    "总特征数",
			"update_total": "更新特征数",
			"name":         "ET open/emerging-scan",
		}
	} else {

		this.Data["version"] = maps.Map{
			"update_time":  version.UTime,
			"version":      version.Version,
			"all_total":    version.Total,
			"update_total": version.UTotal,
			"name":         version.Name,
		}
	}
	this.Data["nodes"] = node
	this.Data["selectNode"] = params.NodeId

	this.Show()

}
