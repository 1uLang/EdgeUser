package certs

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/certs/acme"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/certs/acme/users"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(new(Helper)).
			Data("teaMenu", "certs").
			Prefix("/servers/certs").
			Data("leftMenuItem", "cert").
			Get("", new(IndexAction)).
			Post("/count", new(CountAction)).
			Get("/selectPopup", new(SelectPopupAction)).
			GetPost("/uploadPopup", new(UploadPopupAction)).
			Get("/datajs", new(DatajsAction)).
			Get("/viewKey", new(ViewKeyAction)).
			Get("/viewCert", new(ViewCertAction)).
			Get("/downloadKey", new(DownloadKeyAction)).
			Get("/downloadCert", new(DownloadCertAction)).
			Get("/downloadZip", new(DownloadZipAction)).
			Get("/certPopup", new(CertPopupAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).

			// ACME
			Prefix("/servers/certs/acme").
			Data("leftMenuItem", "acme").
			Get("", new(acme.IndexAction)).
			GetPost("/create", new(acme.CreateAction)).
			Post("/run", new(acme.RunAction)).
			GetPost("/updateTaskPopup", new(acme.UpdateTaskPopupAction)).
			Post("/deleteTask", new(acme.DeleteTaskAction)).
			Prefix("/servers/certs/acme/users").
			Get("", new(users.IndexAction)).
			GetPost("/createPopup", new(users.CreatePopupAction)).
			GetPost("/updatePopup", new(users.UpdatePopupAction)).
			Post("/delete", new(users.DeleteAction)).
			GetPost("/selectPopup", new(users.SelectPopupAction)).
			EndAll()
	})
}
