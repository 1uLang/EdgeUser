package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
	"strings"
)

type ScanAction struct {
	actionutils.ParentAction
}

func (this *ScanAction) RunPost(params struct {
	Opt       string
	MacCode   []string
	ScanItems string
	ServerIp  string

	VirusPath    string
	WebShellPath string
	Must         *actions.Must
	//CSRF         *actionutils.CSRF
}) {

	params.Must.
		Field("opt", params.Opt).
		Require("请输入操作方式")

	if params.Opt != "now" && params.Opt != "cancel" {
		this.ErrorPage(fmt.Errorf("操作方式参数错误"))
		return
	}

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.ServerIp != "" {
		req := &agent.SearchReq{}
		req.ServerIp = params.ServerIp
		req.UserName, err = this.UserName()
		if err != nil {
			this.ErrorPage(fmt.Errorf("获取用户信息失败:%v", err))
			return
		}
		list, err := agent_server.List(req)
		if err != nil {
			this.Error(fmt.Sprintf("获取主机信息失败：%v", err), 400)
			return
		}
		if len(list.List) == 0 {
			this.Error(fmt.Sprintf("该主机不存在"), 400)
			return
		} else if state, isExist := list.List[0]["agentState"].(string); isExist && state != "2" { //启用 主机
			this.Error("失败：该主机agent已暂停服务，命令无法执行！", 400)
			return
		}
	}
	if params.Opt == "now" {
		req := &examine.ScanReq{MacCode: params.MacCode}
		//去掉 ','
		params.ScanItems = strings.TrimPrefix(params.ScanItems, ",")
		params.ScanItems = strings.TrimSuffix(params.ScanItems, ",")
		if len(params.ScanItems) > 0 {
			req.ScanItems = strings.Split(params.ScanItems, ",")
		}
		req.ScanConfig.VirusPath = params.VirusPath
		req.ScanConfig.WebShellPath = params.WebShellPath
		err = examine_server.ScanServerNow(req)
	} else {
		err = examine_server.ScanServerCancel(params.MacCode)
	}

	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
