package host

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AccessKeyId int64
}) {
	//defer this.CreateLogInfo("删除AccessKey %d", params.AccessKeyId)
	//
	//_, err := this.RPC().UserAccessKeyRPC().DeleteUserAccessKey(this.UserContext(), &pb.DeleteUserAccessKeyRequest{UserAccessKeyId: params.AccessKeyId})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}

	this.Success()
}
