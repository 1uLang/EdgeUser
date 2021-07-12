package targets

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
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

	Show int
}) {
	data := make([]interface{}, 0)
	this.Data["nodeErr"] = ""
	this.Data["targets"] = data
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取扫描节点信息失败：%v", err))
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	list, err := targets_server.List(&targets.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list != nil {
		this.ErrorPage(err)
		return
	}
	if lists, ok := list["targets"]; ok {
		this.Data["targets"] = lists
	}
	if params.Show == 1 {
		this.Success()
	}
	this.Show()
}
func (this *IndexAction) RunPost(params struct {
	PageSize int
	PageNo   int
}) {
	data := make([]interface{}, 0)
	this.Data["nodeErr"] = ""
	this.Data["targets"] = data
	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取扫描节点信息失败：%v", err))
		return
	}
	if params.PageNo < 0 {
		params.PageNo = 0
	}
	if params.PageSize < 0 {
		params.PageSize = 20
	}
	list, err := targets_server.List(&targets.ListReq{Limit: params.PageSize, C: params.PageNo * params.PageSize, UserId: uint64(this.UserId())})
	if err != nil && list != nil {
		this.ErrorPage(err)
		return
	}
	if lists, ok := list["targets"]; ok {
		this.Data["targets"] = lists
	}
	this.Success()
}
