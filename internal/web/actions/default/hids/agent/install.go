package agent

import (
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"strings"
)

type InstallAction struct {
	actionutils.ParentAction
}

func (this *InstallAction) Init() {
	this.FirstMenu("index")
}

func (this *InstallAction) RunGet(params struct {
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	linux, err := agent_server.Install( "Linux")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	windows, err := agent_server.Install( "Windows")
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//172.18.200.29 替换
	url := hids.ServerAddr
	if url == "" {
		url = "https://hids.zhiannet.com"
	}
	linux = strings.Replace(linux, "https://172.18.200.29", url, 1)
	linux = strings.ReplaceAll(linux, "172.18.200.29", "119.8.57.159")

	windows = strings.Replace(windows, "172.18.200.29", "119.8.57.159", 1)

	this.Data["linux"] = linux
	this.Data["windows"] = windows
	this.Show()
}
