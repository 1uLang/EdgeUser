package report

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_from"
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
	Id uint64
}) {
	if params.Id > 0 {
		info, _ := audit_from.GetFrom(&audit_from.GetFromReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Id: params.Id,
		})
		this.Data["id"] = info.Data.Info.ID
		this.Data["name"] = info.Data.Info.Name
		this.Data["email"] = info.Data.Info.Email
		this.Data["format"] = info.Data.Info.Format
		this.Data["assets_id"] = info.Data.Info.AssetsID
		this.Data["assets_type"] = info.Data.Info.AssetsType
		this.Data["cycle"] = info.Data.Info.Cycle
		this.Data["cycle_day"] = info.Data.Info.CycleDay
		this.Data["send_time"] = info.Data.Info.SendTime
	} else {

		this.Data["id"] = 0
		this.Data["name"] = ""
		this.Data["email"] = ""
		this.Data["format"] = 1
		this.Data["assets_id"] = 0
		this.Data["assets_type"] = 1
		this.Data["cycle"] = 1
		this.Data["cycle_day"] = 1
		this.Data["send_time"] = "00:00"
	}

	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name       string
	Email      string
	Format     int
	AssetsType int
	AssetsId   uint64
	Cycle      int
	CycleDay   int
	SendTime   string
	Id         uint64

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")
	params.Must.
		Field("email", params.Email).
		Require("请输入邮箱")

	if params.Id == 0 {
		res, err := audit_from.AddFrom(&audit_from.FromReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:       params.Name,
			Email:      params.Email,
			Format:     params.Format,
			AssetsType: params.AssetsType,
			AssetsId:   params.AssetsId,
			Cycle:      params.Cycle,
			CycleDay:   params.CycleDay,
			SendTime:   params.SendTime,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		defer this.CreateLogInfo("编辑安全审计-报表 %v", res.Msg)
	} else {
		res, err := audit_from.EditFrom(&audit_from.FromReq{
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
			Name:       params.Name,
			Email:      params.Email,
			Format:     params.Format,
			AssetsType: params.AssetsType,
			AssetsId:   params.AssetsId,
			Cycle:      params.Cycle,
			CycleDay:   params.CycleDay,
			SendTime:   params.SendTime,
			Id:         params.Id,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		defer this.CreateLogInfo("创建安全审计-报表 %v", res.Msg)
	}
	this.Success()
}
