package conds

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/conds/condutils"
	"github.com/iwind/TeaGo/actions"
)

type AddCondPopupAction struct {
	actionutils.ParentAction
}

func (this *AddCondPopupAction) Init() {
}

func (this *AddCondPopupAction) RunGet(params struct{}) {
	this.Data["components"] = condutils.ReadAllAvailableCondTypes()

	this.Show()
}

func (this *AddCondPopupAction) RunPost(params struct {
	CondType string
	CondJSON []byte

	Must *actions.Must
}) {
	condConfig := &shared.HTTPRequestCond{}
	err := json.Unmarshal(params.CondJSON, condConfig)
	if err != nil {
		this.Fail("解析条件设置时发生了错误：" + err.Error())
	}
	condConfig.Type = params.CondType

	this.Data["cond"] = condConfig
	this.Success()
}
