package vulnerabilities

import (
	vulnerabilities_server "github.com/1uLang/zhiannet-api/awvs/server/vulnerabilities"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	VulId  string
	ScanId string
	ScanSessionId string
	Must   *actions.Must
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

	hostScan := strings.HasSuffix(params.ScanId, "-host")

	if err := webscan.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}

	var detail_func func() (interface{}, error)

	if !hostScan {
		detail_func = func() (interface{}, error) {
			return vulnerabilities_server.Details(params.VulId)
		}
	} else {
		detail_func = func() (interface{}, error) {
			req := &nessus_scans_model.PluginsReq{HistoryId: strings.TrimSuffix(params.ScanId, "-host"), VulId: params.VulId,
				ID: strings.TrimSuffix(params.ScanSessionId, "-host")}
			return nessus_scans_server.Plugins(req)
		}
	}

	info, err := detail_func()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

	this.Success()
}
