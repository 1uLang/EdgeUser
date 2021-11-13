package https

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("https")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	serverConfig, err := dao.SharedServerDAO.FindEnabledServerConfig(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if serverConfig == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	httpsConfig := serverConfig.HTTPS
	if httpsConfig == nil {
		httpsConfig = &serverconfigs.HTTPSProtocolConfig{}
	}

	var sslPolicy *sslconfigs.SSLPolicy
	if httpsConfig.SSLPolicyRef != nil && httpsConfig.SSLPolicyRef.SSLPolicyId > 0 {
		sslPolicyConfigResp, err := this.RPC().SSLPolicyRPC().FindEnabledSSLPolicyConfig(this.UserContext(), &pb.FindEnabledSSLPolicyConfigRequest{SslPolicyId: httpsConfig.SSLPolicyRef.SSLPolicyId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		sslPolicyConfigJSON := sslPolicyConfigResp.SslPolicyJSON
		if len(sslPolicyConfigJSON) > 0 {
			sslPolicy = &sslconfigs.SSLPolicy{}
			err = json.Unmarshal(sslPolicyConfigJSON, sslPolicy)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
	}

	this.Data["serverType"] = serverConfig.Type
	this.Data["httpsConfig"] = maps.Map{
		"isOn":      httpsConfig.IsOn,
		"sslPolicy": sslPolicy,
	}

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId int64
	IsOn     bool

	SslPolicyJSON []byte

	Must *actions.Must
}) {
	// 记录日志
	defer this.CreateLog(oplogs.LevelInfo, "修改服务 %d 的HTTPS设置", params.ServerId)

	serverConfig, err := dao.SharedServerDAO.FindEnabledServerConfig(this.UserContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if serverConfig == nil {
		this.NotFound("server", params.ServerId)
		return
	}

	// 校验SSL
	var sslPolicyId = int64(0)
	if params.SslPolicyJSON != nil {
		sslPolicy := &sslconfigs.SSLPolicy{}
		err = json.Unmarshal(params.SslPolicyJSON, sslPolicy)
		if err != nil {
			this.ErrorPage(errors.New("解析SSL配置时发生了错误：" + err.Error()))
			return
		}

		sslPolicyId = sslPolicy.Id

		certsJSON, err := json.Marshal(sslPolicy.CertRefs)
		if err != nil {
			this.ErrorPage(err)
			return
		}

		hstsJSON, err := json.Marshal(sslPolicy.HSTS)
		if err != nil {
			this.ErrorPage(err)
			return
		}

		clientCACertsJSON, err := json.Marshal(sslPolicy.ClientCARefs)
		if err != nil {
			this.ErrorPage(err)
			return
		}

		if sslPolicyId > 0 {
			_, err := this.RPC().SSLPolicyRPC().UpdateSSLPolicy(this.UserContext(), &pb.UpdateSSLPolicyRequest{
				SslPolicyId:       sslPolicyId,
				Http2Enabled:      sslPolicy.HTTP2Enabled,
				MinVersion:        sslPolicy.MinVersion,
				SslCertsJSON:      certsJSON,
				HstsJSON:          hstsJSON,
				ClientAuthType:    types.Int32(sslPolicy.ClientAuthType),
				ClientCACertsJSON: clientCACertsJSON,
				CipherSuitesIsOn:  sslPolicy.CipherSuitesIsOn,
				CipherSuites:      sslPolicy.CipherSuites,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
		} else {
			resp, err := this.RPC().SSLPolicyRPC().CreateSSLPolicy(this.UserContext(), &pb.CreateSSLPolicyRequest{
				Http2Enabled:      sslPolicy.HTTP2Enabled,
				MinVersion:        sslPolicy.MinVersion,
				SslCertsJSON:      certsJSON,
				HstsJSON:          hstsJSON,
				ClientAuthType:    types.Int32(sslPolicy.ClientAuthType),
				ClientCACertsJSON: clientCACertsJSON,
				CipherSuitesIsOn:  sslPolicy.CipherSuitesIsOn,
				CipherSuites:      sslPolicy.CipherSuites,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			sslPolicyId = resp.SslPolicyId
		}
	}

	httpsConfig := serverConfig.HTTPS
	if httpsConfig == nil {
		httpsConfig = &serverconfigs.HTTPSProtocolConfig{}
	}

	httpsConfig.SSLPolicyRef = &sslconfigs.SSLPolicyRef{
		IsOn:        true,
		SSLPolicyId: sslPolicyId,
	}
	httpsConfig.IsOn = params.IsOn
	configData, err := json.Marshal(httpsConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	_, err = this.RPC().ServerRPC().UpdateServerHTTPS(this.UserContext(), &pb.UpdateServerHTTPSRequest{
		ServerId:  params.ServerId,
		HttpsJSON: configData,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
