package waf

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	if !this.ValidateFeature("waf") {
		return
	}

	this.Data["serverId"] = params.ServerId
	this.Data["path"] = this.Request.URL.Path

	// 所有的服务列表
	serversResp, err := this.RPC().ServerRPC().ListEnabledServersMatch(this.UserContext(), &pb.ListEnabledServersMatchRequest{
		Offset:         0,
		Size:           100, // 我们这里最多显示前100个
		ServerGroupId:  0,
		Keyword:        "",
		ProtocolFamily: "http",
		UserId:         this.UserId(true),
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

	// 统计数据
	resp, err := this.RPC().ServerHTTPFirewallDailyStatRPC().ComposeServerHTTPFirewallDashboard(this.UserContext(), &pb.ComposeServerHTTPFirewallDashboardRequest{
		Day:      timeutil.Format("Ymd"),
		UserId:   this.UserId(true),
		ServerId: params.ServerId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["countDailyLog"] = resp.CountDailyLog
	this.Data["countDailyBlock"] = resp.CountDailyBlock
	this.Data["countDailyCaptcha"] = resp.CountDailyCaptcha
	this.Data["countWeeklyBlock"] = resp.CountWeeklyBlock
	this.Data["countMonthlyBlock"] = resp.CountMonthlyBlock

	// 分组
	groupStatMaps := []maps.Map{}
	for _, group := range resp.HttpFirewallRuleGroups {
		groupStatMaps = append(groupStatMaps, maps.Map{
			"group": maps.Map{
				"id":   group.HttpFirewallRuleGroup.Id,
				"name": group.HttpFirewallRuleGroup.Name,
			},
			"count": group.Count,
		})
	}
	this.Data["groupStats"] = groupStatMaps

	// 每日趋势
	logStatMaps := []maps.Map{}
	blockStatMaps := []maps.Map{}
	captchaStatMaps := []maps.Map{}
	for _, stat := range resp.LogDailyStats {
		logStatMaps = append(logStatMaps, maps.Map{
			"day":   stat.Day,
			"count": stat.Count,
		})
	}
	for _, stat := range resp.BlockDailyStats {
		blockStatMaps = append(blockStatMaps, maps.Map{
			"day":   stat.Day,
			"count": stat.Count,
		})
	}
	for _, stat := range resp.CaptchaDailyStats {
		captchaStatMaps = append(captchaStatMaps, maps.Map{
			"day":   stat.Day,
			"count": stat.Count,
		})
	}
	this.Data["logDailyStats"] = logStatMaps
	this.Data["blockDailyStats"] = blockStatMaps
	this.Data["captchaDailyStats"] = captchaStatMaps

	this.Show()
}
