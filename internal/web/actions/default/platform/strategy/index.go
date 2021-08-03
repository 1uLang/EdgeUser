package strategy

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "platform", "index")
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
	//list, _ := audit_app.GetAuditAppList(&audit_app.ReqSearch{
	//	PageSize: params.PageSize,
	//	PageNum:  params.PageNum,
	//	AppType:  params.Type,
	//	Ip:       params.Ip,
	//	Name:     params.Name,
	//	Status:   params.Status,
	//	User: &request.UserReq{
	//		AdminUserId: uint64(this.AdminId()),
	//	},
	//})
	////this.Data["appList"] = list.Data.List
	//if list != nil && len(list.Data.List) > 0 {
	//	this.Data["appList"] = list.Data.List
	//} else {
	//	this.Data["appList"] = []maps.Map{}
	//}
	//if params.Json {
	//	this.Success()
	//}
	this.Show()
}
