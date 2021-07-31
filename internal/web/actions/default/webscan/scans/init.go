package scans

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//漏洞

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "scans").
			Prefix("/webscan/scans").
			GetPost("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			Get("/vulnerabilities", new(VulnerabilitiesAction)).
			Post("/delete", new(DeleteAction)).
			Post("/stop", new(StopAction)).
			Get("/statistics", new(StatisticsAction)).
			Post("/resume", new(ResumeAction)).
			EndAll()
	})
}
