package basic

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("basic")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	resp, err := this.RPC().ServerRPC().FindEnabledUserServerBasic(this.UserContext(), &pb.FindEnabledUserServerBasicRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	server := resp.Server
	if server == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	this.Data["name"] = server.Name

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId int64
	Name     string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入服务名称")

	_, err := this.RPC().ServerRPC().UpdateEnabledUserServerBasic(this.UserContext(), &pb.UpdateEnabledUserServerBasicRequest{
		ServerId: params.ServerId,
		Name:     params.Name,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
