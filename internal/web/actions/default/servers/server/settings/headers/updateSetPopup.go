package headers

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type UpdateSetPopupAction struct {
	actionutils.ParentAction
}

func (this *UpdateSetPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateSetPopupAction) RunGet(params struct {
	HeaderPolicyId int64
	HeaderId       int64
}) {
	this.Data["headerPolicyId"] = params.HeaderPolicyId
	this.Data["headerId"] = params.HeaderId

	headerResp, err := this.RPC().HTTPHeaderRPC().FindEnabledHTTPHeaderConfig(this.UserContext(), &pb.FindEnabledHTTPHeaderConfigRequest{HeaderId: params.HeaderId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	headerConfig := &shared.HTTPHeaderConfig{}
	err = json.Unmarshal(headerResp.HeaderJSON, headerConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["headerConfig"] = headerConfig

	this.Show()
}

func (this *UpdateSetPopupAction) RunPost(params struct {
	HeaderId int64
	Name     string
	Value    string

	Must *actions.Must
}) {
	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "修改设置请求Header，HeaderPolicyId:%d, Name:%s, Value:%s", params.HeaderId, params.Name, params.Value)

	params.Must.
		Field("name", params.Name).
		Require("请输入Header名称")

	_, err := this.RPC().HTTPHeaderRPC().UpdateHTTPHeader(this.UserContext(), &pb.UpdateHTTPHeaderRequest{
		HeaderId: params.HeaderId,
		Name:     params.Name,
		Value:    params.Value,
	})
	if err != nil {
		this.ErrorPage(err)
	}

	this.Success()
}
