package databackup

import (
	"fmt"
	"io"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DownLoadAction struct {
	actionutils.ParentAction
}

func (this *DownLoadAction) Init() {
	this.Nav("", "", "")
}

func (this *DownLoadAction) RunGet(params struct {
	Name string
}) {
	// 获取token
	token, err := model.QueryTokenByUID(this.UserId())
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 获取下载链接
	// dURL, err := request.DownLoadFileWithURL(token, params.Name)
	// if err != nil {
	// 	this.ErrorPage(err)
	// 	return
	// }
	// log.Printf("token: %s\n",token)
	// log.Printf("url: %s\n",dURL)
	// this.Data["token"] = token
	// this.Data["url"] = dURL

	// 获取下载字节流
	rsp, err := request.DownLoadFile(token, params.Name)
	if err != nil {
		this.ErrorPage(fmt.Errorf("下载文件失败：%w", err))
		return
	}
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取文件字节流失败：%w", err))
		return
	}

	this.Data["body"] = body
	this.Data["fileName"] = params.Name

	this.Success()
}
