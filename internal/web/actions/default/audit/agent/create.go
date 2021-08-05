package agent

import (
	"github.com/1uLang/zhiannet-api/agent/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct {
	Id int64 `json:"id"`

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请输入文件id")
	if params.Id <= 0 {
		this.Fail("id必须大于0")
	}

	af, err := server.GetFileInfo(params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["name"] = af.Name
	this.Data["describe"] = af.Describe
	this.Data["id"] = af.ID

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Describe string `json:"describe"`

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Id).
		Require("请输入文件id")
	if params.Id <= 0 {
		this.Fail("id必须大于0")
	}

	err := server.UpdateFileInfo(params.Name, params.Describe, params.Id)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
