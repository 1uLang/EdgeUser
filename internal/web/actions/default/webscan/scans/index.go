package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
}) {
	this.Data["nodeErr"] = ""
	this.Data["scans"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		//this.ErrorPage(err)
		this.Data["nodeErr"] = "web扫描节点错误"
		this.Show()
		return
	}
	if params.PageNo <= 0 {
		params.PageNo = 0
	}
	if params.PageSize <= 0 {
		params.PageSize = 20
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list != nil {
		this.ErrorPage(err)
		this.Show()
		return
	}
	if lists, ok := list["scans"]; ok {
		this.Data["scans"] = lists
	}
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	PageSize int
	PageNo   int
}) {
	this.Data["scans"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.PageNo <= 0 {
		params.PageNo = 0
	}
	if params.PageSize <= 0 {
		params.PageSize = 20
	}
	list, err := scans_server.List(&scans.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list != nil {
		this.ErrorPage(err)
		return
	}
	if lists, ok := list["scans"]; ok {
		this.Data["scans"] = lists
	}
	this.Success()
}
