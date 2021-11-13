package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	FirewallPolicyId int64
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "删除WAF策略 %d", params.FirewallPolicyId)

	countResp, err := this.RPC().NodeClusterRPC().CountAllEnabledNodeClustersWithHTTPFirewallPolicyId(this.UserContext(), &pb.CountAllEnabledNodeClustersWithHTTPFirewallPolicyIdRequest{HttpFirewallPolicyId: params.FirewallPolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if countResp.Count > 0 {
		this.Fail("此WAF策略正在被有些集群引用，请修改后再删除。")
	}

	_, err = this.RPC().HTTPFirewallPolicyRPC().DeleteHTTPFirewallPolicy(this.UserContext(), &pb.DeleteHTTPFirewallPolicyRequest{HttpFirewallPolicyId: params.FirewallPolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
