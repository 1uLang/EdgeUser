package http

import (
	"encoding/json"
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
	this.SecondMenu("http")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	serverConfig, err := dao.SharedServerDAO.FindServerConfig(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if serverConfig == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	httpConfig := serverConfig.HTTP
	if httpConfig == nil {
		httpConfig = &serverconfigs.HTTPProtocolConfig{}
		httpConfig.IsOn = true
	}

	this.Data["serverType"] = serverConfig.Type
	this.Data["httpConfig"] = maps.Map{
		"isOn":      httpConfig.IsOn,
		"addresses": httpConfig.Listen,
	}

	// 跳转相关设置
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["webId"] = webConfig.Id
	this.Data["redirectToHTTPSConfig"] = webConfig.RedirectToHttps

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId  int64
	IsOn      bool
	Addresses string

	WebId               int64
	RedirectToHTTPSJSON []byte

	Must *actions.Must
}) {
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "修改服务 %d 的HTTP设置", params.ServerId)

	serverConfig, err := dao.SharedServerDAO.FindServerConfig(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if serverConfig == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	httpConfig := serverConfig.HTTP
	if httpConfig == nil {
		httpConfig = &serverconfigs.HTTPProtocolConfig{}
	}
	httpConfig.IsOn = params.IsOn
	configData, err := json.Marshal(httpConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().ServerRPC().UpdateServerHTTP(this.UserContext(), &pb.UpdateServerHTTPRequest{
		ServerId: params.ServerId,
		HttpJSON: configData,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 设置跳转到HTTPS
	// TODO 校验设置
	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebRedirectToHTTPS(this.UserContext(), &pb.UpdateHTTPWebRedirectToHTTPSRequest{
		WebId:               params.WebId,
		RedirectToHTTPSJSON: params.RedirectToHTTPSJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
