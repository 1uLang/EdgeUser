package headers

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/websocket/webutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("header")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	webConfig, err := webutils.FindWebConfigWithServerId(this.Parent(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	webId := webConfig.Id

	isChanged := false
	if webConfig.RequestHeaderPolicy == nil {
		createHeaderPolicyResp, err := this.RPC().HTTPHeaderPolicyRPC().CreateHTTPHeaderPolicy(this.UserContext(), &pb.CreateHTTPHeaderPolicyRequest{})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		headerPolicyId := createHeaderPolicyResp.HeaderPolicyId
		ref := &shared.HTTPHeaderPolicyRef{
			IsPrior:        false,
			IsOn:           true,
			HeaderPolicyId: headerPolicyId,
		}
		refJSON, err := json.Marshal(ref)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebRequestHeader(this.UserContext(), &pb.UpdateHTTPWebRequestHeaderRequest{
			WebId:      webId,
			HeaderJSON: refJSON,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		isChanged = true
	}
	if webConfig.ResponseHeaderPolicy == nil {
		createHeaderPolicyResp, err := this.RPC().HTTPHeaderPolicyRPC().CreateHTTPHeaderPolicy(this.UserContext(), &pb.CreateHTTPHeaderPolicyRequest{})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		headerPolicyId := createHeaderPolicyResp.HeaderPolicyId
		ref := &shared.HTTPHeaderPolicyRef{
			IsPrior:        false,
			IsOn:           true,
			HeaderPolicyId: headerPolicyId,
		}
		refJSON, err := json.Marshal(ref)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebResponseHeader(this.UserContext(), &pb.UpdateHTTPWebResponseHeaderRequest{
			WebId:      webId,
			HeaderJSON: refJSON,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		isChanged = true
	}

	// 重新获取配置
	if isChanged {
		webConfig, err = webutils.FindWebConfigWithServerId(this.Parent(), params.ServerId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Data["requestHeaderRef"] = webConfig.RequestHeaderPolicyRef
	this.Data["requestHeaderPolicy"] = webConfig.RequestHeaderPolicy
	this.Data["responseHeaderRef"] = webConfig.ResponseHeaderPolicyRef
	this.Data["responseHeaderPolicy"] = webConfig.ResponseHeaderPolicy

	this.Show()
}
