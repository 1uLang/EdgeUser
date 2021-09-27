// 主机防护使用wazuh组件

package wazuh

import (
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type CheckAction struct {
	actionutils.ParentAction
}

func (this *CheckAction) Init() {
	this.Nav("", "", "agents")
}

func (this *CheckAction) RunPost(params struct {
	Agent string
}) {

	err := InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = server.AgentCheck(params.Agent)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.CreateLogInfo("主机防护 - 重新扫描:[%v]成功", params.Agent)
	this.Success()
}
