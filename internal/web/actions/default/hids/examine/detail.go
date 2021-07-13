package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	MacCode string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	if err := hids.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := examine_server.Details(params.MacCode)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userName, err := this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取当前用户信息失败：%v", err))
		return
	}
	list, err := examine_server.List(&examine.SearchReq{UserName: userName, Type: -1, Score: -1, State: -1})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["details"] = info
	this.Data["otherDetails"] = list.ServerExamineResultInfoList[0]

	this.Show()
}