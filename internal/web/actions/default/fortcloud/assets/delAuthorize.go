package assets

import (
	"fmt"
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type DelAuthorizeAction struct {
	actionutils.ParentAction
}

func (this *DelAuthorizeAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
}
func (this *DelAuthorizeAction) RunPost(params struct {
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
	err = req.Assets.DelAuthorize(&assets_model.DelAuthorizeReq{
		Asset: params.Id,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 删除授权资产:%v成功", params.Id)

	this.Success()
}
