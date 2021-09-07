package targets

import (
	"fmt"
	targets_model "github.com/1uLang/zhiannet-api/awvs/model/targets"
	"github.com/1uLang/zhiannet-api/awvs/server/targets"
	scan_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	"github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	Id   uint64
	Addr string
	Desc string
	Host bool
	Must *actions.Must
}) {
	params.Must.Field("id", params.Id).Require("请输入id")

	if params.Host {
		config, err := scans.GetConfig(fmt.Sprintf("%v", params.Id))
		if err != nil {
			this.ErrorPage(err)
			return
		}
		this.Data["username"] = config.Username
		this.Data["password"] = config.Password
		this.Data["port"] = config.Port
		this.Data["osSet"] = config.Os
		this.Data["type"] = 2
	} else {
		config, err := targets.GetConfig(params.Id)
		if err != nil {
			this.ErrorPage(err)
			return
		}

		this.Data["username"] = config.Username
		this.Data["password"] = config.Password
		this.Data["port"] = ""
		this.Data["osSet"] = ""
		this.Data["type"] = 1
	}
	this.Data["address"] = params.Addr
	this.Data["desc"] = params.Desc
	this.Data["id"] = params.Id

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	Id       uint64
	Username string
	Password string
	Port     int
	Os       int
	Host     bool
	Address  string
	Desc     string
	More     bool
	Must     *actions.Must
	CSRF     *actionutils.CSRF
}) {
	params.Must.Field("id", params.Id).Require("请输入id")

	if params.Host {

		req := &scan_model.AddReq{ID: fmt.Sprintf("%v", params.Id)}
		if params.More { //高级设置 身份登录

			params.Must.
				Field("username", params.Username).
				Require("请输入用户名").
				Field("password", params.Password).
				Require("请输入密码").
				Field("port", params.Port).
				Require("请输入端口")

			req.Username = params.Username
			req.Password = params.Password
			req.Port = params.Port
			req.Os = params.Os

		}

		req.Settings.Name = params.Address
		req.Settings.Text_targets = params.Address
		req.Settings.Description = params.Desc
		err := scans.Update(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	} else {

		req := &targets_model.SetConfigReq{Id: params.Id}
		if params.More { //高级设置 身份登录

			params.Must.
				Field("username", params.Username).
				Require("请输入用户名").
				Field("password", params.Password).
				Require("请输入密码")

			req.Username = params.Username
			req.Password = params.Password
		}
		err := targets.SetConfig(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Success()
}
