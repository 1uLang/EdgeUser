package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	TargetIds []string
}) {

	if len(params.TargetIds) == 0 {
		this.FailField("username", "请选择需要扫描的目标")
		return
	}

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &scans.AddReq{ProfileId: "11111111-1111-1111-1111-111111111111"}
	for _, targetId := range params.TargetIds {
		req.TargetId = targetId
		err = scans_server.Add(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("WEB漏洞扫描 - 扫描任务目标:%v成功", params.TargetIds)

	this.Success()
}
