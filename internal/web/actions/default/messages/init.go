package messages

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/messages").
			Get("", new(IndexAction)).
			GetPost("/readAll", new(ReadAllAction)).
			GetPost("/readPage", new(ReadPageAction)).
			Post("/badge", new(BadgeAction)).
			EndAll()
	})
}
