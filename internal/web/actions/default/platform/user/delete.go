package user

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	UserId uint64
}) {
	defer this.CreateLogInfo("删除用户 %d", params.UserId)

	// TODO 关联组件的账号是否需要删除

	err := server.DeleteUser(&model.DeleteUserReq{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
