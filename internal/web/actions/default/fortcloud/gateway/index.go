package gateway

import (
	"fmt"
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "fortcloud", "index")
}

func (this *IndexAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {

	err := fortcloud.InitAPIServer()
	if err != nil {
		return nil, err
	}

	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}
func (this *IndexAction) RunGet(params struct {
	PageSize  int
	PageNo    int
	PageState int
	Gateway   string

	Must *actions.Must
}) {

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	list, _, err := req.GateWay.List(&gateway_model.ListReq{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["gateways"] = list
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	HostName   string
	Ip         string
	Type       string
	Password   string
	Port       int
	Username   string
	PrivateKey string
	Passphrase string
	Must       *actions.Must
}) {

	params.Must.
		Field("hostName", params.HostName).
		Require("请输入主机名").
		Field("port", params.Port).
		Require("请输入端口号").
		Field("type", params.Type).
		Require("请选择账户类型").
		Field("ip", params.Ip).
		Require("请输入主机地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的主机地址")

	if params.Type == "password" {
		params.Must.
			Field("username", params.Username).
			Require("请输入授权账号").
			Field("password", params.Password).
			Require("请输入密码")
	} else {
		params.Must.
			Field("privateKey", params.PrivateKey).
			Require("请选择密钥")
	}
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &gateway_model.CreateReq{}
	args.Name = params.HostName
	args.IP = params.Ip
	args.AccountType = params.Type
	args.Password = params.Password
	args.Port = params.Port
	args.Username = params.Username
	args.PrivateKey = params.PrivateKey
	args.Passphrase = params.Passphrase
	args.UserId = uint64(this.UserId())
	err = req.GateWay.Create(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 新增网关:[%v]成功", params.HostName)
	this.Success()
}
