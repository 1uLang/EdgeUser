package stat

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/serverutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(serverutils.NewServerHelper()).
			Prefix("/servers/server/stat").
			Get("", new(IndexAction)).
			Get("/providers", new(ProvidersAction)).
			Get("/clients", new(ClientsAction)).
			Get("/waf", new(WafAction)).
			EndAll()
	})
}
