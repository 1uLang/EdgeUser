package risk

import (
	"fmt"
	risk_model "github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
}) {

	defer this.Show()

	var risk, weak, dangerAccount, configDefect risk_model.SystemDistributedResp
	this.Data["risk"] = risk
	this.Data["weak"] = weak
	this.Data["dangerAccount"] = dangerAccount
	this.Data["configDefect"] = configDefect
	this.Data["names"] = []string{"系统漏洞", "弱口令", "风险账号", "配置缺陷"}
	this.Data["datas"] = []map[string]interface{}{
		{"name": "系统漏洞", "value": 0},
		{"name": "弱口令", "value": 0},
		{"name": "风险账号", "value": 0},
		{"name": "配置缺陷", "value": 0},
	}

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk_model.SearchReq{}
	req.UserId = uint64(this.UserId(true))

	//系统漏洞数汇总
	risk, err = risk_server.SystemDistributed(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取系统漏洞信息失败：%v", err)
		return
	}
	//弱口令
	weak, err = risk_server.WeakList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Sprintf("获取弱口令信息失败：%v", err)
		return
	}
	//风险账号
	dangerAccount, err = risk_server.DangerAccountList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号信息失败：%v", err)
		return
	}
	//缺陷配置
	configDefect, err = risk_server.ConfigDefectList(req)
	if err != nil {

		this.Data["errorMessage"] = fmt.Sprintf("获取缺陷配置信息失败：%v", err)
		return
	}
	this.Data["risk"] = risk
	this.Data["weak"] = weak
	this.Data["dangerAccount"] = dangerAccount
	this.Data["configDefect"] = configDefect
	fmt.Println(risk)
	this.Data["datas"] = []map[string]interface{}{
		{"name": "系统漏洞", "value": risk.Total},
		{"name": "弱口令", "value": weak.Total},
		{"name": "风险账号", "value": dangerAccount.Total},
		{"name": "配置缺陷", "value": configDefect.Total},
	}

}
