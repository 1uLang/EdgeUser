package databackup

import (
	em "github.com/1uLang/zhiannet-api/edgeUsers/model"
	es "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) Init() {
	this.Nav("", "", "")
}

func (this *DeleteAction) RunPost(params struct {
	Name string
}) {
	uid, _ := es.GetParentId(&em.GetParentIdReq{UserId: uint64(this.UserId())})
	if uid == 0 {
		uid = uint64(this.UserId())
	}
	// 获取token
	token, err := model.QueryTokenByUID(int64(uid))
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 删除文件
	err = request.DeleteFile(token, params.Name)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除数据备份文件 %v", params.Name)

	this.Success()
}
