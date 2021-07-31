package databackup

import (
	"bytes"
	"log"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	// 获取token
	req := model.LoginReq{
		User:     "admin",
		Password: "Dengbao123!@#",
	}
	token := request.GenerateToken(&req)

	// 文件列表（不包含目录）
	list, err := request.ListFolders(token)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["list"] = list.List
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	UploadFile *actions.File `json:"uploadFile"`
}) {
	// 获取token
	req := model.LoginReq{
		User:     "admin",
		Password: "Dengbao123!@#",
	}
	token := request.GenerateToken(&req)
	log.Println(123)
	// 上传文件
	if params.UploadFile == nil {
		this.Fail("请选择要上传的文件")
	}
	upFile, err := params.UploadFile.Read()
	if err != nil {
		this.Fail("读取文件内容错误，请重新上传")
	}
	name := params.UploadFile.Filename
	err = request.UploadFile(token, name, bytes.NewReader(upFile))
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "上传nextcloud文件 %v", name)

	this.Success()
}
