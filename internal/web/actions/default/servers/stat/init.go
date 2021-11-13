package stat

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "stat").
			Prefix("/servers/stat").
			Get("", new(IndexAction)).
			EndAll()
	})
}
