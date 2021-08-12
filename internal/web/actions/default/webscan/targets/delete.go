package targets

import (
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"strings"
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

	var delFunc func(string) (err error)

	for _, targetId := range params.TargetIds {

		//主机漏洞扫描
		if strings.HasSuffix(targetId, "-host") {
			delFunc = func(id string) (err error) {
				return nessus_scans_server.Delete(&nessus_scans_model.DeleteReq{ID: strings.TrimSuffix(id, "-host")})
			}
		} else {
			delFunc = targets_server.Delete
		}

		err = delFunc(targetId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("漏洞扫描 - 删除任务目标:%v成功", params.TargetIds)

	this.Success()
}
