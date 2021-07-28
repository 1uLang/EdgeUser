package app

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
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
	Type     string
	Ip       string
	Name     string
	Status   string
	Json     bool

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	list, _ := audit_app.GetAuditAppList(&audit_app.ReqSearch{
		PageSize: params.PageSize,
		PageNum:  params.PageNum,
		AppType:  params.Type,
		Ip:       params.Ip,
		Name:     params.Name,
		Status:   params.Status,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	this.Data["appList"] = list.Data.List
	fmt.Println("params==", params)
	fmt.Println("list==", list)
	if params.Json {
		this.Success()
	}
	this.Show()
}
