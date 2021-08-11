package app

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//审计-数据库
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "app").
			Prefix("/audit/app").
			GetPost("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/authorize", new(AuthorizeAction)).
			GetPost("/auth", new(AuthAction)).
			GetPost("/delete", new(DeleteAction)).
			EndAll()
	})
}
