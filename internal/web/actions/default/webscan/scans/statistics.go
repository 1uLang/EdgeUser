package scans

import (
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//

type StatisticsAction struct {
	actionutils.ParentAction
}

func (this *StatisticsAction) Init() {
	this.FirstMenu("index")
}

func (this *StatisticsAction) RunGet(params struct {
	ScanId        string
	ScanSessionId string
	TargetId      string

	Must *actions.Must
}) {

	params.Must.
		Field("ScanId", params.ScanId).
		Require("请输入扫描id")

	params.Must.
		Field("ScanSessionId", params.ScanSessionId).
		Require("请输入扫描会话id")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	info, err := scans_server.Statistics(params.ScanId, params.ScanSessionId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["statistics"] = info
	var severity float64

	scanning_app, isExist := info["scanning_app"].(map[string]interface{})
	if isExist {
		wvs, isExist := scanning_app["wvs"].(map[string]interface{})
		if isExist {
			main, isExist := wvs["main"].(map[string]interface{})
			if isExist {
				vulns, isExist := main["vulns"]
				if isExist {
					for _, vul := range vulns.([]interface{}) {
						s, isExist := vul.(map[string]interface{})["severity"].(float64)
						if isExist && s > severity {
							severity = s
						}
					}
				}
			}
		}
	}
	this.Data["severity"] = severity
	this.Success()
}
