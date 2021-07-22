package app

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Description string

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	//params.Must.
	//	Field("description", params.Description).
	//	Require("请输入备注")

	//accessKeyIdResp, err := this.RPC().UserAccessKeyRPC().CreateUserAccessKey(this.UserContext(), &pb.CreateUserAccessKeyRequest{
	//	UserId:      this.UserId(),
	//	Description: params.Description,
	//})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	//
	//defer this.CreateLogInfo("创建AccessKey %d", accessKeyIdResp.UserAccessKeyId)

	this.Success()
}
