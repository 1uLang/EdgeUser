package gateway

import (
	"fmt"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *DeleteAction) RunPost(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择网关")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	err = req.GateWay.Delete(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 删除网关:[%v]成功", params.Id)
	this.Success()
}
