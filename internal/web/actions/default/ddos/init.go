package ddos

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "ddos").
			Prefix("/ddos").
			Data("teaMenu", "host").
			Get("/host", new(HostAction)).//主机状态
			Data("teaMenu", "shield").
			Get("/shield", new(ShieldAction)). //连接监控
			Data("teaMenu", "link").
			Get("/link", new(LinkAction)). //屏蔽列表
			EndAll()
	})
}
