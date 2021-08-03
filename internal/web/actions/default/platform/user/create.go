package user

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Show()
}