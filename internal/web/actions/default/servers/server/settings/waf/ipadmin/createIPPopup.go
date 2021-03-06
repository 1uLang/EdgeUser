package ipadmin

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreateIPPopupAction struct {
	actionutils.ParentAction
}

func (this *CreateIPPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreateIPPopupAction) RunGet(params struct {
	ListId int64
	Type   string
}) {
	// 校验权限
	//if !this.ValidateFeature("server.waf") {
	//	return
	//}

	this.Data["type"] = params.Type
	this.Data["listId"] = params.ListId

	this.Show()
}

func (this *CreateIPPopupAction) RunPost(params struct {
	ListId    int64
	IpFrom    string
	IpTo      string
	ExpiredAt int64
	Reason    string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	// 校验权限
	//if !this.ValidateFeature("server.waf") {
	//	return
	//}

	// TODO 校验ListId所属用户

	params.Must.
		Field("ipFrom", params.IpFrom).
		Require("请输入开始IP")

	// 校验IP格式（ipFrom/ipTo）
	ipFromLong := utils.IP2Long(params.IpFrom)
	if len(params.IpFrom) > 0 {
		if ipFromLong == 0 {
			this.Fail("请输入正确的开始IP")
		}
	}

	ipToLong := utils.IP2Long(params.IpTo)
	if len(params.IpTo) > 0 {
		if ipToLong == 0 {
			this.Fail("请输入正确的结束IP")
		}
	}

	if ipFromLong > 0 && ipToLong > 0 && ipFromLong > ipToLong {
		params.IpTo, params.IpFrom = params.IpFrom, params.IpTo
	}

	createResp, err := this.RPC().IPItemRPC().CreateIPItem(this.UserContext(), &pb.CreateIPItemRequest{
		IpListId:  params.ListId,
		IpFrom:    params.IpFrom,
		IpTo:      params.IpTo,
		ExpiredAt: params.ExpiredAt,
		Reason:    params.Reason,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	itemId := createResp.IpItemId


	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "在WAF策略 %d 名单中添加IP %d", params.ListId, itemId)

	this.Success()
}
