package scans

import (
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
)

//主机漏洞扫描 - 重新扫描
type ResumeAction struct {
	actionutils.ParentAction
}

func (this *ResumeAction) RunPost(params struct {
	TargetIds []string
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要扫描的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	for _, targetId := range params.TargetIds {

		err = nessus_scans_server.Resume(&nessus_scans_model.ResumeReq{ID: targetId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("主机漏洞扫描 - 重新扫描任务目标:%v成功", params.TargetIds)

	this.Success()
}
