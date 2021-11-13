package cert

import (
	"fmt"
	cert_model "github.com/1uLang/zhiannet-api/next-terminal/model/cert"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}
func (this *IndexAction) RunGet() {

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.Data["nodeErr"] = "堡垒机组件错误:" + err.Error()
		this.Show()
		return
	}

	list, _, err := req.Cert.List(&cert_model.ListReq{
		UserId: uint64(this.UserId()),
	})
	this.Data["certs"] = list
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
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

	args := &cert_model.CreateReq{}
	args.Name = params.Name
	args.Username = params.Username
	args.Password = params.Password
	args.UserId = uint64(this.UserId())
	err = req.Cert.Create(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 新增授权凭证:[%v]成功", params.Name)
	this.Success()
}
