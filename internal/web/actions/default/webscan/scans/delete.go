package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ScanIds []string
	Ids     []string
	Must    *actions.Must
}) {

	params.Must.
		Field("ScanIds", params.ScanIds).
		Require("请输入扫描id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	for k, scanid := range params.ScanIds {

		//主机漏洞扫描
		if strings.HasSuffix(scanid, "-host") {
			err = nessus_scans_server.DelHistory(&nessus_scans_model.DelHistoryReq{
				ID:        strings.TrimSuffix(params.Ids[k], "-host"),
				HistoryId: strings.TrimSuffix(scanid, "-host"),
			})
		} else {
			err = scans_server.Delete(scanid)
		}
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("漏洞扫描 - 删除扫描任务目标:%v成功", params.ScanIds)
	this.Success()
}
