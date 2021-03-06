package headers

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteDeletingHeaderAction struct {
	actionutils.ParentAction
}

func (this *DeleteDeletingHeaderAction) RunPost(params struct {
	HeaderPolicyId int64
	HeaderName     string
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "删除需要删除的请求Header，HeaderPolicyId:%d, HeaderName:%s", params.HeaderPolicyId, params.HeaderName)

	policyConfigResp, err := this.RPC().HTTPHeaderPolicyRPC().FindEnabledHTTPHeaderPolicyConfig(this.UserContext(), &pb.FindEnabledHTTPHeaderPolicyConfigRequest{HeaderPolicyId: params.HeaderPolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	policyConfigJSON := policyConfigResp.HeaderPolicyJSON
	policyConfig := &shared.HTTPHeaderPolicy{}
	err = json.Unmarshal(policyConfigJSON, policyConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	headerNames := []string{}
	for _, h := range policyConfig.DeleteHeaders {
		if h == params.HeaderName {
			continue
		}
		headerNames = append(headerNames, h)
	}
	_, err = this.RPC().HTTPHeaderPolicyRPC().UpdateHTTPHeaderPolicyDeletingHeaders(this.UserContext(), &pb.UpdateHTTPHeaderPolicyDeletingHeadersRequest{
		HeaderPolicyId: params.HeaderPolicyId,
		HeaderNames:    headerNames,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
