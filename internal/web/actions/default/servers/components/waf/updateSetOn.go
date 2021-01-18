package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateSetOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateSetOnAction) RunPost(params struct {
	SetId int64
	IsOn  bool
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "修改WAF规则集 %d 开启状态", params.SetId)

	_, err := this.RPC().HTTPFirewallRuleSetRPC().UpdateHTTPFirewallRuleSetIsOn(this.UserContext(), &pb.UpdateHTTPFirewallRuleSetIsOnRequest{
		FirewallRuleSetId: params.SetId,
		IsOn:              params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
