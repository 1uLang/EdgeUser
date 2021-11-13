package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
	"strings"
	"sync"
)

type CheckAction struct {
	actionutils.ParentAction
}

var gl_baseline_check_maps sync.Map

func (this *CheckAction) RunPost(params struct {
	MacCode    []string `json:"macCodes"`
	TemplateId int      `json:"templateId"`
	ServerIp   string
	Must       *actions.Must
	//CSRF *actionutils.CSRF
}) {

	params.Must.
		Field("templateId", params.TemplateId).
		Require("请输入合规基线模板")

	if len(params.MacCode) == 0 {
		this.Error("请选择机器码", 400)
		return
	}
	err := hids.InitAPIServer()
	if err != nil {
		this.Error(err.Error(), 400)
		return
	}
	if params.ServerIp != "" {
		params.ServerIp = strings.ReplaceAll(params.ServerIp, "/", ".")
		req := &agent.SearchReq{}
		req.ServerIp = params.ServerIp
		req.UserId = uint64(this.UserId(true))

		list, err := agent_server.List(req)
		if err != nil {
			this.Error(fmt.Sprintf("获取主机信息失败：%v", err), 400)
			return
		}
		if len(list.List) == 0 {
			this.Error(fmt.Sprintf("该主机不存在"), 400)
			return
		} else if state, isExist := list.List[0]["agentState"].(string); state != "2" && isExist { //启用 主机
			this.Error("失败：该主机agent已暂停服务，命令无法执行！", 400)
			return
		}
	}

	req := &baseline.CheckReq{MacCodes: params.MacCode, TemplateId: params.TemplateId}
	err = baseline_server.Check(req)

	if err != nil {
		this.Error(fmt.Sprintf("检测失败：%v", err), 400)
		return
	}
	userName, _ := this.UserName()
	if len(params.MacCode) > 0 && params.MacCode[0] != "" { //该主机进行合规基线
		gl_baseline_check_maps.Store(params.MacCode[0]+userName, true)
	}

	this.Success()
}
