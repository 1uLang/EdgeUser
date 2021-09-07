package reports

import (
	"github.com/1uLang/zhiannet-api/awvs/server/reports"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DownloadAction struct {
	actionutils.ParentAction
}

func (this *DownloadAction) RunGet(params struct {
	Path string
	PDF  bool
	Must *actions.Must
}) {

	params.Must.
		Field("path", params.Path).
		Require("请输入下载路径")

	bytes, contents, err := reports.Download(webscan.ServerUrl+params.Path, params.PDF)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.AddHeader("Content-Disposition", contents)
	this.Write(bytes)
}
