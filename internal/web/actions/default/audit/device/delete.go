package device

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_device"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	Id uint64
}) {
	res, err := audit_device.DelDevice(&audit_device.DelDeviceReq{
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
		Id: params.Id,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除安全审计-主机 %v", res.Msg)

	this.Success()
}
