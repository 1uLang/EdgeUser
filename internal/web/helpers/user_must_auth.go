package helpers

import (
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"net/http"
	"reflect"
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
			"name": "概览",
			"icon": "dashboard",
		},
		{
			"code": "servers",
			"name": "CDN加速",
			"icon": "clone outline",
			"subItems": []maps.Map{
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
				/**{
					"name": "用量统计",
					"code": "stat",
					"url":  "/servers/stat",
				},**/
			},
		},
		{
			"code": "lb",
			"name": "负载均衡",
			"icon": "clone outline",
		},
		{
			"code": "waf",
			"name": "WAF安全",
			"icon": "clone outline",
			"subItems": []maps.Map{
				{
					"name": "拦截日志",
					"code": "wafLogs",
					"url":  "/waf/logs",
				},
			},
		},
		{
			"code": "finance",
			"name": "费用账单",
			"icon": "yen sign",
		},
		{
			"code": "acl",
			"name": "访问控制",
			"icon": "address book",
		},
		/**{
			"code": "tickets",
			"name": "工单",
			"icon": "question circle outline",
		},**/
	}

	result := []maps.Map{}
	config, _ := configloaders.LoadUIConfig()
	for _, m := range allMaps {
		if m.GetString("code") == "finance" {
			if config != nil && !config.ShowFinance {
				continue
			}
			if !lists.ContainsString(featureCodes, "finance") {
				continue
			}
		}
		if m.GetString("code") == "lb" && !lists.ContainsString(featureCodes, "server.tcp") {
			continue
		}
		if m.GetString("code") == "waf" && !lists.ContainsString(featureCodes, "server.waf") {
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
