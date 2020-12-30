package accesskeys

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	AccessKeyId int64
	IsOn        bool
}) {
	defer this.CreateLogInfo("设置AccessKey %d 启用状态", params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().UpdateUserAccessKeyIsOn(this.UserContext(), &pb.UpdateUserAccessKeyIsOnRequest{
		UserAccessKeyId: params.AccessKeyId,
		IsOn:            params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
