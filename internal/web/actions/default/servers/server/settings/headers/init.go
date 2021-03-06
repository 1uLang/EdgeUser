package headers

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
			Prefix("/servers/server/settings/headers").
			Get("", new(IndexAction)).
			GetPost("/createSetPopup", new(CreateSetPopupAction)).
			GetPost("/updateSetPopup", new(UpdateSetPopupAction)).
			GetPost("/createDeletePopup", new(CreateDeletePopupAction)).
			Post("/deleteDeletingHeader", new(DeleteDeletingHeaderAction)).
			Post("/delete", new(DeleteAction)).
			EndAll()
	})
}
