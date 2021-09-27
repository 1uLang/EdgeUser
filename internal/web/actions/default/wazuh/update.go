// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "agents")
}

func (this *UpdateAction) RunGet(params struct {
	Agent  string
	Remake string
}) {

	this.Data["agent"] = params.Agent
	this.Data["remake"] = params.Remake
	this.Show()
}
func (this *UpdateAction) RunPost(params struct {
	Agent  string
	Remake string
}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = server.AgentUpdate(agents.UpdateReq{ID: params.Agent, Remake: params.Remake})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.CreateLogInfo("主机防护 - 修改资产:[%v]成功", params.Agent)
	this.Success()
}
