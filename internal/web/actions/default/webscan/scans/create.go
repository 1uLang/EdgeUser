package scans

import (
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	nessus_scans_model "github.com/1uLang/zhiannet-api/nessus/model/scans"
	nessus_scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"strings"
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
	var scan_func func(string) error
	for _, targetId := range params.TargetIds {

		//主机漏洞扫描
		if strings.HasSuffix(targetId,"-host") {
			scan_func = func(id string) (err error) {
				return nessus_scans_server.Scans(&nessus_scans_model.ScanReq{ID: strings.TrimSuffix(id,"-host")})
			}
		}else{
			scan_func = func(id string) error {
				req := &scans.AddReq{ProfileId: "11111111-1111-1111-1111-111111111111"}
				req.TargetId = id
				return scans_server.Add(req)
			}
		}
		err = scan_func(targetId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	// 日志
	this.CreateLogInfo("漏洞扫描 - 扫描任务目标:%v成功", params.TargetIds)

	this.Success()
}
