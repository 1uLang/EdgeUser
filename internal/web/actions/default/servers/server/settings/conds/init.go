package conds

import (
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/servers/server/settings/conds").
			GetPost("/addGroupPopup", new(AddGroupPopupAction)).
			GetPost("/addCondPopup", new(AddCondPopupAction)).
			EndAll()
	})
}
