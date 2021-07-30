package reports

import (
	"fmt"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//主机漏洞扫描 - 导出报表

type ExportAction struct {
	actionutils.ParentAction
}

func (this *ExportAction) RunPost(params struct {
	Id     string
	Format string
	Must   *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择指定生成报表的扫描目标")

	params.Must.
		Field("format", params.Format).
		Require("请选择指定生成报表的类型")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	ret,err := nessus_scans_server.Export(&nessus_scans_model.ExportReq{ID: params.Id,Format: params.Format})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["url"] = fmt.Sprintf("%s/tokens/%s/download",webscan.NessusServerUrl,ret.Token)
	// 日志
	this.CreateLogInfo("漏洞扫描 - 生成目标扫描报表:%v成功", params.Id)

	this.Success()
}
