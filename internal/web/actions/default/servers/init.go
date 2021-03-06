package servers

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "servers").
			Prefix("/servers").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Get("/serverNamesPopup", new(ServerNamesPopupAction)).
			GetPost("/addServerNamePopup", new(AddServerNamePopupAction)).
			Post("/status", new(StatusAction)).
			Post("/delete", new(DeleteAction)).
			Post("/updateOn", new(UpdateOnAction)).
			Get("/server", new(ServerAction)).
			GetPost("/addPortPopup", new(AddPortPopupAction)).
			EndAll()
	})
}
