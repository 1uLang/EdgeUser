package log

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/serverutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(serverutils.NewServerHelper()).
			Prefix("/servers/server/log").
			GetPost("", new(IndexAction)).
			GetPost("/today", new(TodayAction)).
			GetPost("/history", new(HistoryAction)).
			Get("/viewPopup", new(ViewPopupAction)).
			EndAll()
	})
}
