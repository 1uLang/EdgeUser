package admins

import (
	"fmt"
	admin_users_model "github.com/1uLang/zhiannet-api/jumpserver/model/admin_users"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
}
func (this *UpdateAction) RunPost(params struct {
	Id       string
	Name     string
	Username string
	Password string
	Comment  string

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请选择管理用户")

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

	args := &admin_users_model.UpdateReq{}
	args.ID = params.Id
	args.Name = params.Name
	args.UserName = params.Username
	args.Password = params.Password
	args.Comment = params.Comment
	args.UserId = uint64(this.UserId())
	_, err = req.AdminUser.Update(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 修改管理用户:[%v]成功", params.Id)
	this.Success()
}
