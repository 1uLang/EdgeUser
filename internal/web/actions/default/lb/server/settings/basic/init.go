package basic

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/serverutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(serverutils.NewServerHelper()).
			Prefix("/lb/server/settings/basic").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
