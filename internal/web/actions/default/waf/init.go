package waf

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "waf").
			Prefix("/waf").
			Get("", new(IndexAction)).
			Data("teaMenu", "ddos").
			Get("/ddos", new(DdosAction)).
			Data("teaMenu", "alarm").
			Get("/alarm", new(AlarmAction)).
			Data("teaMenu", "logs").
			Get("/logs", new(LogsAction)).
			EndAll()
	})
}
