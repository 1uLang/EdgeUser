package assets

import (
	"fmt"
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
	"strconv"
	"strings"
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
	Id        string
	HostName  string
	Ip        string
	Platform  string
	Protocols []string
	Active    bool
	AdminUser string
	Comment   string
	PublicIp  string

	Proto 	  string
	Must *actions.Must
}) {

	params.Must.
		Field("hostName", params.HostName).
		Require("请输入主机名")

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	params.Must.
		Field("adminUser", params.AdminUser).
		Require("请输入认证账号")

	params.Must.
		Field("platform", params.Platform).
		Require("请选择系统平台")

	params.Must.
		Field("ip", params.Ip).
		Require("请输入主机地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的主机地址")


	checkProtocols := func() bool {
		if len(params.Protocols) == 0 {
			return true
		}

		for _, v := range params.Protocols {
			tmp := strings.Split(v, "/")
			///1-65535
			if len(tmp) == 1 {
				return true
			}else{
				port,err := strconv.Atoi(tmp[1])
				if err != nil || port < 1 || port >65535{
					return true
				}
			}
		}
		return false
	}
	if checkProtocols() {

		params.Must.
			Field("proto", params.Proto).
			Require("请设置协议组")
	}

	if params.PublicIp != "" {
		params.Must.
			Field("publicIp", params.PublicIp).
			Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的公网ip")
	}

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &assets_model.UpdateReq{}
	args.ID = params.Id
	args.HostName = params.HostName
	args.IP = params.Ip
	args.Platform = params.Platform
	args.Protocols = params.Protocols
	args.Active = params.Active
	args.AdminUser = params.AdminUser
	args.Comment = params.Comment
	args.PublicIp = params.PublicIp
	args.UserId = uint64(this.UserId())

	_, err = req.Assets.Update(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 日志
	this.CreateLogInfo("堡垒机 - 修改资产:[%v]成功", params.HostName)
	this.Success()
}
