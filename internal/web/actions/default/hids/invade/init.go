package invade

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//入侵威胁

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "invade").
			Prefix("/hids/invade").
			Get("", new(IndexAction)).
			EndAll()
	})
}
