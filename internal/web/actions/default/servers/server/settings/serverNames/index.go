package serverNames

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

// 域名管理
type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	serverNamesResp, err := this.RPC().ServerRPC().FindServerNames(this.UserContext(), &pb.FindServerNamesRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	serverNamesConfig := []*serverconfigs.ServerNameConfig{}
	this.Data["isAuditing"] = serverNamesResp.IsAuditing
	this.Data["auditingResult"] = maps.Map{
		"isOk": true,
	}
	if serverNamesResp.IsAuditing {
		serverNamesResp.ServerNamesJSON = serverNamesResp.AuditingServerNamesJSON
	} else if serverNamesResp.AuditingResult != nil {
		if !serverNamesResp.AuditingResult.IsOk {
			serverNamesResp.ServerNamesJSON = serverNamesResp.AuditingServerNamesJSON
		}

		this.Data["auditingResult"] = maps.Map{
			"isOk":        serverNamesResp.AuditingResult.IsOk,
			"reason":      serverNamesResp.AuditingResult.Reason,
			"createdTime": timeutil.FormatTime("Y-m-d H:i:s", serverNamesResp.AuditingResult.CreatedAt),
		}
	}
	if len(serverNamesResp.ServerNamesJSON) > 0 {
		err := json.Unmarshal(serverNamesResp.ServerNamesJSON, &serverNamesConfig)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Data["serverNames"] = serverNamesConfig

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId    int64
	ServerNames string
	Must        *actions.Must
	CSRF        *actionutils.CSRF
}) {
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "修改代理服务 %d 域名", params.ServerId)

	serverNames := []*serverconfigs.ServerNameConfig{}
	err := json.Unmarshal([]byte(params.ServerNames), &serverNames)
	if err != nil {
		this.Fail("域名解析失败：" + err.Error())
	}

	_, err = this.RPC().ServerRPC().UpdateServerNames(this.UserContext(), &pb.UpdateServerNamesRequest{
		ServerId:        params.ServerId,
		ServerNamesJSON: []byte(params.ServerNames),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}