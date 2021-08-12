package agent

import (
	agent_model "github.com/1uLang/zhiannet-api/hids/model/agent"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) Init() {
	this.FirstMenu("index")
}

func (this *DeleteAction) RunPost(params struct {
	Id  uint64

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请选择agent主机")


	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = agent_server.Delete(&agent_model.DeleteReq{
		Id: params.Id,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
