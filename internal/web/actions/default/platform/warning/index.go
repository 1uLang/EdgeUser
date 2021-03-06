package warning

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
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

	this.Data["email"] = user.Email
	this.Data["enable"] = user.Email != ""
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Enable bool
	Email  string

	Must *actions.Must
}) {

	if params.Enable {
		params.Must.
			Field("email", params.Email).
			Require("请输入告警邮箱")
	}
	_, err := this.RPC().UserRPC().UpdateUser(this.UserContext(), &pb.UpdateUserRequest{
		UserId: this.UserId(),
		Email:  params.Email,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
