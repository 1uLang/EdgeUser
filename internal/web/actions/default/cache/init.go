package cache

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "cache").
			Prefix("/cache").
			GetPost("", new(IndexAction)).
			GetPost("/preheat", new(PreheatAction)).
			EndAll()
	})
}
