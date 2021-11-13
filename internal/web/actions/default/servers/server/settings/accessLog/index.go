package accessLog

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("accessLog")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	//this.Data["featureIsOn"] = this.ValidateFeature("server.accessLog")
	this.Data["featureIsOn"] = true
	// 获取配置
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["webId"] = webConfig.Id
	this.Data["accessLogConfig"] = webConfig.AccessLogRef

	// 可选的缓存策略
	this.Data["accessLogPolicies"] = []maps.Map{}

	// 通用变量
	this.Data["fields"] = serverconfigs.HTTPAccessLogShortFields
	this.Data["defaultFieldCodes"] = serverconfigs.HTTPAccessLogDefaultFieldsCodes

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId         int64
	AccessLogJSON []byte

	Must *actions.Must
}) {
	//if !this.ValidateFeature("server.accessLog") {
	//	this.Fail("Permission denied")
	//	return
	//}

	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "修改Web %d 的访问日志设置", params.WebId)

	// TODO 检查参数

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebAccessLog(this.UserContext(), &pb.UpdateHTTPWebAccessLogRequest{
		WebId:         params.WebId,
		AccessLogJSON: params.AccessLogJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
