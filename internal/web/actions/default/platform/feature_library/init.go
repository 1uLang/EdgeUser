package feature_library

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(NewHelper("virus")).
			Data("teaMenu", "virus").
			Prefix("/platform/feature_library/virus").
			Get("", new(IndexAction)).
			EndAll()
	})
}
