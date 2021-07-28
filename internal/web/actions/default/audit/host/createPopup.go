package host

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct {
	Name   string
	Ip     string
	System uint
	Status uint
	Id     uint64
	Edit   bool
}) {

	if !params.Edit {
		params.System = 1
		params.Status = 1
	}

	this.Data["name"] = params.Name
	this.Data["ip"] = params.Ip
	this.Data["systemSelect"] = params.System
	this.Data["openState"] = params.Status
	this.Data["id"] = params.Id
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name   string
	Ip     string
	Port   string
	System uint
	Status uint
	Id     uint64

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")
	params.Must.
		Field("ip", params.Ip).
		Require("请输入ip")
	if params.Id == 0 {
		res, err := audit_host.AddHost(&audit_host.HostReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:   params.Name,
			IP:     params.Ip,
			System: params.System,
			Status: params.Status,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		defer this.CreateLogInfo("创建安全审计-主机 %v", res.Msg)
	} else {
		res, err := audit_host.EditHost(&audit_host.HostEditReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:   params.Name,
			Id:     params.Id,
			Status: params.Status,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		defer this.CreateLogInfo("修改安全审计-主机 %v", res.Msg)
	}

	this.Success()
}
