package servers

import "github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	if !this.ValidateFeature("server.waf") {
		return
	}

	this.Show()
}
