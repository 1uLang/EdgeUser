package agent

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//审计-agent
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "agent").
			Prefix("/audit/agent").
			GetPost("", new(IndexAction)).
			// Post("/delete", new(DeleteAction)).
			Get("/download", new(DownLoadAction)).
			// GetPost("/create", new(CreateAction)).
			// Get("/create", new(CreatePopupAction)).
			//GetPost("/auth", new(AuthAction)).
			//GetPost("/delete", new(DeleteAction)).
			EndAll()
	})
}
