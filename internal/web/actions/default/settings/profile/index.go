package profile

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
		"fullname": user.Fullname,
	}

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Fullname string

	Must *actions.Must
}) {
	defer this.CreateLogInfo("修改个人资料")

	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入你的姓名")

	_, err := this.RPC().UserRPC().UpdateUserInfo(this.UserContext(), &pb.UpdateUserInfoRequest{
		UserId:  this.UserId(),
		Fullname: params.Fullname,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
