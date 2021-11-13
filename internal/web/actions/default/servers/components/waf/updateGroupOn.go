package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateGroupOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateGroupOnAction) RunPost(params struct {
	GroupId int64
	IsOn    bool
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "设置WAF规则分组 %d 开启状态", params.GroupId)

	_, err := this.RPC().HTTPFirewallRuleGroupRPC().UpdateHTTPFirewallRuleGroupIsOn(this.UserContext(), &pb.UpdateHTTPFirewallRuleGroupIsOnRequest{
		FirewallRuleGroupId: params.GroupId,
		IsOn:                params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
