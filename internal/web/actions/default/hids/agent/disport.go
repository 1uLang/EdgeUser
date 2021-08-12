package agent

import (
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DisportAction struct {
	actionutils.ParentAction
}

func (this *DisportAction) Init() {
	this.FirstMenu("index")
}

func (this *DisportAction) RunPost(params struct {
	MacCode string
	Opt     string

	Must *actions.Must
}) {
	params.Must.
		Field("MacCode", params.MacCode).
		Require("请输入机器码")

	params.Must.
		Field("Opt", params.Opt).
		Require("请选择操作方式")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	err = agent_server.Disport(params.MacCode, params.Opt)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
