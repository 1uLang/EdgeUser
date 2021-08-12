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

type StopAction struct {
	actionutils.ParentAction
}

func (this *StopAction) RunPost(params struct {
	ScanIds []string
	Type    int
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
	var stop_func func(string) error

	for _, scanId := range params.ScanIds {

		//主机漏洞扫描
		if strings.HasSuffix(scanId, "-host") {
			stop_func = func(id string) (err error) {
				return nessus_scans_server.Pause(&nessus_scans_model.PauseReq{ID: strings.TrimSuffix(id, "-host")})
			}
		} else {
			stop_func = scans_server.Abort
		}
		err = stop_func(scanId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Success()
}
