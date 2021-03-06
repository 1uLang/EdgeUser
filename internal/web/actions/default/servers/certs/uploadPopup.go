package certs

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type UploadPopupAction struct {
	actionutils.ParentAction
}

func (this *UploadPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *UploadPopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *UploadPopupAction) RunPost(params struct {
	Name        string
	IsCA        bool
	Description string
	IsOn        bool

	CertFile *actions.File
	KeyFile  *actions.File

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入证书说明")

	certData := []byte{}
	keyData := []byte{}

	if params.CertFile == nil {
		this.Fail("请选择要上传的证书文件")
	}
	var err error
	certData, err = params.CertFile.Read()
	if err != nil {
		this.Fail("读取证书文件内容错误，请重新上传")
	}

	if !params.IsCA {
		if params.KeyFile == nil {
			this.Fail("请选择要上传的私钥文件")
		} else {
			keyData, err = params.KeyFile.Read()
			if err != nil {
				this.Fail("读取密钥文件内容错误，请重新上传")
			}
		}
	}

	// 校验
	sslConfig := &sslconfigs.SSLCertConfig{
		IsCA:     params.IsCA,
		CertData: certData,
		KeyData:  keyData,
	}
	err = sslConfig.Init()
	if err != nil {
		if params.IsCA {
			this.Fail("证书校验错误：" + err.Error())
		} else {
			this.Fail("证书或密钥校验错误：" + err.Error())
		}
	}

	// 保存
	createResp, err := this.RPC().SSLCertRPC().CreateSSLCert(this.UserContext(), &pb.CreateSSLCertRequest{
		IsOn:        params.IsOn,
		Name:        params.Name,
		Description: params.Description,
		ServerName:  "",
		IsCA:        params.IsCA,
		CertData:    certData,
		KeyData:     keyData,
		TimeBeginAt: sslConfig.TimeBeginAt,
		TimeEndAt:   sslConfig.TimeEndAt,
		DnsNames:    sslConfig.DNSNames,
		CommonNames: sslConfig.CommonNames,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 查询已创建的证书并返回，方便调用者进行后续处理
	certId := createResp.SslCertId
	configResp, err := this.RPC().SSLCertRPC().FindEnabledSSLCertConfig(this.UserContext(), &pb.FindEnabledSSLCertConfigRequest{SslCertId: certId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	certConfig := &sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(configResp.SslCertJSON, certConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	certConfig.CertData = nil // 去掉不必要的数据
	certConfig.KeyData = nil  // 去掉不必要的数据
	this.Data["cert"] = certConfig
	this.Data["certRef"] = &sslconfigs.SSLCertRef{
		IsOn:   true,
		CertId: certId,
	}

	// 创建日志
	defer this.CreateLog(oplogs.LevelInfo, "上传SSL证书 %d", certId)

	this.Success()
}
