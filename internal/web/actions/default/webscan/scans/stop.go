package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type StopAction struct {
	actionutils.ParentAction
}

func (this *StopAction) RunPost(params struct {
	ScanIds []string

	Must *actions.Must
}) {

	params.Must.
		Field("ScanIds", params.ScanIds).
		Require("请输入扫描id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, scanId := range params.ScanIds {
		err = scans_server.Abort(scanId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Success()
}
