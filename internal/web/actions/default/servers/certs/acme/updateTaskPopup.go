package acme

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/dns/domains/domainutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"strings"
)

type UpdateTaskPopupAction struct {
	actionutils.ParentAction
}

func (this *UpdateTaskPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateTaskPopupAction) RunGet(params struct {
	TaskId int64
}) {
	taskResp, err := this.RPC().ACMETaskRPC().FindEnabledACMETask(this.UserContext(), &pb.FindEnabledACMETaskRequest{AcmeTaskId: params.TaskId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	task := taskResp.AcmeTask
	if task == nil {
		this.NotFound("acmeTask", params.TaskId)
		return
	}

	var dnsProviderMap maps.Map
	if task.DnsProvider != nil {
		dnsProviderMap = maps.Map{
			"id": task.DnsProvider.Id,
		}
	} else {
		dnsProviderMap = maps.Map{
			"id": 0,
		}
	}

	var acmeUserMap maps.Map
	if task.AcmeUser != nil {
		acmeUserMap = maps.Map{
			"id": task.AcmeUser.Id,
		}
	} else {
		acmeUserMap = maps.Map{
			"id": 0,
		}
	}

	this.Data["task"] = maps.Map{
		"id":          task.Id,
		"authType":    task.AuthType,
		"acmeUser":    acmeUserMap,
		"dnsDomain":   task.DnsDomain,
		"domains":     task.Domains,
		"autoRenew":   task.AutoRenew,
		"isOn":        task.IsOn,
		"authURL":     task.AuthURL,
		"dnsProvider": dnsProviderMap,
	}

	// 域名解析服务商
	providersResp, err := this.RPC().DNSProviderRPC().FindAllEnabledDNSProviders(this.UserContext(), &pb.FindAllEnabledDNSProvidersRequest{
		AdminId: 0,
		UserId:  this.UserId(true),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	providerMaps := []maps.Map{}
	for _, provider := range providersResp.DnsProviders {
		providerMaps = append(providerMaps, maps.Map{
			"id":       provider.Id,
			"name":     provider.Name,
			"typeName": provider.TypeName,
		})
	}
	this.Data["providers"] = providerMaps

	this.Show()
}

func (this *UpdateTaskPopupAction) RunPost(params struct {
	TaskId        int64
	AuthType      string
	AcmeUserId    int64
	DnsProviderId int64
	DnsDomain     string
	Domains       []string
	AutoRenew     bool
	AuthURL       string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("修改证书申请任务 %d", params.TaskId)

	if params.AuthType != "dns" && params.AuthType != "http" {
		this.Fail("无法识别的认证方式'" + params.AuthType + "'")
	}

	if params.AcmeUserId <= 0 {
		this.Fail("请选择一个申请证书的用户")
	}

	dnsDomain := strings.ToLower(params.DnsDomain)
	if params.AuthType == "dns" {
		if params.DnsProviderId <= 0 {
			this.Fail("请选择DNS服务商")
		}
		if len(params.DnsDomain) == 0 {
			this.Fail("请输入顶级域名")
		}
		if !domainutils.ValidateDomainFormat(dnsDomain) {
			this.Fail("请输入正确的顶级域名")
		}
	}

	if len(params.Domains) == 0 {
		this.Fail("请输入证书域名列表")
	}
	realDomains := []string{}
	for _, domain := range params.Domains {
		domain = strings.ToLower(domain)
		if params.AuthType == "dns" {
			if !strings.HasSuffix(domain, "."+dnsDomain) && domain != dnsDomain {
				this.Fail("证书域名中的" + domain + "和顶级域名不一致")
			}
		} else if params.AuthType == "http" { // HTTP认证
			if strings.Contains(domain, "*") {
				this.Fail("在HTTP认证时域名" + domain + "不能包含通配符")
			}
		}
		realDomains = append(realDomains, domain)
	}

	_, err := this.RPC().ACMETaskRPC().UpdateACMETask(this.UserContext(), &pb.UpdateACMETaskRequest{
		AcmeTaskId:    params.TaskId,
		AcmeUserId:    params.AcmeUserId,
		DnsProviderId: params.DnsProviderId,
		DnsDomain:     dnsDomain,
		Domains:       realDomains,
		AutoRenew:     params.AutoRenew,
		AuthURL:       params.AuthURL,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}