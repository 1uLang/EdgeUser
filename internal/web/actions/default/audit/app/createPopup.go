package app

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"github.com/1uLang/zhiannet-api/common/server/audit_assets_relation_server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"time"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct {
	Name   string
	Type   int
	Ip     string
	Status uint
	Id     uint64
	Edit   bool
}) {
	if !params.Edit && params.Type == 0 {
		params.Type = -1
	}
	if !params.Edit {
		params.Type = -1
		params.Status = 1
	}

	this.Data["name"] = params.Name
	this.Data["typeSelect"] = params.Type
	this.Data["ip"] = params.Ip
	this.Data["openState"] = params.Status
	this.Data["id"] = params.Id
	this.Data["edit"] = params.Edit
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name   string
	Type   int
	Ip     string
	Status uint
	Id     uint64

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")
	params.Must.
		Field("ip", params.Ip).
		Require("请输入ip").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的ip")
	params.Must.
		Field("type", params.Type).
		In([]int{0, 1}, "请选择类型")
	if params.Id == 0 {
		res, err := audit_app.AddApp(&audit_app.AppReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:    params.Name,
			AppType: uint(params.Type),
			IP:      params.Ip,
			Status:  params.Status,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if res != nil && res.Code != 0 {
			this.Fail(res.Msg)
			return
		}
		//关联账号审计ID
		if res != nil {
			_, err = audit_assets_relation_server.Add(&audit_assets_relation.AuditAssetsRelation{
				UserId:     uint64(this.UserId()),
				AssetsId:   res.Data.Id,
				AssetsType: 2,
				CreateTime: int(time.Now().Unix()),
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
		defer this.CreateLogInfo("创建安全审计-应用 %v", res.Msg)
	} else {
		res, err := audit_app.EditApp(&audit_app.AppEditReq{
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
		if res != nil && res.Code != 0 {
			this.Fail(res.Msg)
			return
		}
		defer this.CreateLogInfo("修改安全审计-应用 %v", res.Msg)
	}

	this.Success()
}
