package report

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
	"github.com/1uLang/zhiannet-api/audit/server/audit_db"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type AssetsAction struct {
	actionutils.ParentAction
}

func (this *AssetsAction) RunGet(params struct {
	AssetsType int

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	fmt.Println("AssetsType====", params.AssetsType)
	this.Data["assetsList"] = []maps.Map{}
	switch params.AssetsType {
	case 1:
		list, _ := audit_db.GetAuditBdList(&audit_db.ReqSearch{
			PageSize: 999,
			PageNum:  1,
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		this.Data["assetsList"] = list.Data.List
	case 2:
		list, _ := audit_host.GetAuditHostList(&audit_host.ReqSearch{
			PageSize: 999,
			PageNum:  1,
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		this.Data["assetsList"] = list.Data.List
	case 3:
		list, _ := audit_app.GetAuditAppList(&audit_app.ReqSearch{
			PageSize: 999,
			PageNum:  1,
			User: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		this.Data["assetsList"] = list.Data.List
	}

	this.Success()
}
