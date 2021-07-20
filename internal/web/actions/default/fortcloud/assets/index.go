package assets

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "fortcloud", "index")
}

func (this *IndexAction) RunGet() {
	this.Show()
}
