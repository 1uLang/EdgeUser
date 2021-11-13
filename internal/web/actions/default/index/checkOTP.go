package index

import (
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

// 检查是否需要OTP
type CheckOTPAction struct {
	actionutils.ParentAction
}

func (this *CheckOTPAction) Init() {
	this.Nav("", "", "")
}

func (this *CheckOTPAction) RunPost(params struct {
	Username string

	Must *actions.Must
}) {
	info, err := edge_logins_server.GetOtpByName(params.Username)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["requireOTP"] = info
	this.Success()
}
