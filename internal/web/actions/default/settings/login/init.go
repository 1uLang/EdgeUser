package login

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/settingutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(settingutils.NewHelper("login")).
			Prefix("/settings/login").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
