package bwlist

import (
	bwlist_model "github.com/1uLang/zhiannet-api/hids/model/bwlist"
	"github.com/1uLang/zhiannet-api/hids/server/bwlist"
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

func (this *CreatePopupAction) RunGet(params struct{}) {

	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Addr  string
	White string
	Must  *actions.Must
}) {
	params.Must.
		Field("Addr", params.Addr).
		Require("请输入ip地址")

	err := bwlist.AddBWList(&bwlist_model.HIDSBWList{UserId: uint64(this.UserId(true)), IP: params.Addr, CreateTime: time.Now().Unix(),Black: params.White =="黑名单"})
	if err != nil {
		this.Fail(err.Error())
		return
	}
	this.Success()
}
