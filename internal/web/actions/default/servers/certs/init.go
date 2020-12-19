package certs

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "certs").
			Prefix("/servers/certs").
			Get("", new(IndexAction)).
			Post("/count", new(CountAction)).
			Get("/selectPopup", new(SelectPopupAction)).
			GetPost("/uploadPopup", new(UploadPopupAction)).

			EndAll()
	})
}
