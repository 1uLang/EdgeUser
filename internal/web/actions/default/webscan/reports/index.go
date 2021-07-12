package reports

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
	reports_server "github.com/1uLang/zhiannet-api/awvs/server/reports"
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
	this.Data["reports"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		//this.ErrorPage(err)
		this.Data["nodeErr"] = "web漏洞扫描节点错误"
		this.Show()
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	list, err := reports_server.List(&reports.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list == nil {
		//this.ErrorPage(err)
		this.Show()
		return
	}
	//this.Data["reports"] = list["reports"]
	if lists, ok := list["reports"]; ok {
		this.Data["reports"] = lists
	}
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	PageSize int
	PageNo   int
}) {
	this.Data["nodeErr"] = ""
	this.Data["reports"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(fmt.Errorf("web漏洞扫描节点错误:%v",err))
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	list, err := reports_server.List(&reports.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list == nil {
		this.ErrorPage(err)
		return
	}
	//this.Data["reports"] = list["reports"]
	if lists, ok := list["reports"]; ok {
		this.Data["reports"] = lists
	}
	this.Show()
}
