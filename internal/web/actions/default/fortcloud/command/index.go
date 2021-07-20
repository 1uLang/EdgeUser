package command

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "command", "index")
}

func (this *IndexAction) RunGet() {
	this.Show()
}
