package feature_library

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
	action.Data["teaMenu"] = "virus"
	// 标签栏
	tabbar := actionutils.NewTabbar()
	tabbar.Add("病毒库", "", "/platform/feature_library/virus", "", this.tab == "virus")
	tabbar.Add("漏洞库", "", "/platform/feature_library/loopholes", "", this.tab == "loopholes")
	tabbar.Add("NIPS规则库", "", "/platform/feature_library/feature", "", this.tab == "feature")
	tabbar.Add("HIPS规则库", "", "/platform/feature_library/rule", "", this.tab == "rule")
	tabbar.Add("更新设置", "", "/platform/feature_library/update", "", this.tab == "update")
	actionutils.SetTabbar(actionPtr, tabbar)

	return
}
