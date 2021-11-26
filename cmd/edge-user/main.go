package main

import (
	"fmt"
	ag_ser "github.com/1uLang/zhiannet-api/agent/server"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/server"
	nc_model "github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/TeaOSLab/EdgeUser/internal/apps"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/nodes"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	_ "github.com/iwind/TeaGo/bootstrap"
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
	server.SetApiDbPath(utils.Path() + "/build/configs/api_db.yaml")
	server.InitMysqlLink()
	// 初始化agengt和nextcloud配置
	ag_ser.AgentInit(server.GetApiDbPath())
	nc_model.InitialAdminUser()
	cache.ApiDbPath = utils.Path() + "/build/configs/api_db.yaml"
	cache.InitClient()
	app.Run(func() {
		userNode := nodes.NewUserNode()
		userNode.Run()
	})
}
