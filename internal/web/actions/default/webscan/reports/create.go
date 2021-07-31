package reports

import (
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	Ids    []string
	TarIds []string
	Must   *actions.Must
}) {

	params.Must.
		Field("id_list", params.Ids).
		Require("请选择指定生成报表的扫描目标")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	webscan_ids := []string{}
	for _,v := range params.Ids {
		if !strings.HasSuffix(v,"-host"){//去掉主机漏洞扫描
			webscan_ids = append(webscan_ids, v)
		}
	}

	req := &reports.CreateResp{
		Source: struct {
			IDS  []string `json:"id_list"`
			Type string   `json:"list_type"`
		}{IDS: webscan_ids, Type: "scans"},
		TemplateId:  "11111111-1111-1111-1111-111111111112", //快速
		UserId: uint64(this.UserId()),
	}
	_, err = reports_server.Create(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 生成目标扫描报表:%v成功", params.Ids)

	this.Success()
}
