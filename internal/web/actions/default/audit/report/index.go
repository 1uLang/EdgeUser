package report

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_from"
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
	PageNum  int
	PageSize int
	Type     string
	Name     string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	list, _ := audit_from.GetAuditFromList(&audit_from.ReqSearch{
		PageSize:   params.PageSize,
		PageNum:    params.PageNum,
		AssetsType: params.Type,
		Name:       params.Name,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	//this.Data["fromList"] = list.Data.List
	if list != nil && len(list.Data.List) > 0 {
		this.Data["fromList"] = list.Data.List
	} else {
		this.Data["fromList"] = []maps.Map{}
	}
	this.Show()
}
