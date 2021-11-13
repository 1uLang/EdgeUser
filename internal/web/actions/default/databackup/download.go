package databackup

import (
	"fmt"
	"io"

	em "github.com/1uLang/zhiannet-api/edgeUsers/model"
	es "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
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
	Name string
	Fp   string

	Must *actions.Must
}) {
	params.Must.
		Field("id", params.Fp).
		Require("请输入文件id")
	if params.Fp == "" {
		this.Fail("fp不能为空")
	}
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
	// rsp, err := request.DownLoadFile(token, params.Name)
	rsp, err := request.DownLoadFileWithPath(token, params.Fp)
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
