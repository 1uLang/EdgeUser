package main

import (
	"fmt"
	"github.com/TeaOSLab/EdgeUser/internal/apps"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/nodes"
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
	app.Run(func() {
		userNode := nodes.NewUserNode()
		userNode.Run()
	})
}
