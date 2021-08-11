package ip_white

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/settingsutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(settingutils.NewHelper("ip_white")).
			Prefix("/platform/ip_white").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
