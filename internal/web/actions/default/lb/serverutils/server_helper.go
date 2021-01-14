package serverutils

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
	"net/http"
	"strconv"
)

type ServerHelper struct {
}

func NewServerHelper() *ServerHelper {
	return &ServerHelper{}
}

func (this *ServerHelper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["teaMenu"] = "lb"

	// 左侧菜单
	this.createLeftMenu(action)
}

func (this *ServerHelper) createLeftMenu(action *actions.ActionObject) {
	// 初始化
	if !action.Data.Has("leftMenuItemIsDisabled") {
		action.Data["leftMenuItemIsDisabled"] = false
	}
	action.Data["leftMenuItems"] = []maps.Map{}
	mainTab, _ := action.Data["mainTab"]
	secondMenuItem, _ := action.Data["secondMenuItem"]

	serverId := action.ParamInt64("serverId")
	if serverId == 0 {
		return
	}
	serverIdString := strconv.FormatInt(serverId, 10)
	action.Data["serverId"] = serverId

	// 读取server信息
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		logs.Error(err)
		return
	}

	userId := action.Context.GetInt64("userId")
	ctx := rpcClient.Context(userId)
	serverResp, err := rpcClient.ServerRPC().FindEnabledServer(ctx, &pb.FindEnabledServerRequest{ServerId: serverId})
	if err != nil {
		logs.Error(err)
		return
	}
	server := serverResp.Server
	if server == nil {
		logs.Error(errors.New("can not find the server"))
		return
	}

	// 服务管理
	serverConfig := &serverconfigs.ServerConfig{}
	err = json.Unmarshal(server.Config, serverConfig)
	if err != nil {
		logs.Error(err)
		return
	}

	// TABBAR
	selectedTabbar, _ := action.Data["mainTab"]
	tabbar := actionutils.NewTabbar()
	tabbar.Add("服务列表", "", "/lb", "", false)
	//tabbar.Add("看板", "", "/servers/server/board?serverId="+serverIdString, "dashboard", selectedTabbar == "board")
	//tabbar.Add("日志", "", "/servers/server/log?serverId="+serverIdString, "history", selectedTabbar == "log")
	//tabbar.Add("统计", "", "/servers/server/stat?serverId="+serverIdString, "chart area", selectedTabbar == "stat")
	tabbar.Add("设置", "", "/lb/server?serverId="+serverIdString, "setting", selectedTabbar == "setting")
	//tabbar.Add("删除", "", "/servers/server/delete?serverId="+serverIdString, "trash", selectedTabbar == "delete")

	actionutils.SetTabbar(action, tabbar)

	// 左侧操作子菜单
	switch types.String(mainTab) {
	case "board":
		action.Data["leftMenuItems"] = this.createBoardMenu(types.String(secondMenuItem), serverIdString, serverConfig)
	case "log":
		action.Data["leftMenuItems"] = this.createLogMenu(types.String(secondMenuItem), serverIdString, serverConfig)
	case "stat":
		action.Data["leftMenuItems"] = this.createStatMenu(types.String(secondMenuItem), serverIdString, serverConfig)
	case "setting":
		action.Data["leftMenuItems"] = this.createSettingsMenu(types.String(secondMenuItem), serverIdString, serverConfig)
	case "delete":
		action.Data["leftMenuItems"] = this.createDeleteMenu(types.String(secondMenuItem), serverIdString, serverConfig)
	}
}

// 看板菜单
func (this *ServerHelper) createBoardMenu(secondMenuItem string, serverIdString string, serverConfig *serverconfigs.ServerConfig) []maps.Map {
	menuItems := []maps.Map{}
	menuItems = append(menuItems, maps.Map{
		"name":     "看板",
		"url":      "/servers/server/board?serverId=" + serverIdString,
		"isActive": secondMenuItem == "index",
	})
	return menuItems
}

