package index

import (
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/dlclark/regexp2"
	"github.com/iwind/TeaGo/actions"
	stringutil "github.com/iwind/TeaGo/utils/string"
)

type UpdatePwdAction struct {
	actionutils.ParentAction
}

// 首页（登录页）

var TokenSalt1 = stringutil.Rand(32)

func (this *UpdatePwdAction) RunGet(params struct {
	From string

	Auth *helpers.UserShouldAuth
}) {

	//fmt.Println("username = ", params.Auth.GetUpdatePwdToken())
	if params.Auth.GetUpdatePwdToken() <= 0 {
		this.RedirectURL("/")
		return
	}

	this.Show()
}

// RunPost 提交
func (this *UpdatePwdAction) RunPost(params struct {
	Password        string
	ConfirmPassword string

	Must *actions.Must
	Auth *helpers.UserShouldAuth
	CSRF *actionutils.CSRF
}) {
	//fmt.Println("username = ", params.Auth.GetUpdatePwdToken())
	userId := params.Auth.GetUpdatePwdToken()
	if userId <= 0 {
		this.FailField("refresh", "页面信息已过期，请刷新后重试")
	}

	if params.Password != params.ConfirmPassword {
		this.FailField("password", "两次输入的密码不一致")
	}
	if params.Password == stringutil.Md5("") {
		this.FailField("password", "请输入密码")
	}
	reg, err := regexp2.Compile(
		`^(?![A-z0-9]+$)(?=.[^%&',;=?$\x22])(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,30}$`, 0)
	if err != nil {
		this.FailField("pass1", "密码格式不正确")
	}
	if match, err := reg.FindStringMatch(params.ConfirmPassword); err != nil || match == nil {
		this.FailField("pass1", "密码格式不正确")
	}
	res, err := edge_users_server.UpdatePwd(uint64(userId), stringutil.Md5(params.Password))
	if err != nil || res == 0 {
		this.Fail("修改密码失败")
	}
	//清除session
	params.Auth.Logout()

	this.Success()
}
