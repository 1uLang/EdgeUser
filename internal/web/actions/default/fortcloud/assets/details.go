package assets

import (
	"fmt"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type DetailsAction struct {
	actionutils.ParentAction
}

func (this *DetailsAction) Init() {
	this.Nav("", "fortcloud", "index")
}

func (this *DetailsAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *DetailsAction) RunPost(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &asset_model.DetailsReq{}
	args.Id = params.Id
	info, err := req.Assets.Details(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["asset"] = info
	// 日志
	this.CreateLogInfo("堡垒机 - 资产详情:[%v]成功", params.Id)
	this.Success()
}
