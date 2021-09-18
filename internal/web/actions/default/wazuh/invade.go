// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type InvadeAction struct {
	actionutils.ParentAction
}

func (this *InvadeAction) Init() {
	this.Nav("", "", "virus")
}

func (this *InvadeAction) RunGet(params struct {
	Agent string
}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	agent, err := server.AgentList(&agents.ListReq{
		UserId: this.UserId(true),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if len(agent.AffectedItems) == 0 {
		this.Data["errorMsg"] = "请先添加资产"
		this.Data["baselines"] = []map[string]string{}
		this.Data["agents"] = []map[string]string{}
		this.Show()
		return
	}
	if params.Agent == "" {
		params.Agent = agent.AffectedItems[0].ID
		//params.Agent = "007"
	}
	list, err := server.InvadeThreatESList(agents.ESListReq{
		Agent:  params.Agent,
		Limit:  1,
		Offset: 0,
		//Start:  time.Now().AddDate(0, 0, -1).Unix(),
		//End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	page := this.NewPage(int64(list.Total))
	this.Data["page"] = page.AsHTML()

	list, err = server.InvadeThreatESList(agents.ESListReq{
		Agent:  params.Agent,
		Limit:  int(page.Size),
		Offset: int(page.Offset),
		//Start:  time.Now().AddDate(0, 0, -1).Unix(),
		//End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["invades"] = list.Hits

	this.Data["agents"] = agent.AffectedItems

	this.Data["agent"] = params.Agent

	this.Show()
}
