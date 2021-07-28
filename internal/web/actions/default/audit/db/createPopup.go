package db

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_db"
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
	Name    string
	Type    int
	Version string
	Ip      string
	Port    string
	System  uint
	Status  uint
	Id      uint64
	Edit    bool
}) {
	if !params.Edit && params.Type == 0 {
		params.Type = -1
	}
	if params.Version == "" {
		params.Version = "-1"
	}
	if !params.Edit {
		params.System = 1
		params.Status = 1
	}

	this.Data["name"] = params.Name
	this.Data["typeSelect"] = params.Type
	this.Data["verSelect"] = params.Version
	this.Data["ip"] = params.Ip
	this.Data["port"] = params.Port
	this.Data["systemSelect"] = params.System
	this.Data["openState"] = params.Status
	this.Data["id"] = params.Id
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name    string
	Type    uint
	Version string
	Ip      string
	Port    string
	System  uint
	Status  uint
	Id      uint64

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")
	params.Must.
		Field("ip", params.Ip).
		Require("请输入ip")
	params.Must.
		Field("port", params.Port).
		Require("请输入port")
	if params.Id == 0 {
		res, err := audit_db.AddDb(&audit_db.DBReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:    params.Name,
			Type:    params.Type,
			Version: params.Version,
			IP:      params.Ip,
			Port:    params.Port,
			System:  params.System,
			Status:  params.Status,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		defer this.CreateLogInfo("创建安全审计-数据库 %v", res.Msg)
	} else {
		res, err := audit_db.EditDb(&audit_db.DBEditReq{
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

		defer this.CreateLogInfo("修改安全审计-数据库 %v", res.Msg)
	}

	this.Success()
}
