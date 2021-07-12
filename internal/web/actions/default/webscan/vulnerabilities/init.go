package vulnerabilities

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "webscan").
			Prefix("/webscan/vulnerabilities").
			Get("", new(IndexAction)).
			Get("/details", new(DetailAction)).
			EndAll()
	})
}
