package databackup

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "databackup").
			Prefix("/databackup").
			GetPost("", new(IndexAction)).
			Get("/create", new(CreateAction)).
			Post("/delete", new(DeleteAction)).
			Post("/dir", new(DirAction)).
			Get("/download", new(DownLoadAction)).
			GetPost("/rename", new(RenameAction)).
			EndAll()
	})
}
