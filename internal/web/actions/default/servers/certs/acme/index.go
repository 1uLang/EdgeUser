package acme

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "task")
	this.SecondMenu("list")
}

func (this *IndexAction) RunGet(params struct {
	Type    string
	Keyword string
}) {
	this.Data["type"] = params.Type
	this.Data["keyword"] = params.Keyword

	countAll := int64(0)
	countAvailable := int64(0)
	countExpired := int64(0)
	count7Days := int64(0)
	count30Days := int64(0)

	// 计算数量
	{
		// all
		resp, err := this.RPC().ACMETaskRPC().CountAllEnabledACMETasks(this.UserContext(), &pb.CountAllEnabledACMETasksRequest{
			UserId: this.UserId(true),
			Keyword: params.Keyword,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countAll = resp.Count

		// available
		resp, err = this.RPC().ACMETaskRPC().CountAllEnabledACMETasks(this.UserContext(), &pb.CountAllEnabledACMETasksRequest{
			UserId: this.UserId(true),
			IsAvailable: true,
			Keyword:     params.Keyword,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countAvailable = resp.Count

		// expired
		resp, err = this.RPC().ACMETaskRPC().CountAllEnabledACMETasks(this.UserContext(), &pb.CountAllEnabledACMETasksRequest{
			UserId: this.UserId(true),
			IsExpired: true,
			Keyword:   params.Keyword,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countExpired = resp.Count

		// expire in 7 days
		resp, err = this.RPC().ACMETaskRPC().CountAllEnabledACMETasks(this.UserContext(), &pb.CountAllEnabledACMETasksRequest{
			UserId: this.UserId(true),
			ExpiringDays: 7,
			Keyword:      params.Keyword,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		count7Days = resp.Count

		// expire in 30 days
		resp, err = this.RPC().ACMETaskRPC().CountAllEnabledACMETasks(this.UserContext(), &pb.CountAllEnabledACMETasksRequest{
			UserId: this.UserId(true),
			ExpiringDays: 30,
			Keyword:      params.Keyword,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		count30Days = resp.Count
	}

	this.Data["countAll"] = countAll
	this.Data["countAvailable"] = countAvailable
	this.Data["countExpired"] = countExpired
	this.Data["count7Days"] = count7Days
	this.Data["count30Days"] = count30Days

	// 分页
	var page *actionutils.Page
	var tasksResp *pb.ListEnabledACMETasksResponse
	var err error
	switch params.Type {
	case "":
		page = this.NewPage(countAll)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			Offset:  page.Offset,
			Size:    page.Size,
			Keyword: params.Keyword,
		})
	case "available":
		page = this.NewPage(countAvailable)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			IsAvailable: true, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword})
	case "expired":
		page = this.NewPage(countExpired)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			IsExpired: true, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword})
	case "7days":
		page = this.NewPage(count7Days)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			ExpiringDays: 7, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword})
	case "30days":
		page = this.NewPage(count30Days)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			ExpiringDays: 30, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword})
	default:
		page = this.NewPage(countAll)
		tasksResp, err = this.RPC().ACMETaskRPC().ListEnabledACMETasks(this.UserContext(), &pb.ListEnabledACMETasksRequest{
			UserId: this.UserId(true),
			Keyword: params.Keyword,
			Offset:  page.Offset,
			Size:    page.Size,
		})
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["page"] = page.AsHTML()

	taskMaps := []maps.Map{}
	for _, task := range tasksResp.AcmeTasks {
		if task.AcmeUser == nil {
			continue
		}
		dnsProviderMap := maps.Map{}
		if task.AuthType == "dns" && task.DnsProvider != nil {
			dnsProviderMap = maps.Map{
				"id":   task.DnsProvider.Id,
				"name": task.DnsProvider.Name,
			}
		}

		// 证书
		var certMap maps.Map = nil
		if task.SslCert != nil {
			certMap = maps.Map{
				"id":        task.SslCert.Id,
				"name":      task.SslCert.Name,
				"beginTime": timeutil.FormatTime("Y-m-d", task.SslCert.TimeBeginAt),
				"endTime":   timeutil.FormatTime("Y-m-d", task.SslCert.TimeEndAt),
			}
		}

		// 日志
		var logMap maps.Map = nil
		if task.LatestACMETaskLog != nil {
			logMap = maps.Map{
				"id":          task.LatestACMETaskLog.Id,
				"isOk":        task.LatestACMETaskLog.IsOk,
				"error":       task.LatestACMETaskLog.Error,
				"createdTime": timeutil.FormatTime("m-d", task.CreatedAt),
			}
		}

		taskMaps = append(taskMaps, maps.Map{
			"id":       task.Id,
			"authType": task.AuthType,
			"acmeUser": maps.Map{
				"id":    task.AcmeUser.Id,
				"email": task.AcmeUser.Email,
			},
			"dnsProvider": dnsProviderMap,
			"dnsDomain":   task.DnsDomain,
			"domains":     task.Domains,
			"autoRenew":   task.AutoRenew,
			"cert":        certMap,
			"log":         logMap,
		})
	}
	this.Data["tasks"] = taskMaps

	this.Show()
}
