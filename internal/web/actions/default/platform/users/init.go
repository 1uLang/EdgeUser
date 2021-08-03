package users

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/manage/users").
			Get("", new(IndexAction)).
			GetPost("/user", new(UserAction)).
			GetPost("/create", new(CreateAction)).
			GetPost("/features", new(FeaturesAction)).
			EndAll()
	})
}
