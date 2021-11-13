package ip_white

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {

	config, err := configloaders.LoadSecurityConfig(this.UserId(true))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if config.AllowIPs == nil {
		config.AllowIPs = []string{}
	}
	this.Data["config"] = config
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	AllowIPs []string

	Must *actions.Must
}) {
	//defer this.CreateLogInfo("修改登录设置")
	config, err := configloaders.LoadSecurityConfig(this.UserId(true))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	// 允许的IP
	if len(params.AllowIPs) > 0 {
		for _, ip := range params.AllowIPs {
			_, err := shared.ParseIPRange(ip)
			if err != nil {
				this.Fail("允许访问的IP '" + ip + "' 格式错误：" + err.Error())
			}
		}
		config.AllowIPs = params.AllowIPs
	} else {
		config.AllowIPs = []string{}
	}

	// 允许本地
	config.AllowLocal = true

	err = configloaders.UpdateSecurityConfig(config, this.UserId(true))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
