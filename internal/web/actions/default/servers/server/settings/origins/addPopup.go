package origins

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"regexp"
	"strings"
)

// 添加源站
type AddPopupAction struct {
	actionutils.ParentAction
}

func (this *AddPopupAction) RunGet(params struct {
	ServerId       int64
	ReverseProxyId int64
	OriginType     string
}) {
	this.Data["reverseProxyId"] = params.ReverseProxyId
	this.Data["originType"] = params.OriginType

	serverTypeResp, err := this.RPC().ServerRPC().FindEnabledServerType(this.UserContext(), &pb.FindEnabledServerTypeRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverType := serverTypeResp.Type
	this.Data["serverType"] = serverType

	this.Show()
}

func (this *AddPopupAction) RunPost(params struct {
	OriginType string

	ReverseProxyId int64
	Weight         int32
	Protocol       string
	Addr           string
	Name           string
	Description    string
	IsOn           bool

	Must *actions.Must
}) {
	params.Must.
		Field("addr", params.Addr).
		Require("请输入源站地址")

	addr := regexp.MustCompile(`\s+`).ReplaceAllString(params.Addr, "")
	portIndex := strings.LastIndex(params.Addr, ":")
	if portIndex < 0 {
		this.Fail("地址中需要带有端口")
	}
	host := addr[:portIndex]
	port := addr[portIndex+1:]

	createResp, err := this.RPC().OriginRPC().CreateOrigin(this.UserContext(), &pb.CreateOriginRequest{
		Name: params.Name,
		Addr: &pb.NetworkAddress{
			Protocol:  params.Protocol,
			Host:      host,
			PortRange: port,
		},
		Description: params.Description,
		Weight:      params.Weight,
		IsOn:        params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	originId := createResp.OriginId
	originRef := &serverconfigs.OriginRef{
		IsOn:     true,
		OriginId: originId,
	}

	reverseProxyResp, err := this.RPC().ReverseProxyRPC().FindEnabledReverseProxy(this.UserContext(), &pb.FindEnabledReverseProxyRequest{ReverseProxyId: params.ReverseProxyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	reverseProxy := reverseProxyResp.ReverseProxy
	if reverseProxy == nil {
		this.ErrorPage(errors.New("reverse proxy should not be nil"))
		return
	}

	origins := []*serverconfigs.OriginRef{}
	switch params.OriginType {
	case "primary":
		if len(reverseProxy.PrimaryOriginsJSON) > 0 {
			err = json.Unmarshal(reverseProxy.PrimaryOriginsJSON, &origins)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
	case "backup":
		if len(reverseProxy.BackupOriginsJSON) > 0 {
			err = json.Unmarshal(reverseProxy.BackupOriginsJSON, &origins)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
	}
	origins = append(origins, originRef)
	originsData, err := json.Marshal(origins)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	switch params.OriginType {
	case "primary":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyPrimaryOrigins(this.UserContext(), &pb.UpdateReverseProxyPrimaryOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    originsData,
		})
	case "backup":
		_, err = this.RPC().ReverseProxyRPC().UpdateReverseProxyBackupOrigins(this.UserContext(), &pb.UpdateReverseProxyBackupOriginsRequest{
			ReverseProxyId: params.ReverseProxyId,
			OriginsJSON:    originsData,
		})
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 日志
	defer this.CreateLog(oplogs.LevelInfo, "为反向代理服务 %d 添加源站 %d", params.ReverseProxyId, originId)

	this.Success()
}
