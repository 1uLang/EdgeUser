package lb

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "lb").
			Prefix("/lb").
			GetPost("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Post("/updateOn", new(UpdateOnAction)).
			Post("/delete", new(DeleteAction)).
			Get("/server", new(ServerAction)).
			EndAll()
	})
}
