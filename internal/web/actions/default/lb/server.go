package lb

import (
	"github.com/TeaOSLab/EdgeUser/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type ServerAction struct {
	actionutils.ParentAction
}

func (this *ServerAction) Init() {
	this.Nav("", "", "")
}

func (this *ServerAction) RunGet(params struct {
	ServerId int64
}) {
	// TODO 先跳转到设置页面，将来实现日志查看、统计看板等
	this.RedirectURL("/lb/server/settings/basic?serverId=" + numberutils.FormatInt64(params.ServerId))
}
