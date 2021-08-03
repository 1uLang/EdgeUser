package user

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//平台管理-用户
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "platform").
			Prefix("/platform/user").
			GetPost("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			//GetPost("/delete", new(DeleteAction)).
			EndAll()
	})
}
