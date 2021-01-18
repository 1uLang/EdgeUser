package waf

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct{}) {
	countResp, err := this.RPC().HTTPFirewallPolicyRPC().CountAllEnabledHTTPFirewallPolicies(this.UserContext(), &pb.CountAllEnabledHTTPFirewallPoliciesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)

	listResp, err := this.RPC().HTTPFirewallPolicyRPC().ListEnabledHTTPFirewallPolicies(this.UserContext(), &pb.ListEnabledHTTPFirewallPoliciesRequest{
		Offset: page.Offset,
		Size:   page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	policyMaps := []maps.Map{}
	for _, policy := range listResp.HttpFirewallPolicies {
		countInbound := 0
		countOutbound := 0
		if len(policy.InboundJSON) > 0 {
			inboundConfig := &firewallconfigs.HTTPFirewallInboundConfig{}
			err = json.Unmarshal(policy.InboundJSON, inboundConfig)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			countInbound = len(inboundConfig.GroupRefs)
		}
		if len(policy.OutboundJSON) > 0 {
			outboundConfig := &firewallconfigs.HTTPFirewallInboundConfig{}
			err = json.Unmarshal(policy.OutboundJSON, outboundConfig)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			countOutbound = len(outboundConfig.GroupRefs)
		}

		countClustersResp, err := this.RPC().NodeClusterRPC().CountAllEnabledNodeClustersWithHTTPFirewallPolicyId(this.UserContext(), &pb.CountAllEnabledNodeClustersWithHTTPFirewallPolicyIdRequest{HttpFirewallPolicyId: policy.Id})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		countClusters := countClustersResp.Count

		policyMaps = append(policyMaps, maps.Map{
			"id":            policy.Id,
			"isOn":          policy.IsOn,
			"name":          policy.Name,
			"countInbound":  countInbound,
			"countOutbound": countOutbound,
			"countClusters": countClusters,
		})
	}

	this.Data["policies"] = policyMaps

	this.Data["page"] = page.AsHTML()

	this.Show()
}
