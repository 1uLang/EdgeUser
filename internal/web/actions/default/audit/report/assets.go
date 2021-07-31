package report

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
	"github.com/1uLang/zhiannet-api/audit/server/audit_db"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type AssetsAction struct {
	actionutils.ParentAction
}

func (this *AssetsAction) RunGet(params struct {
	AssetsType int

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	dblist, _ := audit_db.GetAuditBdList(&audit_db.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	this.Data["dbAssetsList"] = dblist.Data.List

	hostlist, _ := audit_host.GetAuditHostList(&audit_host.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	this.Data["hostAssetsList"] = hostlist.Data.List

	applist, _ := audit_app.GetAuditAppList(&audit_app.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	this.Data["appAssetsList"] = applist.Data.List

	this.Success()
}
