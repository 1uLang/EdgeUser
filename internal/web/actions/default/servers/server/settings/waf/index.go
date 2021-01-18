package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("waf")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	// 校验权限
	this.Data["featureIsOn"] = this.ValidateFeature("server.waf")

	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["firewallConfig"] = webConfig.FirewallRef

	// 当前的Server独立设置
	if webConfig.FirewallRef == nil || webConfig.FirewallRef.FirewallPolicyId == 0 {
		firewallPolicyId, err := dao.SharedHTTPWebDAO.InitEmptyHTTPFirewallPolicy(this.UserContext(), webConfig.Id, webConfig.FirewallRef != nil && webConfig.FirewallRef.IsOn)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		this.Data["firewallPolicyId"] = firewallPolicyId
	} else {
		this.Data["firewallPolicyId"] = webConfig.FirewallRef.FirewallPolicyId
	}

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId        int64
	FirewallJSON []byte

	Must *actions.Must
}) {
	// 校验权限
	if !this.ValidateFeature("server.waf") {
		return
	}

	defer this.CreateLogInfo("修改Web %d 的WAF设置", params.WebId)

	// TODO 检查配置

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebFirewall(this.UserContext(), &pb.UpdateHTTPWebFirewallRequest{
		WebId:        params.WebId,
		FirewallJSON: params.FirewallJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
