package targets

import (
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	TargetIds []string `json:"target_ids"`
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要删除的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _, targetId := range params.TargetIds {
		err = targets_server.Delete(targetId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 删除任务目标:%v成功", params.TargetIds)

	this.Success()
}
