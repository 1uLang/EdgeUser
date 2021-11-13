package accesskeys

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "acl").
			Prefix("/acl/accesskeys").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Post("/updateIsOn", new(UpdateIsOnAction)).
			EndAll()
	})
}
