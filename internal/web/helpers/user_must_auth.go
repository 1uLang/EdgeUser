package helpers

import (
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/iwind/TeaGo/actions"
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
	action.Data["teaUsername"] = this.findUserFullname(userId)

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
	allMaps := []maps.Map{
		{
			"code": "servers",
			"name": "网站服务",
			"icon": "clone outsize",
		},
	}

	return allMaps
}

// 跳转到登录页
func (this *userMustAuth) login(action *actions.ActionObject) {
	action.RedirectURL("/")
}

// 查找用户名称
func (this *userMustAuth) findUserFullname(userId int64) string {
	// TODO
	return ""
}
