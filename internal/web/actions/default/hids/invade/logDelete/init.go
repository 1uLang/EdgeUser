package logDelete

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
			Prefix("/hids/invade/logDelete").
			GetPost("", new(IndexAction)).
			Get("/detailList", new(DetailListAction)). //日志异常删除 详情列表
			Get("/detail", new(DetailAction)).         //日志异常删除 详情

			EndAll()
	})
}
