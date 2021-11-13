package servers

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateOnAction) RunPost(params struct {
	ServerId int64
	IsOn     bool
}) {
	defer this.CreateLogInfo("修改服务 %d 启用状态", params.ServerId)

	_, err := this.RPC().ServerRPC().UpdateServerIsOn(this.UserContext(), &pb.UpdateServerIsOnRequest{
		ServerId: params.ServerId,
		IsOn:     params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
