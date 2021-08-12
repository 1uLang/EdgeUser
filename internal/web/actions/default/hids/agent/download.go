package agent

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DownloadAction struct {
	actionutils.ParentAction
}

func (this *DownloadAction) Init() {
	this.FirstMenu("index")
}

func (this *DownloadAction) RunGet(params struct {
	UserName string
	osType   string

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	this.Show()
	return
	//params.Must.
	//	Field("username", params.UserName).
	//	Require("请输入用户名")
	//
	//params.Must.
	//	Field("osType", params.osType).
	//	Require("请选择主机操作系统")

}
