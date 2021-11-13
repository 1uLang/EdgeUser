package cert

import (
	"fmt"
	cert_model "github.com/1uLang/zhiannet-api/next-terminal/model/cert"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *UpdateAction) RunPost(params struct {
	Id       string
	Name     string
	Username string
	Password string

	Must *actions.Must
}) {

	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	params.Must.
		Field("username", params.Username).
		Require("请输入用户名")

	params.Must.
		Field("password", params.Password).
		Require("请输入密码")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}

	args := &cert_model.UpdateReq{}
	args.ID = params.Id
	args.Name = params.Name
	args.Username = params.Username
	args.Password = params.Password
	args.UserId = uint64(this.UserId())
	err = req.Cert.Update(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 修改凭证:[%v]成功", params.Id)
	this.Success()
}
