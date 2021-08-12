package cert

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "cert").
			Prefix("/fortcloud/cert").
			GetPost("", new(IndexAction)).
			Post("/update", new(UpdateAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/authorize",new(AuthorizeAction)).
			Post("/details",new(DetailsAction)).
			EndAll()
	})
}
