package db

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//审计-日志
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "logs").
			Prefix("/audit/logs").
			GetPost("", new(IndexAction)).
			GetPost("/export", new(ExportAction)).
			EndAll()
	})
}
