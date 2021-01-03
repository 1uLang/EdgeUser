package ipadminutils

import (
	"context"
	"github.com/TeaOSLab/EdgeCommon/pkg/messageconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/nodeutils"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
)

// 通知使用此WAF策略的集群更新
func NotifyUpdateToClustersWithFirewallPolicyId(ctx context.Context, firewallPolicyId int64) error {
	client, err := rpc.SharedRPC()
	if err != nil {
		return err
	}
	resp, err := client.NodeClusterRPC().FindAllEnabledNodeClustersWithHTTPFirewallPolicyId(ctx, &pb.FindAllEnabledNodeClustersWithHTTPFirewallPolicyIdRequest{HttpFirewallPolicyId: firewallPolicyId})
	if err != nil {
		return err
	}
	for _, cluster := range resp.NodeClusters {
		_, err = nodeutils.SendMessageToCluster(ctx, cluster.Id, messageconfigs.MessageCodeIPListChanged, &messageconfigs.IPListChangedMessage{}, 3)
		if err != nil {
			return err
		}
	}
	return nil
}
