package reports

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

type DownloadAction struct {
	actionutils.ParentAction
}

func (this *DownloadAction) RunGet(params struct {
	Path string
	Must *actions.Must
}) {

	params.Must.
		Field("path", params.Path).
		Require("请输入下载路径")

	this.Data["url"] = webscan.ServerUrl + params.Path
	this.Success()
}
