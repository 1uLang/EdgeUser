package users

import "github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct{}) {

	this.Show()
}
