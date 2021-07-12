package agent

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageNo        int
	PageSize      int
	ServerIp      string
	ServerLocalIp string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	defer this.Show()
	this.Data["agents"] = nil

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	req := &agent.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ServerIp = params.ServerIp
	req.ServerLocalIp = params.ServerLocalIp

	list, err := agent_server.List(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取agent主机列表失败：%v", err)
		return
	}
	this.Data["agents"] = list.List
}

func (this *IndexAction) RunPost(params struct {
	PageNo        int
	PageSize      int
	ServerIp      string
	ServerLocalIp string

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	req := &agent.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取用户信息失败：%v", err))
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ServerIp = params.ServerIp
	req.ServerLocalIp = params.ServerLocalIp

	list, err := agent_server.List(req)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取agent主机列表失败：%v", err))
		return
	}
	this.Data["agents"] = list.List
	this.Success()
}
