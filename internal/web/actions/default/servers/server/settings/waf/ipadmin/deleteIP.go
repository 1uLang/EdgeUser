package ipadmin

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteIPAction struct {
	actionutils.ParentAction
}

func (this *DeleteIPAction) RunPost(params struct {
	FirewallPolicyId int64
	ItemId           int64
}) {
	// 校验权限
	if !this.ValidateFeature("server.waf") {
		return
	}

	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "从WAF策略 %d 名单中删除IP %d", params.FirewallPolicyId, params.ItemId)

	// TODO 判断权限

	_, err := this.RPC().IPItemRPC().DeleteIPItem(this.UserContext(), &pb.DeleteIPItemRequest{IpItemId: params.ItemId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
