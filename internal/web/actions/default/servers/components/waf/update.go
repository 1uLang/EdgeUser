package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"net/http"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	FirewallPolicyId int64
}) {
	firewallPolicy, err := dao.SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicyConfig(this.UserContext(), params.FirewallPolicyId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if firewallPolicy == nil {
		this.NotFound("firewallPolicy", params.FirewallPolicyId)
		return
	}

	if firewallPolicy.BlockOptions == nil {
		firewallPolicy.BlockOptions = &firewallconfigs.HTTPFirewallBlockAction{
			StatusCode: http.StatusForbidden,
			Body:       "Blocked By WAF",
			URL:        "",
		}
	}

	this.Data["firewallPolicy"] = maps.Map{
		"id":           firewallPolicy.Id,
		"name":         firewallPolicy.Name,
		"description":  firewallPolicy.Description,
		"isOn":         firewallPolicy.IsOn,
		"blockOptions": firewallPolicy.BlockOptions,
	}

	// 预置分组
	groups := []maps.Map{}
	templatePolicy := firewallconfigs.HTTPFirewallTemplate()
	for _, group := range templatePolicy.AllRuleGroups() {
		if len(group.Code) > 0 {
			usedGroup := firewallPolicy.FindRuleGroupWithCode(group.Code)
			if usedGroup != nil {
				group.IsOn = usedGroup.IsOn
			}
		}

		groups = append(groups, maps.Map{
			"code": group.Code,
			"name": group.Name,
			"isOn": group.IsOn,
		})
	}
	this.Data["groups"] = groups

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	FirewallPolicyId int64
	Name             string
	GroupCodes       []string
	BlockOptionsJSON []byte
	Description      string
	IsOn             bool

	Must *actions.Must
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "修改WAF策略 %d 基本信息", params.FirewallPolicyId)

	params.Must.
		Field("name", params.Name).
		Require("请输入策略名称")

	_, err := this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallPolicy(this.UserContext(), &pb.UpdateHTTPFirewallPolicyRequest{
		HttpFirewallPolicyId: params.FirewallPolicyId,
		IsOn:                 params.IsOn,
		Name:                 params.Name,
		Description:          params.Description,
		FirewallGroupCodes:   params.GroupCodes,
		BlockOptionsJSON:     params.BlockOptionsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
