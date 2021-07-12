package agent

import (
	"fmt"
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
	username, err := this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取当前用户信息失败：%v", err))
		return
	}
	linux, err := agent_server.Install(username, "Linux")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	windows, err := agent_server.Install(username, "Windows")
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//172.18.200.29 替换
	linux = strings.Replace(linux, "https://172.18.200.29", "https://user.cloudhids.net", 1)
	linux = strings.ReplaceAll(linux, "172.18.200.29", "156.240.95.243")

	windows = strings.Replace(windows, "172.18.200.29", "156.240.95.243", 1)

	this.Data["linux"] = linux
	this.Data["windows"] = windows
	this.Show()
}
