package settingutils

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
	action.Data["teaMenu"] = "platform"

	// 标签栏
	tabbar := actionutils.NewTabbar()
	tabbar.Add("安全策略", "", "/platform/strategy", "", this.tab == "strategy")
	tabbar.Add("访问设置", "", "/platform/ip_white", "", this.tab == "ip_white")
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
