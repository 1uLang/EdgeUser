package abnormalAccount

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
			Prefix("/hids/invade/abnormalAccount").
			GetPost("", new(IndexAction)).
			Get("/detailList", new(DetailListAction)). //异常账号 详情列表
			Get("/detail", new(DetailAction)).         //异常账号 详情

			EndAll()
	})
}
