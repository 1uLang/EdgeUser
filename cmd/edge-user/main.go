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
	cache.ApiDbPath = Tea.ConfigFile("api_db.yaml")
	cache.InitClient()
	app.Run(func() {
		userNode := nodes.NewUserNode()
		userNode.Run()
	})
}
