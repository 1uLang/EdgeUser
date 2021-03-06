package certs

import (
	"archive/zip"
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"strconv"
)

type DownloadZipAction struct {
	actionutils.ParentAction
}

func (this *DownloadZipAction) Init() {
	this.Nav("", "", "")
}

func (this *DownloadZipAction) RunGet(params struct {
	CertId int64
}) {
	defer this.CreateLogInfo("下载SSL证书压缩包 %d", params.CertId)

	certResp, err := this.RPC().SSLCertRPC().FindEnabledSSLCertConfig(this.UserContext(), &pb.FindEnabledSSLCertConfigRequest{SslCertId: params.CertId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	certConfig := &sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(certResp.SslCertJSON, certConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	z := zip.NewWriter(this.ResponseWriter)
	defer func() {
		_ = z.Close()
	}()

	this.AddHeader("Content-Disposition", "attachment; filename=\"cert-"+strconv.FormatInt(params.CertId, 10)+".zip\";")

	// cert
	{
		w, err := z.Create("cert.pem")
		if err != nil {
			this.ErrorPage(err)
			return
		}
		_, err = w.Write(certConfig.CertData)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		err = z.Flush()
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// key
	if !certConfig.IsCA {
		w, err := z.Create("key.pem")
		if err != nil {
			this.ErrorPage(err)
			return
		}
		_, err = w.Write(certConfig.KeyData)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		err = z.Flush()
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
}
