package databackup

import (
	em "github.com/1uLang/zhiannet-api/edgeUsers/model"
	es "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type RenameAction struct {
	actionutils.ParentAction
}

func (this *RenameAction) Init() {
	this.Nav("", "", "")
}

func (this *RenameAction) RunGet(params struct {
	Name string
	Url  string
}) {
	this.Data["name"] = params.Name
	this.Data["oldname"] = params.Name
	this.Data["url"] = params.Url
	this.Show()
}

func (this *RenameAction) RunPost(params struct {
	SrcURL  string
	NewName string

	Name string
	Must *actions.Must
}) {
	params.Must.
		Field("srcURL", params.SrcURL).
		Require("请输入原文件或文件夹的url路径").
		Field("newName", params.NewName).
		Require("请输入新的的名字")
	if params.SrcURL == "" || params.NewName == "" {
		this.Fail("url路径，新名字不能为空")
	}

	if params.Name == params.NewName {
		this.Success()
	}
	uid, _ := es.GetParentId(&em.GetParentIdReq{UserId: uint64(this.UserId())})
	if uid == 0 {
		uid = uint64(this.UserId())
	}
	// 获取token
	token, err := model.QueryTokenByUID(int64(uid))
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	err = request.MoveFileOrFolder(params.SrcURL, params.NewName, token)
	if err != nil {
		this.FailField("error", err.Error())
		return
	}

	this.Success()
}
