package command

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "command").
			Prefix("/fortcloud/command").
			Get("", new(IndexAction)).
			EndAll()
	})
}
