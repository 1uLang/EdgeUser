package bwlist

import (
	"github.com/1uLang/zhiannet-api/hids/server/bwlist"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DelAction struct {
	actionutils.ParentAction
}

func (this *DelAction) Init() {
	this.Nav("", "", "")
}

func (this *DelAction) RunGet(params struct{}) {
	this.Show()
}

func (this *DelAction) RunPost(params struct {
	Id   uint64
	Must *actions.Must
}) {
	params.Must.
		Field("Addr", params.Id).
		Require("请选择黑白名单")

	err := bwlist.DeleteBWList(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
