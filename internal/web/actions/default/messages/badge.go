package messages

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type BadgeAction struct {
	actionutils.ParentAction
}

func (this *BadgeAction) RunPost(params struct{}) {
	countResp, err := this.RPC().MessageRPC().CountUnreadMessages(this.UserContext(), &pb.CountUnreadMessagesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["count"] = countResp.Count

	this.Success()
}
