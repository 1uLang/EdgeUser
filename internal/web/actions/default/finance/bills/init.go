package bills

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "finance").

			// 财务管理
			Prefix("/finance/bills").
			Get("", new(IndexAction)).

			EndAll()
	})
}
