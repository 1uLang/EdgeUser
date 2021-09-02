package bwlist

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "bwlist").
			Prefix("/hids/bwlist").
			Get("", new(IndexAction)).
			GetPost("/del", new(DelAction)). //添删除ip
			GetPost("/createPopup", new(CreatePopupAction)).
			EndAll()
	})
}
