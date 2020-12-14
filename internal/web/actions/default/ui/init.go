package ui

import (
	"compress/gzip"
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/actions"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/ui").

			// 以下的需要压缩
			Helper(&actions.Gzip{Level: gzip.BestCompression}).
			Get("/components.js", new(ComponentsAction)).
			EndHelpers().

			// 以下需要登录
			Get("/download", new(DownloadAction)).

			EndAll()
	})
}
