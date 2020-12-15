package messages

import "github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"

type BadgeAction struct {
	actionutils.ParentAction
}

func (this *BadgeAction) RunPost(params struct{}) {
	this.Data["count"] = 0

	this.Success()
}
