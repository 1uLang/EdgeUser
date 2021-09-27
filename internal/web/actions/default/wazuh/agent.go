// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type AgentsAction struct {
	actionutils.ParentAction
}

func (this *AgentsAction) Init() {
	this.Nav("", "", "agents")
}

func (this *AgentsAction) RunGet(params struct{}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	list, err := server.AgentList(&agents.ListReq{
		UserId: this.UserId(true),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	active := 0
	for _, v := range list.AffectedItems {
		if v.Status == "active" {
			active++
		}
	}
	this.Data["agents"] = list.AffectedItems
	this.Data["dashboard"] = maps.Map{
		"total":   list.TotalAffectedItems,
		"active":  active,
		"offline": list.TotalAffectedItems - int64(active),
	}

	this.Data["agents"] = list.AffectedItems
	this.Show()
}

func (this *AgentsAction) RunPost(params struct {
	Agent string
}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = server.AgentDelete([]string{params.Agent})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.CreateLogInfo("主机防护 - 删除资产:[%v]成功", params.Agent)
	this.Success()
}
