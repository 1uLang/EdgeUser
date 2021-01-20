package logs

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"regexp"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Day              string
	RequestId        string
	FirewallPolicyId int64
	GroupId          int64
	ServerId         int64
}) {

	if len(params.Day) == 0 {
		params.Day = timeutil.Format("Y-m-d")
	}

	this.Data["path"] = this.Request.URL.Path
	this.Data["day"] = params.Day
	this.Data["groupId"] = params.GroupId
	this.Data["accessLogs"] = []interface{}{}
	this.Data["serverId"] = params.ServerId

	// 所有的服务列表
	serversResp, err := this.RPC().ServerRPC().ListEnabledServersMatch(this.UserContext(), &pb.ListEnabledServersMatchRequest{
		Offset:         0,
		Size:           100, // 我们这里最多显示前100个
		GroupId:        0,
		Keyword:        "",
		ProtocolFamily: "http",
		UserId:         this.UserId(),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverMaps := []maps.Map{}
	for _, server := range serversResp.Servers {
		if !server.IsOn {
			continue
		}

		// 域名列表
		serverNames := []*serverconfigs.ServerNameConfig{}
		if server.IsAuditing || (server.AuditingResult != nil && !server.AuditingResult.IsOk) {
			server.ServerNamesJSON = server.AuditingServerNamesJSON
		}
		if len(server.ServerNamesJSON) > 0 {
			err = json.Unmarshal(server.ServerNamesJSON, &serverNames)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
		if len(serverNames) == 0 {
			continue
		}

		serverName := server.Name
		if len(serverNames) > 0 {
			if len(serverNames[0].SubNames) == 0 {
				serverName = serverNames[0].Name
			} else {
				serverName = serverNames[0].SubNames[0]
			}
		}

		serverMaps = append(serverMaps, maps.Map{
			"id":         server.Id,
			"serverName": serverName,
		})
	}
	this.Data["servers"] = serverMaps

	// 查询
	day := params.Day
	ipList := []string{}
	if len(day) > 0 && regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(day) {
		day = strings.ReplaceAll(day, "-", "")
		size := int64(20)

		resp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.UserContext(), &pb.ListHTTPAccessLogsRequest{
			RequestId:           params.RequestId,
			UserId:              this.UserId(),
			ServerId:            params.ServerId,
			FirewallPolicyId:    params.FirewallPolicyId,
			FirewallRuleGroupId: params.GroupId,
			HasFirewallPolicy:   true,
			Day:                 day,
			Size:                size,
			Reverse:             false,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		if len(resp.AccessLogs) == 0 {
			this.Data["accessLogs"] = []interface{}{}
		} else {
			this.Data["accessLogs"] = resp.AccessLogs
			for _, accessLog := range resp.AccessLogs {
				if len(accessLog.RemoteAddr) > 0 {
					if !lists.ContainsString(ipList, accessLog.RemoteAddr) {
						ipList = append(ipList, accessLog.RemoteAddr)
					}
				}
			}
		}
		this.Data["hasMore"] = resp.HasMore
		this.Data["nextRequestId"] = resp.RequestId

		// 上一个requestId
		this.Data["hasPrev"] = false
		this.Data["lastRequestId"] = ""
		if len(params.RequestId) > 0 {
			this.Data["hasPrev"] = true
			prevResp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.UserContext(), &pb.ListHTTPAccessLogsRequest{
				UserId:              this.UserId(),
				ServerId:            params.ServerId,
				RequestId:           params.RequestId,
				FirewallPolicyId:    params.FirewallPolicyId,
				FirewallRuleGroupId: params.GroupId,
				HasFirewallPolicy:   true,
				Day:                 day,
				Size:                size,
				Reverse:             true,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			if int64(len(prevResp.AccessLogs)) == size {
				this.Data["lastRequestId"] = prevResp.RequestId
			}
		}
	}

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

	this.Show()
}
