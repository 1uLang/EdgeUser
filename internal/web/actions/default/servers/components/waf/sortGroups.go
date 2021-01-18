package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type SortGroupsAction struct {
	actionutils.ParentAction
}

func (this *SortGroupsAction) RunPost(params struct {
	FirewallPolicyId int64
	Type             string
	GroupIds         []int64
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "修改WAF策略 %d 中的规则分组中的排序", params.FirewallPolicyId)

	firewallPolicy, err := dao.SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicyConfig(this.UserContext(), params.FirewallPolicyId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if firewallPolicy == nil {
		this.NotFound("firewallPolicy", params.FirewallPolicyId)
		return
	}

	switch params.Type {
	case "inbound":
		refMapping := map[int64]*firewallconfigs.HTTPFirewallRuleGroupRef{}
		for _, ref := range firewallPolicy.Inbound.GroupRefs {
			refMapping[ref.GroupId] = ref
		}
		newRefs := []*firewallconfigs.HTTPFirewallRuleGroupRef{}
		for _, groupId := range params.GroupIds {
			ref, ok := refMapping[groupId]
			if ok {
				newRefs = append(newRefs, ref)
			}
		}
		firewallPolicy.Inbound.GroupRefs = newRefs
	case "outbound":
		refMapping := map[int64]*firewallconfigs.HTTPFirewallRuleGroupRef{}
		for _, ref := range firewallPolicy.Outbound.GroupRefs {
			refMapping[ref.GroupId] = ref
		}
		newRefs := []*firewallconfigs.HTTPFirewallRuleGroupRef{}
		for _, groupId := range params.GroupIds {
			ref, ok := refMapping[groupId]
			if ok {
				newRefs = append(newRefs, ref)
			}
		}
		firewallPolicy.Outbound.GroupRefs = newRefs
	}

	inboundJSON, err := firewallPolicy.InboundJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	outboundJSON, err := firewallPolicy.OutboundJSON()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallPolicyGroups(this.UserContext(), &pb.UpdateHTTPFirewallPolicyGroupsRequest{
		HttpFirewallPolicyId: params.FirewallPolicyId,
		InboundJSON:          inboundJSON,
		OutboundJSON:         outboundJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
