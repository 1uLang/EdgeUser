package net

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf/helper"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(helper.NewHelper("net")).
			Data("teaMenu", "apt").
			Prefix("/waf/net").
			Get("", new(NetAction)).
			EndAll()
	})
}
