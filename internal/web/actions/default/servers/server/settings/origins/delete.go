package origins

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ReverseProxyId int64
	OriginId       int64
	OriginType     string
}) {
	reverseProxyResp, err := this.RPC().ReverseProxyRPC().FindEnabledReverseProxy(this.UserContext(), &pb.FindEnabledReverseProxyRequest{ReverseProxyId: params.ReverseProxyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	reverseProxy := reverseProxyResp.ReverseProxy
	if reverseProxy == nil {
		this.ErrorPage(errors.New("reverse proxy is nil"))
		return
	}

	origins := []*serverconfigs.OriginRef{}
	switch params.OriginType {
	case "primary":
		err = json.Unmarshal(reverseProxy.PrimaryOriginsJSON, &origins)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	case "backup":
		err = json.Unmarshal(reverseProxy.BackupOriginsJSON, &origins)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	default:
		this.ErrorPage(errors.New("invalid origin type '" + params.OriginType + "'"))
		return
	}

	result := []*serverconfigs.OriginRef{}
	for _, origin := range origins {
		if origin.OriginId == params.OriginId {
			continue
		}
		result = append(result, origin)
	}
	resultData, err := json.Marshal(result)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	switch params.OriginType {
	case "primary":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyPrimaryOrigins(this.UserContext(), &pb.UpdateReverseProxyPrimaryOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    resultData,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	case "backup":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyBackupOrigins(this.UserContext(), &pb.UpdateReverseProxyBackupOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    resultData,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "删除反向代理服务 %d 的源站 %d", params.ReverseProxyId, params.OriginId)

	this.Success()
}
