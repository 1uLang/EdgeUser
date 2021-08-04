package user

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/platform/user").
			Get("", new(IndexAction)).
			Post("/update", new(UpdateAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/create", new(CreateAction)).
			GetPost("/features", new(FeaturesAction)).
			EndAll()
	})
}
