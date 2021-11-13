package messages

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type ReadAllAction struct {
	actionutils.ParentAction
}

func (this *ReadAllAction) RunPost(params struct{}) {
	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "将所有消息置为已读")

	_, err := this.RPC().MessageRPC().UpdateAllMessagesRead(this.UserContext(), &pb.UpdateAllMessagesReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
