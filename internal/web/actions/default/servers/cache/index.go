package cache

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/messageconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/nodeutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"net/url"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "purge")
}

func (this *IndexAction) RunGet(params struct{}) {
	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	Type    string
	UrlList string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("批量刷新缓存")

	switch params.Type {
	case "file", "dir":
	default:
		this.Fail("请选择正确的刷新类型")
	}

	// 查找当前用户的所有域名
	serverNamesResp, err := this.RPC().ServerRPC().FindAllEnabledServerNamesWithUserId(this.UserContext(), &pb.FindAllEnabledServerNamesWithUserIdRequest{UserId: this.UserId(true)})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverNames := serverNamesResp.ServerNames

	keys := []string{}
	for _, key := range strings.Split(params.UrlList, "\n") {
		key = strings.TrimSpace(key)
		if len(key) == 0 {
			continue
		}
		if lists.ContainsString(keys, key) {
			continue
		}

		// 检查域名
		u, err := url.Parse(key)
		if err != nil || len(u.Host) == 0 || (u.Scheme != "http" && u.Scheme != "https") {
			this.Fail("'" + key + "'不是正确的URL格式")
		}
		if !configutils.MatchDomains(serverNames, u.Host) {
			this.Fail("'" + key + "'中域名'" + u.Host + "'没有绑定")
		}

		if params.Type == "dir" && !strings.HasSuffix(key, "/") {
			key += "/"
		}

		keys = append(keys, key)
	}

	if len(keys) == 0 {
		if params.Type == "file" {
			this.Fail("请输入要刷新的URL列表")
		} else if params.Type == "dir" {
			this.Fail("请输入要刷新的目录列表")
		}
	}

	// 当前用户所在集群
	clusterIdResp, err := this.RPC().UserRPC().FindUserNodeClusterId(this.UserContext(), &pb.FindUserNodeClusterIdRequest{UserId: this.UserId(true)})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	clusterId := clusterIdResp.NodeClusterId
	if clusterId <= 0 {
		this.Fail("当前用户尚未分配集群，不能执行此操作")
	}

	// 缓存策略
	clusterResp, err := this.RPC().NodeClusterRPC().FindEnabledNodeCluster(this.UserContext(), &pb.FindEnabledNodeClusterRequest{NodeClusterId: clusterId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	cluster := clusterResp.NodeCluster
	if cluster == nil {
		this.Fail("找不到对应的集群，请联系管理员操作")
	}
	if cluster.HttpCachePolicyId <= 0 {
		this.Fail("当前用户所在集群没有开启缓存策略，不能执行此操作")
	}
	cachePolicyResp, err := this.RPC().HTTPCachePolicyRPC().FindEnabledHTTPCachePolicyConfig(this.UserContext(), &pb.FindEnabledHTTPCachePolicyConfigRequest{HttpCachePolicyId: cluster.HttpCachePolicyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(cachePolicyResp.HttpCachePolicyJSON) == 0 {
		this.Fail("当前用户所在集群没有开启缓存策略，不能执行此操作")
	}

	// 发送命令
	msg := &messageconfigs.PurgeCacheMessage{
		CachePolicyJSON: cachePolicyResp.HttpCachePolicyJSON,
		Keys:            keys,
		Type:            params.Type,
	}
	results, err := nodeutils.SendMessageToCluster(this.UserContext(), clusterId, messageconfigs.MessageCodePurgeCache, msg, 60)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, result := range results {
		if !result.IsOK {
			this.Fail("操作失败，请联系管理员：节点: " + result.NodeName + ": " + result.Message)
		}
	}

	this.Success()
}
