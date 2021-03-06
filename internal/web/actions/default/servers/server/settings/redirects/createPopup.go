package redirects

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"net/url"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct {
}) {
	this.Data["statusList"] = serverconfigs.AllHTTPRedirectStatusList()

	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	BeforeURL string
	AfterURL  string
	Status    int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("beforeURL", params.BeforeURL).
		Require("请填写跳转前的URL")

	// 校验格式
	{
		u, err := url.Parse(params.BeforeURL)
		if err != nil {
			this.FailField("beforeURL", "请输入正确的跳转前URL")
		}
		if (u.Scheme != "http" && u.Scheme != "https") ||
			len(u.Host) == 0 {
			this.FailField("beforeURL", "请输入正确的跳转前URL")
		}

	}

	params.Must.
		Field("afterURL", params.AfterURL).
		Require("请填写跳转后URL")

	// 校验格式
	{
		u, err := url.Parse(params.AfterURL)
		if err != nil {
			this.FailField("afterURL", "请输入正确的跳转后URL")
		}
		if (u.Scheme != "http" && u.Scheme != "https") ||
			len(u.Host) == 0 {
			this.FailField("afterURL", "请输入正确的跳转后URL")
		}
	}

	params.Must.
		Field("status", params.Status).
		Gte(0, "请选择正确的跳转状态码")

	this.Data["redirect"] = maps.Map{
		"status":    params.Status,
		"beforeURL": params.BeforeURL,
		"afterURL":  params.AfterURL,
		"isOn":      true,
	}

	this.Success()
}
