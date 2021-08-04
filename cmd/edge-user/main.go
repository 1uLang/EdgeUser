package main

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/TeaOSLab/EdgeUser/internal/apps"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/nodes"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	ag_ser "github.com/1uLang/zhiannet-api/agent/server"
	nc_model "github.com/1uLang/zhiannet-api/nextcloud/model"
)

func main() {
	app := apps.NewAppCmd().
		Version(teaconst.Version).
		Product(teaconst.ProductName).
		Usage(teaconst.ProcessName + " [-v|start|stop|restart|service|daemon]")
	app.On("daemon", func() {
		nodes.NewUserNode().Daemon()
	})
	app.On("service", func() {
		err := nodes.NewUserNode().InstallSystemService()
		if err != nil {
			fmt.Println("[ERROR]install failed: " + err.Error())
			return
		}
		fmt.Println("done")
	})

	//初始化 第三方包的配置文件
	model.ApiDbPath = Tea.ConfigFile("api_db.yaml")
	model.InitMysqlLink()
	// 初始化agengt和nextcloud配置
	ag_ser.AgentInit(model.ApiDbPath)
	nc_model.InitialAdminUser()
	cache.ApiDbPath = Tea.ConfigFile("api_db.yaml")
	cache.InitClient()
	app.Run(func() {
		userNode := nodes.NewUserNode()
		userNode.Run()
	})
}
