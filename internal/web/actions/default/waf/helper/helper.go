package helper

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type Helper struct {
	tab string
}

func NewHelper(tab string) *Helper {
	return &Helper{
		tab: tab,
	}
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	goNext = true

	action := actionPtr.Object()

	// 左侧菜单
	action.Data["teaMenu"] = "apt"

	// 标签栏
	tabbar := actionutils.NewTabbar()
	tabbar.Add("恶意IP拦截", "", "/waf/apt", "", this.tab == "apt")
	tabbar.Add("网络防病毒", "", "/waf/net", "", this.tab == "net")
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