// 日志菜单
func (this *ServerHelper) createLogMenu(secondMenuItem string, serverIdString string, serverConfig *serverconfigs.ServerConfig) []maps.Map {
	menuItems := []maps.Map{}
	menuItems = append(menuItems, maps.Map{
		"name":     "实时",
		"url":      "/servers/server/log?serverId=" + serverIdString,
		"isActive": secondMenuItem == "index",
	})
	menuItems = append(menuItems, maps.Map{
		"name":     "今天",
		"url":      "/servers/server/log/today?serverId=" + serverIdString,
		"isActive": secondMenuItem == "today",
	})
	menuItems = append(menuItems, maps.Map{
		"name":     "历史",
		"url":      "/servers/server/log/history?serverId=" + serverIdString,
		"isActive": secondMenuItem == "history",
	})
	return menuItems
}

// 统计菜单
func (this *ServerHelper) createStatMenu(secondMenuItem string, serverIdString string, serverConfig *serverconfigs.ServerConfig) []maps.Map {
	menuItems := []maps.Map{}
	menuItems = append(menuItems, maps.Map{
		"name":     "统计",
		"url":      "/servers/server/stat?serverId=" + serverIdString,
		"isActive": secondMenuItem == "index",
	})
	return menuItems
}

// 设置菜单
func (this *ServerHelper) createSettingsMenu(secondMenuItem string, serverIdString string, serverConfig *serverconfigs.ServerConfig) (items []maps.Map) {
	menuItems := []maps.Map{
		{
			"name":     "基本信息",
			"url":      "/lb/server/settings/basic?serverId=" + serverIdString,
			"isActive": secondMenuItem == "basic",
			"isOff":    !serverConfig.IsOn,
		},
		{
			"name":     "DNS",
			"url":      "/lb/server/settings/dns?serverId=" + serverIdString,
			"isActive": secondMenuItem == "dns",
		},
	}

	if serverConfig.IsTCP() {
		/**menuItems = append(menuItems, maps.Map{
			"name":     "TCP",
			"url":      "/lb/server/settings/tcp?serverId=" + serverIdString,
			"isActive": secondMenuItem == "tcp",
			"isOn":     serverConfig.TCP != nil && serverConfig.TCP.IsOn && len(serverConfig.TCP.Listen) > 0,
		})
		menuItems = append(menuItems, maps.Map{
			"name":     "TLS",
			"url":      "/lb/server/settings/tls?serverId=" + serverIdString,
			"isActive": secondMenuItem == "tls",
			"isOn":     serverConfig.TLS != nil && serverConfig.TLS.IsOn && len(serverConfig.TLS.Listen) > 0,
		})**/
		menuItems = append(menuItems, maps.Map{
			"name":     "源站",
			"url":      "/lb/server/settings/reverseProxy?serverId=" + serverIdString,
			"isActive": secondMenuItem == "reverseProxy",
			"isOn":     serverConfig.ReverseProxyRef != nil && serverConfig.ReverseProxyRef.IsOn,
		})
	} else if serverConfig.IsUnix() {
		menuItems = append(menuItems, maps.Map{
			"name":     "Unix",
			"url":      "/lb/server/settings/unix?serverId=" + serverIdString,
			"isActive": secondMenuItem == "unix",
			"isOn":     serverConfig.Unix != nil && serverConfig.Unix.IsOn && len(serverConfig.Unix.Listen) > 0,
		})
	} else if serverConfig.IsUDP() {
		menuItems = append(menuItems, maps.Map{
			"name":     "UDP",
			"url":      "/lb/server/settings/udp?serverId=" + serverIdString,
			"isActive": secondMenuItem == "udp",
			"isOn":     serverConfig.UDP != nil && serverConfig.UDP.IsOn && len(serverConfig.UDP.Listen) > 0,
		})
	}

	return menuItems
}

// 删除菜单
func (this *ServerHelper) createDeleteMenu(secondMenuItem string, serverIdString string, serverConfig *serverconfigs.ServerConfig) []maps.Map {
	menuItems := []maps.Map{}
	menuItems = append(menuItems, maps.Map{
		"name":     "删除",
		"url":      "/servers/server/delete?serverId=" + serverIdString,
		"isActive": secondMenuItem == "index",
	})
	return menuItems
}
