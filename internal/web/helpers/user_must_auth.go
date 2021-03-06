package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/server/channels_server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// 认证拦截
type userMustAuth struct {
	UserId int64
	module string
}

func NewUserMustAuth(module string) *userMustAuth {
	return &userMustAuth{module: module}
}

func (this *userMustAuth) BeforeAction(actionPtr actions.ActionWrapper, paramName string) (goNext bool) {
	var action = actionPtr.Object()

	// 安全相关
	action.AddHeader("X-Frame-Options", "SAMEORIGIN")
	action.AddHeader("Content-Security-Policy", "default-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'")

	var session = action.Session()
	var userId = session.GetInt64("userId")

	if userId <= 0 {
		this.login(action)
		return false
	}

	// 检查用户是否存在
	// TODO

	this.UserId = userId
	action.Context.Set("userId", this.UserId)

	if action.Request.Method != http.MethodGet {
		return true
	}

	config, err := configloaders.LoadUIConfig()
	if err != nil {
		action.WriteString(err.Error())
		return false
	}

	// 初始化内置方法
	action.ViewFunc("teaTitle", func() string {
		return action.Data["teaTitle"].(string)
	})

	action.Data["teaShowVersion"] = config.ShowVersion
	action.Data["teaTitle"] = config.UserSystemName
	action.Data["teaName"] = config.ProductName
	action.Data["teaFaviconFileId"] = config.FaviconFileId
	action.Data["teaLogoFileId"] = config.LogoFileId
	action.Data["teaUsername"], err = this.findUserFullname(userId)
	if err != nil {
		logs.Println("[USER_MUST_AUTH]" + err.Error())
	}

	action.Data["teaUserAvatar"] = ""

	if !action.Data.Has("teaMenu") {
		action.Data["teaMenu"] = ""
	}
	action.Data["teaModules"] = this.modules(userId)
	action.Data["teaSubMenus"] = []map[string]interface{}{}
	action.Data["teaTabbar"] = []map[string]interface{}{}
	if len(config.Version) == 0 {
		action.Data["teaVersion"] = teaconst.Version
	} else {
		action.Data["teaVersion"] = config.Version
	}
	action.Data["teaShowOpenSourceInfo"] = config.ShowOpenSourceInfo
	action.Data["teaIsSuper"] = false
	action.Data["teaDemoEnabled"] = false
	if !action.Data.Has("teaSubMenu") {
		action.Data["teaSubMenu"] = ""
	}

	// 菜单
	action.Data["firstMenuItem"] = ""

	// 未读消息数
	action.Data["teaBadge"] = 0

	// 调用Init
	initMethod := reflect.ValueOf(actionPtr).MethodByName("Init")
	if initMethod.IsValid() {
		initMethod.Call([]reflect.Value{})
	}
	//fmt.Println("每次都执行的事件url=", action.Request.RequestURI, userId)
	if !utils.UrlIn(action.Request.RequestURI) {
		res, _ := cache.GetInt(fmt.Sprintf("login_success_userid_%v", userId))
		if res == 0 {
			//30分钟没有操作  自动退出
			session.Delete()
			cache.DelKey(fmt.Sprintf("login_success_userid_%v", userId))
			this.login(action)
			return false
		}
		//续期
		cache.Incr(fmt.Sprintf("login_success_userid_%v", userId), time.Minute*30)
	}

	//按照用户渠道信息 读取logo配置
	key := fmt.Sprintf("get-channel-info-uid-%v", userId)
	channelInfo, err := cache.CheckCache(key, func() (interface{}, error) {
		info, err := channels_server.GetInfoByUid(uint64(userId))
		return info, err
	}, 10, true) //10s缓存
	if err == nil {
		b, err := json.Marshal(channelInfo)
		if err == nil {
			list := &channels.Channels{}
			err = json.Unmarshal(b, &list)
			if err == nil {
				if list.ProductName != "" {
					action.Data["teaTitle"] = list.ProductName
				}
				if list.Logo != "" {
					action.Data["teaLogoFileId"] = list.Logo
					action.Data["teaFaviconFileId"], _ = strconv.Atoi(list.Logo)
				}
			}

		}
	}
	return true
}

