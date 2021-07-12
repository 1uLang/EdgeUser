package reports

import (
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ReportIds []string

	Must *actions.Must
}) {

	params.Must.
		Field("ReportIds", params.ReportIds).
		Require("请输入报表id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, reportId := range params.ReportIds {
		err = reports_server.Delete(reportId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 删除目标扫描报表:%v成功", params.ReportIds)
	this.Success()
}
