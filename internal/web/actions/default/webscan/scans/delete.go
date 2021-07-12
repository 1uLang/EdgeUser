package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
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
	for _, scanid := range params.ScanIds {
		err = scans_server.Delete(scanid)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 删除扫描任务目标:%v成功", params.ScanIds)
	this.Success()
}
