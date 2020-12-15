package main

import (
	"github.com/TeaOSLab/EdgeUser/internal/apps"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/nodes"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func main() {
	app := apps.NewAppCmd().
		Version(teaconst.Version).
		Product(teaconst.ProductName).
		Usage(teaconst.ProcessName + " [-v|start|stop|restart]")

	app.Run(func() {
		userNode := nodes.NewUserNode()
		userNode.Run()
	})
}
