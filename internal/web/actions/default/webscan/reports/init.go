package reports

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//漏洞

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "reports").
			Prefix("/webscan/reports").
			GetPost("", new(IndexAction)).
			Post("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Get("/download", new(DownloadAction)).
			EndAll()
	})
}
