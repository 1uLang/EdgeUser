package logs

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/platform/logs").
			Get("", new(IndexAction)).
			Get("/exportExcel", new(ExportExcelAction)).
			EndAll()
	})
}
