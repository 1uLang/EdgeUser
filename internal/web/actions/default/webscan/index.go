package webscan

import (
	dashboard_server "github.com/1uLang/zhiannet-api/awvs/server/dashboard"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "webscan", "index")
}

func (this *IndexAction) RunGet() {
	nodeErr := InitAPIServer()
	data := map[string]interface{}{
		"nodeErr":                 "",
		"scans_running_count":     0,
		"scans_waiting_count":     0,
		"scans_conducted_count":   0,
		"vuln_count":              map[string]interface{}{"low": 0, "med": 0, "high": 0},
		"targets_count":           0,
		"most_vulnerable_targets": []interface{}{},
		"top_vulnerabilities":     []interface{}{},
	}
	if nodeErr != nil {
		//this.ErrorPage(err)
		data["nodeErr"] = nodeErr
		this.Data["data"] = data
		this.Show()
		return
	}
	info, err := dashboard_server.MeState()
	if err != nil || info == nil {
		//this.ErrorPage(err)
		this.Data["data"] = data
		this.Show()
		return
	}
	info["nodeErr"] = ""
	this.Data["data"] = info
	// 日志
	this.CreateLogInfo("WEB漏洞扫描请求")
	this.Show()
}
