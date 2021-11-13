package agent

import (
	agent_model "github.com/1uLang/zhiannet-api/hids/model/agent"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"

	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	Address string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("address", params.Address).
		Require("请输入目标地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的目标地址")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = agent_server.Create(&agent_model.CreateReq{AgentIp: params.Address,UserId: uint64(this.UserId(true))})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
