package dns

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("dns")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	dnsInfoResp, err := this.RPC().ServerRPC().FindEnabledServerDNS(this.UserContext(), &pb.FindEnabledServerDNSRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["dnsName"] = dnsInfoResp.DnsName
	if dnsInfoResp.Domain != nil {
		this.Data["dnsDomain"] = dnsInfoResp.Domain.Name
	} else {
		this.Data["dnsDomain"] = ""
	}

	this.Show()
}
