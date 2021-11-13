package agent

import (
	"github.com/1uLang/zhiannet-api/agent/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DownLoadAction struct {
	actionutils.ParentAction
}

func (this *DownLoadAction) Init() {
	this.Nav("", "", "")
}

func (this *DownLoadAction) RunGet(params struct {
	Id int64 `json:"id"`

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请输入文件id")
	if params.Id <= 0 {
		this.Fail("id必须大于0")
	}
	
	body, err := server.GetFileStream(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["body"] = body
	this.Success()
}
