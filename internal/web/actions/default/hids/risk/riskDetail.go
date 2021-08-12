package risk

import (
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type RiskDetailAction struct {
	actionutils.ParentAction
}

func (this *RiskDetailAction) Init() {
	this.FirstMenu("index")
}

// 系统漏洞列表
func (this *RiskDetailAction) RunGet(params struct {
	MacCode    string
	RiskId     string
	Os         string
	DetailName string
	State      int //是否已处理漏洞
	Must       *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	params.Must.
		Field("riskId", params.RiskId).
		Require("请输入系统漏洞id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	list, err := risk_server.SystemRiskDetail(params.MacCode, params.RiskId, params.State == 2)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["riskDetails"] = list
	this.Data["detailName"] = params.DetailName
	this.Data["os"] = params.Os

	this.Show()
}
