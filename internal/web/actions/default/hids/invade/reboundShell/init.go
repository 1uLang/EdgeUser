package reboundShell

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//入侵威胁

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "hids").
			Prefix("/hids/invade/reboundShell").
			GetPost("", new(IndexAction)).
			Get("/detailList", new(DetailListAction)). //反弹shell 详情列表
			Get("/detail", new(DetailAction)).         //反弹shell 详情

			EndAll()
	})
}
