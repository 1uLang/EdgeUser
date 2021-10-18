package databackup

import (
	em "github.com/1uLang/zhiannet-api/edgeUsers/model"
	es "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type DirAction struct {
	actionutils.ParentAction
}

func (this *DirAction) Init() {
	this.Nav("", "", "")
}

func (this *DirAction) RunPost(params struct {
	Purl string
	Name string

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入文件名")
	if params.Name == "" {
		this.Fail("Purl不能为空")
	}
	uid, _ := es.GetParentId(&em.GetParentIdReq{UserId: uint64(this.UserId())})
	if uid == 0 {
		uid = uint64(this.UserId())
	}
	// 获取token
	token, err := model.QueryTokenByUID(int64(uid))
	if err != nil {
		this.FailField("error",err.Error())
		return
	}

	err = request.CreateFoler(token, params.Purl, params.Name)
	if err != nil {
		this.FailField("error",err.Error())
		return
	}

	this.Success()
}
