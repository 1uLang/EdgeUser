package agent

import (
	"github.com/1uLang/zhiannet-api/agent/server"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "audit", "index")
}

func (this *IndexAction) RunGet(params struct {
	PageNum  int
	PageSize int
}) {
	rsp, err := server.ListFile(params.PageNum, params.PageSize)
	// rsp, err := server.ListFile()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["list"] = rsp.List
	this.Data["total"] = rsp.Total
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	UploadFile *actions.File `json:"uploadFile"`
	FileDesc   string        `json:"fileDesc"`
	Format     string        `json:"format"`
}) {
	// 上传文件
	if params.UploadFile == nil {
		this.Fail("请选择要上传的文件")
	}
	upFile, err := params.UploadFile.Read()
	if err != nil {
		this.Fail("读取文件内容错误，请重新上传")
	}
	name := params.UploadFile.Filename
	err = server.UploadFile(name, params.Format, params.FileDesc, upFile)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "上传backup文件 %v", name)

	this.Success()
	// this.Show()
}
