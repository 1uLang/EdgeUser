package invade

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct{}) {

	defer this.Show()

	dashboard := []map[string]interface{}{
		{"name": "病毒木马", "url": "virus", "value": 0},
		{"name": "网页后门", "url": "webShell", "value": 0},
		{"name": "反弹shell", "url": "reboundShell", "value": 0},
		{"name": "异常账号", "url": "abnormalAccount", "value": 0},
		{"name": "日志异常删除", "url": "logDelete", "value": 0},
		{"name": "异常登录", "url": "abnormalLogin", "value": 0},
		{"name": "异常进程", "url": "abnormalProcess", "value": 0},
		{"name": "系统命令篡改", "url": "systemCmd", "value": 0},
	}
	names := []string{"病毒木马", "网页后门", "反弹shell", "异常账号", "日志异常删除", "异常登录", "异常进程", "系统命令篡改"}
	this.Data["dashboard"] = dashboard
	this.Data["names"] = names

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	//invadeLock := sync.RWMutex{}
	//invadeWg := sync.WaitGroup{}

	fns := []func(*risk.RiskSearchReq) (risk.RiskSearchResp, error){
		risk.VirusList,
		risk.WebShellList,
		risk.ReboundList,
		risk.AbnormalAccountList,
		risk.LogDeleteList,
		risk.AbnormalLoginList,
		risk.AbnormalProcessList,
		risk.SystemCmdList,
	}
	args := &risk.RiskSearchReq{}

	args.UserId = uint64(this.UserId(true))
	args.PageSize = 50

	for idx, fn := range fns {
		//invadeWg.Add(1)
		//go func(idx int, fn func(*risk.RiskSearchReq) (risk.RiskSearchResp, error)) {
		//	defer invadeWg.Done()
			risk, _ := fn(args)
			//invadeLock.Lock()
			dashboard[idx]["value"] = risk.TotalData
			//invadeLock.Unlock()
		//}(i, f)
	}
	//invadeWg.Wait()

	this.Data["dashboard"] = dashboard

}
