package risk

import (
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type WeakDetailAction struct {
	actionutils.ParentAction
}

func (this *WeakDetailAction) Init() {
	this.FirstMenu("index")
}

// 弱口令详情
func (this *WeakDetailAction) RunGet(params struct {
	MacCode      string
	RiskId       string
	ProcessState int

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)

	}

	info, err := risk_server.WeakDetail(params.MacCode, params.RiskId, params.ProcessState == 2)
	if err != nil {
		this.ErrorPage(err)

	}
	this.Data["weakDetails"] = info

	this.Show()
}
