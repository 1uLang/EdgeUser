package lb

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ServerId int64
}) {
	defer this.CreateLogInfo("删除负载均衡服务 %d", params.ServerId)

	_, err := this.RPC().ServerRPC().DeleteServer(this.UserContext(), &pb.DeleteServerRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
