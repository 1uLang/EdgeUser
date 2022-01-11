package host

import (
	"github.com/1uLang/zhiannet-api/audit"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "audit", "index")
}

func (this *IndexAction) RunGet(params struct {
	Page     int
	PageSize int
	Ip       string
	Name     string
	Status   string
	Json     bool

	Must *actions.Must
}) {
	list, _ := audit_host.GetAuditHostList(&audit_host.ReqSearch{
		PageSize: params.PageSize,
		PageNum:  params.Page,
		Ip:       params.Ip,
		Name:     params.Name,
		Status:   params.Status,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	//this.Data["hostList"] = list.Data.List
	count := int64(0)
	if list != nil && len(list.Data.List) > 0 {
		count = int64(list.Data.Total)
		this.Data["hostList"] = list.Data.List
	} else {
		this.Data["hostList"] = []maps.Map{}
	}
	page := this.NewPage(int64(count))
	this.Data["page"] = page.AsHTML()
	this.Data["log_submit_addr"] = audit.LogSubmitAddr
	if params.Json {
		this.Success()
	}
	this.Show()
}