// 菜单配置
func (this *userMustAuth) modules(userId int64) []maps.Map {
	// 开通的功能
	featureCodes := []string{}
	rpcClient, err := rpc.SharedRPC()
	if err == nil {
		userFeatureResp, err := rpcClient.UserRPC().FindUserFeatures(rpcClient.Context(userId), &pb.FindUserFeaturesRequest{UserId: userId})
		if err == nil {
			for _, feature := range userFeatureResp.Features {
				featureCodes = append(featureCodes, feature.Code)
			}
		}
	}
	allMaps := []maps.Map{
		{
			"code": "dashboard",
			"name": "业务概览",
			"icon": "dashboard",
			"url":  "/dashboard",
		},
		{
			"code": "waf",
			"name": "态势感知",
			"icon": "eye",
			"url":  "/waf",
			"subItems": []maps.Map{
				{
					"name": "安全概览",
					"code": "waf",
					"url":  "/waf",
				},
				{
					"name": "DDoS日志",
					"code": "ddos",
					"url":  "/waf/ddos",
				},
				{
					"name": "IPS日志",
					"code": "alarm",
					"url":  "/waf/alarm",
				},
				{
					"name": "APT日志",
					"code": "apt",
					"url":  "/waf/apt",
				},
				{
					"name": "WAF日志",
					"code": "logs",
					"url":  "/waf/logs",
				},
			},
		},
		//{
		//	"code": "hostlist",
		//	"url":  "/hostlist",
		//	"name": "云主机管理",
		//	"icon": "tv",
		//},
		{
			"code": "ddos",
			"name": "DDoS防护",
			"icon": "shield",
			"url":  "/ddos",
			"subItems": []maps.Map{
				{
					"name": "主机状态",
					"code": "host",
					"url":  "/ddos/host",
				},
				{
					"name": "连接监控",
					"code": "link",
					"url":  "/ddos/link",
				},
				{
					"name": "屏蔽列表",
					"code": "shield",
					"url":  "/ddos/shield",
				},
			},
		},
		{
			"code": "nfw",
			"name": "云防火墙",
			"icon": "bars",
			"url":  "/nfw",
			"subItems": []maps.Map{
				{
					"name": "ACL规则",
					"code": "acl",
					"url":  "/nfw/acl",
				},
				{
					"name": "IPS规则",
					"code": "ips",
					"url":  "/nfw/ips",
				},
				//{
				//	"name": "病毒库",
				//	"code": "virus",
				//	"url":  "/nfw/virus",
				//},
				{
					"name": "会话列表",
					"code": "conversation",
					"url":  "/nfw/conversation",
				},
			},
		},
		{
			"code": "servers",
			"name": "WAF服务",
			"url":  "/servers",
			"icon": "skyatlas",
			"subItems": []maps.Map{
				{
					"name": "域名管理",
					"code": "servers",
					"url":  "/servers",
				},
				{
					"name": "刷新预热",
					"code": "cache",
					"url":  "/servers/cache",
				},
			},
		},
		{
			"code": "certs",
			"name": "证书服务",
			"icon": "leanpub",
			"url":  "/servers/certs",
		},
		{
			"code": "lb",
			"name": "负载均衡",
			"icon": "paper plane",
			"url":  "/lb",
		},
		{
			"code": "hids",
			"url":  "/hids/examine",
			"name": "主机防护",
			"icon": "linux",
			"subItems": []maps.Map{
				{
					"name": "主机体检",
					"url":  "/hids/examine",
					"code": "examine",
				},
				{
					"name": "漏洞风险",
					"url":  "/hids/risk",
					"code": "risk",
				},
				{
					"name": "入侵威胁",
					"url":  "/hids/invade",
					"code": "invade",
				},
				{
					"name": "合规基线",
					"url":  "/hids/baseline",
					"code": "baseline",
				},
				{
					"name": "Agent管理",
					"url":  "/hids/agent",
					"code": "agent",
				},
				//{
				//	"name": "黑白名单",
				//	"url":  "/hids/bwlist",
				//	"code": "bwlist",
				//},
			},
		},
		{
			"code": "nhids",
			"url":  "/hids/agents",
			"name": "端点防护",
			"icon": "laptop",
			"subItems": []maps.Map{
				{
					"name": "资产管理",
					"url":  "/hids/agents",
					"code": "agents",
				},
				{
					"name": "安全事件",
					"url":  "/hids/attck",
					"code": "attck",
				},
				//{
				//	"name": "漏洞风险",
				//	"url":  "/hids/vulnerability",
				//	"code": "vulnerability",
				//},
				{
					"name": "病毒查杀",
					"url":  "/hids/virus",
					"code": "virus",
				},
				{
					"name": "合规基线",
					"url":  "/hids/baseLine",
					"code": "baseLine",
				},
				{
					"name": "文件监控",
					"url":  "/hids/syscheck",
					"code": "syscheck",
				},
			},
		},
		{
			"code": "webscan",
			"url":  "/webscan/targets",
			"name": "漏洞扫描",
			"icon": "yelp",
			"subItems": []maps.Map{
				{
					"name": "扫描目标",
					"url":  "/webscan/targets",
					"code": "targets",
				},
				{
					"name": "扫描任务",
					"url":  "/webscan/scans",
					"code": "scans",
				},
				{
					"name": "扫描报告",
					"url":  "/webscan/reports",
					"code": "reports",
				},
			},
		},
		{
			"code": "fortcloud",
			"url":  "/fortcloud/assets",
			"name": "云堡垒机",
			"icon": "ioxhost",
			"subItems": []maps.Map{
				{
					"name": "资产管理",
					"url":  "/fortcloud/assets",
					"code": "assets",
				},
				{
					"name": "授权凭证",
					"url":  "/fortcloud/cert",
					"code": "cert",
				},
				{
					"name": "会话管理",
					"url":  "/fortcloud/sessions",
					"code": "sessions",
				},
				{
					"name": "运维审计",
					"url":  "/fortcloud/audit",
					"code": "audit",
				},
			},
		},
		//{
		//	"code": "finance",
		//	"name": "费用账单",
		//	"icon": "yen sign",
		//},
		//{
		//	"code": "acl",
		//	"name": "访问控制",
		//	"icon": "address book",
		//},
		/**{
			"code": "tickets",
			"name": "工单",
			"icon": "question circle outline",
		},**/
		{
			"code": "audit",
			"url":  "/audit/db",
			"name": "安全审计",
			"icon": "sellsy",
			"subItems": []maps.Map{
				{
					"name": "数据库管理",
					"url":  "/audit/db",
					"code": "db",
				},
				{
					"name": "主机管理",
					"url":  "/audit/host",
					"code": "host",
				},
				{
					"name": "应用管理",
					"url":  "/audit/app",
					"code": "app",
				},
				{
					"name": "审计日志",
					"url":  "/audit/logs",
					"code": "logs",
				},
				{
					"name": "订阅报告",
					"url":  "/audit/report",
					"code": "report",
				},
				{
					"name": "Agent管理",
					"url":  "/audit/agent",
					"code": "agent",
				},
			},
		},
		{
			"code": "databackup",
			"url":  "/databackup",
			"name": "数据备份",
			"icon": "copy",
			//"subItems": []maps.Map{
			//	{
			//		"name": "文件管理",
			//		"url":  "/databackup",
			//		"code": "assets",
			//	},
			//	{
			//		"name": "上传文件",
			//		"url":  "/databackup/createPopUp",
			//		"code": "admins",
			//	},
			//},
		},
		{
			"code": "platform",
			"url":  "/platform/user",
			"name": "平台管理",
			"icon": "sitemap",
			"subItems": []maps.Map{
				{
					"name": "账号管理",
					"url":  "/platform/user",
					"code": "user",
				},
				{
					"name": "操作日志",
					"url":  "/platform/logs",
					"code": "logs",
				},
				{
					"name": "安全策略",
					"url":  "/platform/strategy",
					"code": "strategy",
				},
				{
					"name": "特征库管理",
					"url":  "/platform/feature_library/virus",
					"code": "virus",
				},
			},
		},
	}

	result := []maps.Map{}
	configloaders.LoadUIConfig()

	for _, m := range allMaps {

		//默认展示该组件
		//if m.GetString("code") == "hids" || m.GetString("code") == "webscan" || m.GetString("code") == "fortcloud" {
		//	result = append(result, m)
		//	continue
		//}
		code := m.GetString("code")
		if lists.ContainsString(featureCodes, code) || (code == "lb" && (lists.Contains(featureCodes, "lb-tcp") || lists.Contains(featureCodes, "lb-tcp-port"))) {
			result = append(result, m)
		} else { //判断子菜单是否已授权
			sub := m.GetSlice("subItems")
			newSub := []maps.Map{}
			if sub != nil {
				for _, item := range sub {
					sub := item.(maps.Map)
					subCode := sub.GetString("code")
					if lists.ContainsString(featureCodes, code+"."+subCode) { //表示子菜单包含
						newSub = append(newSub, sub)
					}
				}
			}
			if len(newSub) > 0 {
				m["subItems"] = newSub
				result = append(result, m)
			}
		}
	}

	return result
}

// 跳转到登录页
func (this *userMustAuth) login(action *actions.ActionObject) {
	action.RedirectURL("/")
}

// 查找用户名称
func (this *userMustAuth) findUserFullname(userId int64) (string, error) {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return "", err
	}
	resp, err := rpcClient.UserRPC().FindEnabledUser(rpcClient.Context(userId), &pb.FindEnabledUserRequest{UserId: userId})
	if err != nil {
		return "", err
	}
	if resp.User == nil {
		return "", errors.New("can not find user")
	}

	return resp.User.Fullname, nil
}

func (this *userMustAuth) FirstMenuUrl(userId int64) string {
	menus := this.modules(userId)

	if sub := menus[0].GetSlice("subItems"); sub != nil {
		return sub[0].(maps.Map).GetString("url")
	} else {
		return menus[0].GetString("url")
	}
}
