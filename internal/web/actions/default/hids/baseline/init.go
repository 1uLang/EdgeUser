package examine

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//合规基线

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "baseline").
			Prefix("/hids/baseline").
			Get("", new(IndexAction)).
			Get("/detail", new(DetailAction)).
			Post("/check", new(CheckAction)).
			Get("/template", new(TemplateAction)).
			Get("/template/detail", new(TemplateDetailAction)).
			EndAll()
	})
}
