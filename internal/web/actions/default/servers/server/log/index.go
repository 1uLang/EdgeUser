package log

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "log", "")
	this.SecondMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	ServerId  int64
	RequestId string
}) {
	this.Data["featureIsOn"] = this.ValidateFeature("server.viewAccessLog")

	this.Data["serverId"] = params.ServerId
	this.Data["requestId"] = params.RequestId

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId  int64
	RequestId string

	Must *actions.Must
}) {
	if !this.ValidateFeature("server.viewAccessLog") {
		return
	}

	isReverse := len(params.RequestId) > 0
	accessLogsResp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.UserContext(), &pb.ListHTTPAccessLogsRequest{
		ServerId:  params.ServerId,
		RequestId: params.RequestId,
		Size:      20,
		Day:       timeutil.Format("Ymd"),
		Reverse:   isReverse,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	ipList := []string{}
	accessLogs := accessLogsResp.HttpAccessLogs
	if len(accessLogs) == 0 {
		accessLogs = []*pb.HTTPAccessLog{}
	} else {
		for _, accessLog := range accessLogs {
			if len(accessLog.RemoteAddr) > 0 {
				if !lists.ContainsString(ipList, accessLog.RemoteAddr) {
					ipList = append(ipList, accessLog.RemoteAddr)
				}
			}
		}
	}
	this.Data["accessLogs"] = accessLogs
	if len(accessLogs) > 0 {
		this.Data["requestId"] = accessLogs[0].RequestId
	} else {
		this.Data["requestId"] = params.RequestId
	}
	this.Data["hasMore"] = accessLogsResp.HasMore

	// 根据IP查询区域
	regionMap := map[string]string{} // ip => region
	if len(ipList) > 0 {
		resp, err := this.RPC().IPLibraryRPC().LookupIPRegions(this.UserContext(), &pb.LookupIPRegionsRequest{IpList: ipList})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if resp.IpRegionMap != nil {
			for ip, region := range resp.IpRegionMap {
				regionMap[ip] = region.Summary
			}
		}
	}
	this.Data["regions"] = regionMap

	this.Success()
}
