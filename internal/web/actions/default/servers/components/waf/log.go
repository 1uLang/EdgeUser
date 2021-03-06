package waf

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"regexp"
	"strings"
)

type LogAction struct {
	actionutils.ParentAction
}

func (this *LogAction) Init() {
	this.Nav("", "", "log")
}

func (this *LogAction) RunGet(params struct {
	Day              string
	RequestId        string
	FirewallPolicyId int64
	GroupId          int64
}) {
	if len(params.Day) == 0 {
		params.Day = timeutil.Format("Y-m-d")
	}

	this.Data["path"] = this.Request.URL.Path
	this.Data["day"] = params.Day
	this.Data["groupId"] = params.GroupId
	this.Data["accessLogs"] = []interface{}{}

	day := params.Day
	if len(day) > 0 && regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(day) {
		day = strings.ReplaceAll(day, "-", "")
		size := int64(10)

		resp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.UserContext(), &pb.ListHTTPAccessLogsRequest{
			RequestId:           params.RequestId,
			FirewallPolicyId:    params.FirewallPolicyId,
			FirewallRuleGroupId: params.GroupId,
			Day:                 day,
			Size:                size,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}

		if len(resp.HttpAccessLogs) == 0 {
			this.Data["accessLogs"] = []interface{}{}
		} else {
			this.Data["accessLogs"] = resp.HttpAccessLogs
		}
		this.Data["hasMore"] = resp.HasMore
		this.Data["nextRequestId"] = resp.RequestId

		// 上一个requestId
		this.Data["hasPrev"] = false
		this.Data["lastRequestId"] = ""
		if len(params.RequestId) > 0 {
			this.Data["hasPrev"] = true
			prevResp, err := this.RPC().HTTPAccessLogRPC().ListHTTPAccessLogs(this.UserContext(), &pb.ListHTTPAccessLogsRequest{
				RequestId:           params.RequestId,
				FirewallPolicyId:    params.FirewallPolicyId,
				FirewallRuleGroupId: params.GroupId,
				Day:                 day,
				Size:                size,
				Reverse:             true,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			if int64(len(prevResp.HttpAccessLogs)) == size {
				this.Data["lastRequestId"] = prevResp.RequestId
			}
		}
	}

	// 所有分组
	policyResp, err := this.RPC().HTTPFirewallPolicyRPC().FindEnabledHTTPFirewallPolicyConfig(this.UserContext(), &pb.FindEnabledHTTPFirewallPolicyConfigRequest{
		HttpFirewallPolicyId: params.FirewallPolicyId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	policyConfig := &firewallconfigs.HTTPFirewallPolicy{}
	err = json.Unmarshal(policyResp.HttpFirewallPolicyJSON, policyConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	groupMaps := []maps.Map{}
	for _, group := range policyConfig.AllRuleGroups() {
		groupMaps = append(groupMaps, maps.Map{
			"id":   group.Id,
			"name": group.Name,
		})
	}
	this.Data["groups"] = groupMaps

	this.Show()
}
