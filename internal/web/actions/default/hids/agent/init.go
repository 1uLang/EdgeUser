package agent

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//agent管理

func init() {
	_ = hids.InitAPIServer()
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "agent").
			Prefix("/hids/agent").
			GetPost("", new(IndexAction)).
			Get("/download", new(DownloadAction)).
			Get("/install", new(InstallAction)).
			Post("/disport", new(DisportAction)).
			EndAll()
	})
}
