package login

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	userResp, err := this.RPC().UserRPC().FindEnabledUser(this.UserContext(), &pb.FindEnabledUserRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	user := userResp.User
	if user == nil {
		this.NotFound("user", this.UserId())
		return
	}
	this.Data["user"] = maps.Map{
		"username": user.Username,
		"fullname": user.Fullname,
	}

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Username  string
	Password  string
	Password2 string

	Must *actions.Must
}) {
	defer this.CreateLogInfo("修改登录设置")

	params.Must.
		Field("username", params.Username).
		Require("请输入登录用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能包含英文、数字或下划线")

	existsResp, err := this.RPC().UserRPC().CheckUserUsername(this.UserContext(), &pb.CheckUserUsernameRequest{
		UserId:   this.UserId(),
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if existsResp.Exists {
		this.FailField("username", "此用户名已经被别的用户使用，请换一个")
	}

	if len(params.Password) > 0 {
		if params.Password != params.Password2 {
			this.FailField("password2", "两次输入的密码不一致")
		}
	}

	_, err = this.RPC().UserRPC().UpdateUserLogin(this.UserContext(), &pb.UpdateUserLoginRequest{
		UserId:   this.UserId(),
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
