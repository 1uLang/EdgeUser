// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"time"
)

type SysCheckAction struct {
	actionutils.ParentAction
}

func (this *SysCheckAction) Init() {
	this.Nav("", "", "virus")
}

func (this *SysCheckAction) RunGet(params struct {
	Agent string
	Event string
	Path  string
}) {
	this.Data["event"] = params.Event
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
	if params.Event == "" { //文件列表
		list, err := server.SysCheckList(agents.SysCheckListReq{
			Agent:  params.Agent,
			Limit:  1,
			Offset: 0,
			//Start: 1630982235, End: 1631068635,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		page := this.NewPage(list.TotalAffectedItems)
		this.Data["page"] = page.AsHTML()

		list, err = server.SysCheckList(agents.SysCheckListReq{
			Agent:  params.Agent,
			Limit:  page.Size,
			Offset: page.Offset,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		this.Data["syschecks"] = list.AffectedItems

		this.Data["agents"] = agent.AffectedItems

		this.Data["agent"] = params.Agent

		this.Show()
		return
	}

	list, err := server.SysCheckESList(agents.ESListReq{
		Agent:  params.Agent,
		Path:   params.Path,
		Limit:  1,
		Offset: 0,
		Start:  time.Now().AddDate(0, 0, -1).Unix(),
		End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	page := this.NewPage(int64(list.Total))
	this.Data["page"] = page.AsHTML()

	list, err = server.SysCheckESList(agents.ESListReq{
		Agent:  params.Agent,
		Path:   params.Path,
		Limit:  int(page.Size),
		Offset: int(page.Offset),
		Start:  time.Now().AddDate(0, 0, -1).Unix(),
		End:    time.Now().Unix(),
		//Start: 1630982235, End: 1631068635,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["syschecks"] = list.Hits

	this.Data["agents"] = agent.AffectedItems

	this.Data["agent"] = params.Agent

	this.Show()
}
