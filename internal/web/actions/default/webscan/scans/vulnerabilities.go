package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type VulnerabilitiesAction struct {
	actionutils.ParentAction
}

func (this *VulnerabilitiesAction) RunGet(params struct {
	ScanId        string
	ScanSessionId string
	VulId         string

	Must *actions.Must
}) {

	hostscans := strings.HasSuffix(params.ScanId,"-host")
	params.Must.
		Field("scanId", params.ScanId).
		Require("请输入扫描id")
	params.Must.
		Field("scanSessionId", params.ScanSessionId).
		Require("请输入扫描会话id")

	if !hostscans{
		params.Must.
			Field("vulId", params.VulId).
			Require("请输入漏洞id")
	}

	if err := webscan.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}
	var vul_func func()(interface{},error)

	if !hostscans{
		vul_func = func( ) (interface{}, error) {
			req := &scans.VulnerabilitiesReq{
				ScanId:        params.ScanId,
				ScanSessionId: params.ScanSessionId,
				VulId:         params.VulId,
			}
			return scans_server.Vulnerabilities(req)
		}
	}else {
		vul_func = func() (interface{}, error) {
			req := &nessus_scans_model.VulnerabilitiesReq{
				ID: strings.TrimSuffix(params.ScanSessionId,"-host"),
				HistoryId: strings.TrimSuffix(params.ScanId,"-host")}
			return nessus_scans_server.Vulnerabilities(req)
		}
	}

	info,err := vul_func()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

	this.Success()
}
