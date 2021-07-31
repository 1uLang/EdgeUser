package databackup

import (
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
	// 获取token
	req := model.LoginReq{
		User:     "admin",
		Password: "Dengbao123!@#",
	}
	token := request.GenerateToken(&req)

	// 删除文件
	err := request.DeleteFile(token, params.Name)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	defer this.CreateLogInfo("删除数据备份文件 %v", params.Name)

	this.Success()
}
