package examine

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//主机体检
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "examine").
			Prefix("/hids/examine").
			GetPost("", new(IndexAction)).
			Get("/detail", new(DetailAction)).
			Post("/scans", new(ScanAction)).
			Get("/examine", new(ExamineAction)).
			EndAll()
	})
}
