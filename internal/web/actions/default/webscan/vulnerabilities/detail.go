package vulnerabilities

import (
	vulnerabilities_server "github.com/1uLang/zhiannet-api/awvs/server/vulnerabilities"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	VulId string

	Must *actions.Must
}) {
	params.Must.
		Field("macCode", params.VulId).
		Require("请输入漏洞id")

	if err := webscan.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}

	info, err := vulnerabilities_server.Details(params.VulId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["data"] = info

	this.Success()
}
