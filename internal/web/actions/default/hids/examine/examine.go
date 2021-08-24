package examine

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type ExamineAction struct {
	actionutils.ParentAction
}

func (this *ExamineAction) Init() {
	this.FirstMenu("index")
}

func (this *ExamineAction) RunGet(params struct {
	MacCode  string
	ServerIp string
	Must     *actions.Must
}) {
	params.Must.Field("macCode", params.MacCode).Require("请输入机器码").
		Field("serverIp", params.ServerIp).Require("请输入主机ip")

	this.Data["macCode"] = params.MacCode
	this.Data["serverIp"] = params.ServerIp
	this.Show()
}