package db

import (
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server/audit_app"
	"github.com/1uLang/zhiannet-api/audit/server/audit_db"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"

	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "audit", "index")
}

func (this *IndexAction) RunGet() {
	this.Data["pageState"] = 1

	//数据库列表
	dblist, _ := audit_db.GetAuditBdList(&audit_db.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	if dblist != nil && len(dblist.Data.List) > 0 {
		this.Data["dbList"] = dblist.Data.List
	} else {
		this.Data["dbList"] = []maps.Map{}
	}

	//主机
	hostlist, _ := audit_host.GetAuditHostList(&audit_host.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	if hostlist != nil && len(hostlist.Data.List) > 0 {
		this.Data["hostList"] = hostlist.Data.List
	} else {
		this.Data["hostList"] = []maps.Map{}
	}

	//应用
	applist, _ := audit_app.GetAuditAppList(&audit_app.ReqSearch{
		PageSize: 999,
		PageNum:  1,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	//this.Data["appList"] = applist.Data.List
	if applist != nil && len(applist.Data.List) > 0 {
		this.Data["appList"] = applist.Data.List
	} else {
		this.Data["appList"] = []maps.Map{}
	}
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	PageNum   int64
	PageSize  int64
	TimeType  string
	StartTime string
	EndTime   string
	AuditId   []string
	CIp       string
	User      string
	Risk      string
	Message   string
	LogType   int
	Export    bool

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {

	var Stime, Etime time.Time
	if params.StartTime != "" {
		Stime, _ = time.ParseInLocation("2006-01-02", params.StartTime, time.Local)
	}
	if params.EndTime != "" {
		Etime, _ = time.ParseInLocation("2006-01-02", params.EndTime, time.Local)
		Etime = Etime.AddDate(0, 0, 1).Add(-time.Second)
	}
	var risk string
	if len(params.Risk) > 0 {
		risk = "true"
	}
	if len(params.AuditId) == 0 {
		params.AuditId = []string{"0"}
	}
	this.Data["list"] = []maps.Map{}
	switch params.LogType {
	case 1:
		list, _ := audit_db.GetDbLog(&audit_db.DbLogReq{
			Size:       params.PageSize,
			Page:       params.PageNum,
			StartTime:  Stime,
			EndTime:    Etime,
			Message:    params.Message,
			ClientHost: params.CIp,
			User:       params.User,
			TimeType:   params.TimeType,
			Risk:       risk,
			AuditId:    params.AuditId,
			Export:     params.Export,
			UserId: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		if list == nil {
			break
		}
		if params.Export {
			this.Data["filepath"] = list.Data.Filename
		} else {
			this.Data["list"] = list.Data.Log
		}
	case 2:
		list, _ := audit_host.GetHostLog(&audit_host.HostLogReq{
			Size:      params.PageSize,
			Page:      params.PageNum,
			StartTime: Stime,
			EndTime:   Etime,
			Message:   params.Message,
			TimeType:  params.TimeType,
			AuditId:   params.AuditId,
			Export:    params.Export,
			UserId: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		if list == nil {
			break
		}
		if params.Export {
			this.Data["filepath"] = list.Data.Filename
		} else {
			this.Data["list"] = list.Data.Log
		}
	case 3:
		list, _ := audit_app.GetAppLog(&audit_app.AppLogReq{
			Size:      params.PageSize,
			Page:      params.PageNum,
			StartTime: Stime,
			EndTime:   Etime,
			Message:   params.Message,
			TimeType:  params.TimeType,
			AuditId:   params.AuditId,
			Export:    params.Export,
			UserId: &request.UserReq{
				UserId: uint64(this.UserId()),
			},
		})
		if list == nil {
			break
		}
		if params.Export {
			this.Data["filepath"] = list.Data.Filename
		} else {
			this.Data["list"] = list.Data.Log
		}
	}

	this.Success()
}
