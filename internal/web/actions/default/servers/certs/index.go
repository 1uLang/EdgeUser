package certs

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	Type    string
	Keyword string
}) {
	this.Data["type"] = params.Type
	this.Data["keyword"] = params.Keyword

	countAll := int64(0)
	countCA := int64(0)
	countAvailable := int64(0)
	countExpired := int64(0)
	count7Days := int64(0)
	count30Days := int64(0)

	// 计算数量
	{
		// all
		resp, err := this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			Keyword: params.Keyword,
			UserId:  this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countAll = resp.Count

		// CA
		resp, err = this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			IsCA:    true,
			Keyword: params.Keyword,
			UserId:  this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countCA = resp.Count

		// available
		resp, err = this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			IsAvailable: true,
			Keyword:     params.Keyword,
			UserId:      this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countAvailable = resp.Count

		// expired
		resp, err = this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			IsExpired: true,
			Keyword:   params.Keyword,
			UserId:    this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countExpired = resp.Count

		// expire in 7 days
		resp, err = this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			ExpiringDays: 7,
			Keyword:      params.Keyword,
			UserId:       this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		count7Days = resp.Count

		// expire in 30 days
		resp, err = this.RPC().SSLCertRPC().CountSSLCerts(this.UserContext(), &pb.CountSSLCertRequest{
			ExpiringDays: 30,
			Keyword:      params.Keyword,
			UserId:       this.UserId(true),
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		count30Days = resp.Count
	}

	this.Data["countAll"] = countAll
	this.Data["countCA"] = countCA
	this.Data["countAvailable"] = countAvailable
	this.Data["countExpired"] = countExpired
	this.Data["count7Days"] = count7Days
	this.Data["count30Days"] = count30Days

	// 分页
	var page *actionutils.Page
	var listResp *pb.ListSSLCertsResponse
	var err error
	switch params.Type {
	case "":
		page = this.NewPage(countAll)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{
			Offset:  page.Offset,
			Size:    page.Size,
			Keyword: params.Keyword,
			UserId:  this.UserId(true),
		})
	case "ca":
		page = this.NewPage(countCA)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{
			IsCA:    true,
			Offset:  page.Offset,
			Size:    page.Size,
			Keyword: params.Keyword,
			UserId:  this.UserId(),
		})
	case "available":
		page = this.NewPage(countAvailable)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{IsAvailable: true, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword, UserId: this.UserId(true)})
	case "expired":
		page = this.NewPage(countExpired)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{IsExpired: true, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword, UserId: this.UserId(true)})
	case "7days":
		page = this.NewPage(count7Days)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{ExpiringDays: 7, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword, UserId: this.UserId(true)})
	case "30days":
		page = this.NewPage(count30Days)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{ExpiringDays: 30, Offset: page.Offset, Size: page.Size, Keyword: params.Keyword, UserId: this.UserId(true)})
	default:
		page = this.NewPage(countAll)
		listResp, err = this.RPC().SSLCertRPC().ListSSLCerts(this.UserContext(), &pb.ListSSLCertsRequest{
			Keyword: params.Keyword,
			UserId:  this.UserId(),
		})
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}

	certConfigs := []*sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(listResp.SslCertsJSON, &certConfigs)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["certs"] = certConfigs

	certMaps := []maps.Map{}
	nowTime := time.Now().Unix()
	for _, certConfig := range certConfigs {
		countServersResp, err := this.RPC().ServerRPC().CountAllEnabledServersWithSSLCertId(this.UserContext(), &pb.CountAllEnabledServersWithSSLCertIdRequest{SslCertId: certConfig.Id})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		certMaps = append(certMaps, maps.Map{
			"isOn":         certConfig.IsOn,
			"beginDay":     timeutil.FormatTime("Y-m-d", certConfig.TimeBeginAt),
			"endDay":       timeutil.FormatTime("Y-m-d", certConfig.TimeEndAt),
			"isExpired":    nowTime > certConfig.TimeEndAt,
			"isAvailable":  nowTime <= certConfig.TimeEndAt,
			"countServers": countServersResp.Count,
		})
	}
	this.Data["certInfos"] = certMaps

	this.Data["page"] = page.AsHTML()

	this.Show()
}
