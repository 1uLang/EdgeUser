package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
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
	params.Must.
		Field("vulId", params.VulId).
		Require("请输入漏洞id")
	params.Must.
		Field("scanId", params.ScanId).
		Require("请输入扫描id")
	params.Must.
		Field("scanSessionId", params.ScanSessionId).
		Require("请输入扫描会话id")

	if err := webscan.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}

	req := &scans.VulnerabilitiesReq{
		ScanId:        params.ScanId,
		ScanSessionId: params.ScanSessionId,
		VulId:         params.VulId,
	}
	info, err := scans_server.Vulnerabilities(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

	this.Success()
}
