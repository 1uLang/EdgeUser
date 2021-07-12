package vulnerabilities

import (
	"github.com/1uLang/zhiannet-api/awvs/model/vulnerabilities"
	vulnerabilities_server "github.com/1uLang/zhiannet-api/awvs/server/vulnerabilities"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
)

//任务目标
type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
	Address  string
	Severity string

	List bool
}) {
	this.Data["nodeErr"] = ""
	this.Data["vulnerabilities"] = make([]interface{}, 0)
	err := webscan.InitAPIServer()
	if err != nil {
		//this.ErrorPage(err)
		this.Data["nodeErr"] = "获取web扫描节点错误"
		this.Show()
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	var query string
	if params.Address != "" {
		query += "target_id:" + params.Address
		query += ";"
	}
	if params.Severity != "" {
		query += "severity:" + params.Severity
		query += ";"
	}

	list, err := vulnerabilities_server.List(&vulnerabilities.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, Query: query, UserId: uint64(this.UserId())})
	if err != nil && list != nil {
		//this.ErrorPage(err)
		this.Show()
		return
	}
	//this.Data["vulnerabilities"] = list["vulnerabilities"]
	if lists, ok := list["vulnerabilities"]; ok {
		this.Data["vulnerabilities"] = lists
	}
	if !params.List {
		this.Show()
	} else {
		this.Success()
	}
}
