package assets

import (
	"fmt"
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type AuthorizeAction struct {
	actionutils.ParentAction
}

func (this *AuthorizeAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
}
func (this *AuthorizeAction) RunPost(params struct {
	Id     string
	Emails string
	Must   *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	params.Must.
		Field("emails", params.Emails).
		Require("请输入授权用户邮箱")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	//判断邮箱绑定用户
	emails := strings.Split(params.Emails, "\n")

	err = req.Assets.Authorize(&assets_model.AuthorizeReq{
		Asset: params.Id,
		Emails: emails,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 授权资产:%v成功", params.Id)

	this.Success()
}
