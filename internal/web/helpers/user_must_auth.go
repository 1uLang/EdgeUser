package helpers

import (
	"errors"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"net/http"
	"reflect"
	"time"

	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
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
			this.login(action)
			return false
		}
		//续期
		cache.Incr(fmt.Sprintf("login_success_userid_%v", userId), time.Minute*30)
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
			"icon": "shield",
			"url":  "/waf",
			"subItems": []maps.Map{
				{
					"name": "安全概览",
					"code": "waf",
					"url":  "/waf",
				},
				{
					"name": "拦截日志",
					"code": "wafLogs",
					"url":  "/waf/logs",
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
					"name": "证书管理",
					"code": "certs",
					"url":  "/servers/certs",
				},
				{
					"name": "刷新预热",
					"code": "cache",
					"url":  "/servers/cache",
				},
			},
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
			},
		},
		{
			"code": "webscan",
			"url":  "/webscan/targets",
			"name": "漏洞扫描",
			"icon": "ioxhost",
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
			"name": "堡垒机",
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
					"name": "子账号管理",
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
			},
		},
	}

	result := []maps.Map{}
	configloaders.LoadUIConfig()
	for _, m := range allMaps {

		//默认展示该组件
		if m.GetString("code") == "hids" || m.GetString("code") == "webscan" || m.GetString("code") == "fortcloud" {
			result = append(result, m)
			continue
		}

		//if m.GetString("code") == "finance" {
		//
		//	if config != nil && !config.ShowFinance {
		//		continue
		//	}
		//	if !lists.ContainsString(featureCodes, "finance") {
		//		continue
		//	}
		//}
		if m.GetString("code") == "lb" && !lists.ContainsString(featureCodes, "server.tcp") {
			continue
		}
		if m.GetString("code") == "waf" && !lists.ContainsString(featureCodes, "server.waf") {
			continue
		}
		if code := m.GetString("code"); (code == "hids" || code == "webscan") && !lists.ContainsString(featureCodes, code) {
			continue
		}
		result = append(result, m)
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
