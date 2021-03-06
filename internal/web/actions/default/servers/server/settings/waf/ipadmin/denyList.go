package ipadmin

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type DenyListAction struct {
	actionutils.ParentAction
}

func (this *DenyListAction) Init() {
	this.Nav("", "setting", "denyList")
	this.SecondMenu("waf")
}

func (this *DenyListAction) RunGet(params struct {
	ServerId         int64
	FirewallPolicyId int64
}) {
	this.Data["featureIsOn"] = true
	this.Data["firewallPolicyId"] = params.FirewallPolicyId

	// 校验权限
	//if !this.ValidateFeature("server.waf") {
	//	this.Data["featureIsOn"] = false
	//	this.Show()
	//	return
	//}

	listId, err := dao.SharedIPListDAO.FindDenyIPListIdWithServerId(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建
	if listId == 0 {
		listId, err = dao.SharedIPListDAO.CreateIPListForServerId(this.UserContext(), params.ServerId, "black")
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Data["listId"] = listId

	// 数量
	countResp, err := this.RPC().IPItemRPC().CountIPItemsWithListId(this.UserContext(), &pb.CountIPItemsWithListIdRequest{IpListId: listId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	// 列表
	itemsResp, err := this.RPC().IPItemRPC().ListIPItemsWithListId(this.UserContext(), &pb.ListIPItemsWithListIdRequest{
		IpListId: listId,
		Offset:   page.Offset,
		Size:     page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	itemMaps := []maps.Map{}
	for _, item := range itemsResp.IpItems {
		expiredTime := ""
		if item.ExpiredAt > 0 {
			expiredTime = timeutil.FormatTime("Y-m-d H:i:s", item.ExpiredAt)
		}

		itemMaps = append(itemMaps, maps.Map{
			"id":          item.Id,
			"ipFrom":      item.IpFrom,
			"ipTo":        item.IpTo,
			"expiredTime": expiredTime,
			"reason":      item.Reason,
		})
	}
	this.Data["items"] = itemMaps

	this.Show()
}
