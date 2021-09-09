package hostlist

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Prefix("/hostlist").
			GetPost("", new(IndexAction)).
			//Prefix("/create").
			//GetPost("", new(CreatePopupAction)).
			EndAll()
	})
}
