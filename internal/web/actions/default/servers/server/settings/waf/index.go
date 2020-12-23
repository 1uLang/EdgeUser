package waf

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/websocket/webutils"
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
	webConfig, err := webutils.FindWebConfigWithServerId(this.Parent(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["firewallConfig"] = webConfig.FirewallRef

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId        int64
	FirewallJSON []byte

	Must *actions.Must
}) {
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
