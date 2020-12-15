package bills

import "github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	// 暂时先跳转到账单页面
	this.RedirectURL("/finance/bills")
}
