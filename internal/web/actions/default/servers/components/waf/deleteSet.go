package waf

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteSetAction struct {
	actionutils.ParentAction
}

func (this *DeleteSetAction) RunPost(params struct {
	GroupId int64
	SetId   int64
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "删除WAF规则分组 %d 中的规则集 %d", params.GroupId, params.SetId)

	groupConfig, err := dao.SharedHTTPFirewallRuleGroupDAO.FindRuleGroupConfig(this.UserContext(), params.GroupId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if groupConfig == nil {
		this.NotFound("firewallRuleGroup", params.GroupId)
		return
	}

	newRefs := []*firewallconfigs.HTTPFirewallRuleSetRef{}
	for _, ref := range groupConfig.SetRefs {
		if ref.SetId != params.SetId {
			newRefs = append(newRefs, ref)
		}
	}
	newRefsJSON, err := json.Marshal(newRefs)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, err = this.RPC().HTTPFirewallRuleGroupRPC().UpdateHTTPFirewallRuleGroupSets(this.UserContext(), &pb.UpdateHTTPFirewallRuleGroupSetsRequest{
		FirewallRuleGroupId:  params.GroupId,
		FirewallRuleSetsJSON: newRefsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
