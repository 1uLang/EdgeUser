package waf

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/waf/ipadmin"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/serverutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(serverutils.NewServerHelper()).
			Prefix("/servers/server/settings/waf").
			GetPost("", new(IndexAction)).
			Get("/ipadmin/allowList", new(ipadmin.AllowListAction)).
			Get("/ipadmin/denyList", new(ipadmin.DenyListAction)).
			//GetPost("/ipadmin", new(ipadmin.IndexAction)).
			//GetPost("/ipadmin/provinces", new(ipadmin.ProvincesAction)).
			GetPost("/ipadmin/createIPPopup", new(ipadmin.CreateIPPopupAction)).
			GetPost("/ipadmin/updateIPPopup", new(ipadmin.UpdateIPPopupAction)).
			Post("/ipadmin/deleteIP", new(ipadmin.DeleteIPAction)).
			EndAll()
	})
}
