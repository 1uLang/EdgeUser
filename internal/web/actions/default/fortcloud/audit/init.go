package audit

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "audit").
			Prefix("/fortcloud/audit").
			Get("", new(IndexAction)).
			Post("/replay", new(ReplayAction)).
			Post("/download", new(DownloadAction)).
			Post("/monitor", new(MonitorAction)).
			EndAll()
	})
}
