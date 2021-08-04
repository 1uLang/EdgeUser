package certs

import (
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

			EndAll()
	})
}
