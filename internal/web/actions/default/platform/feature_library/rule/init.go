package feature

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/feature_library"

	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(feature_library.NewHelper("rule")).
			Data("teaMenu", "virus").
			Prefix("/platform/feature_library/rule").
			Get("", new(IndexAction)).
			EndAll()
	})
}
