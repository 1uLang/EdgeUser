package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
)

type SystemRiskAction struct {
	actionutils.ParentAction
}

func (this *SystemRiskAction) Init() {
	this.FirstMenu("index")
}

// 系统漏洞 相关主机
func (this *SystemRiskAction) RunGet(params struct {
	Level    int
	ServerIp string
	PageNo   int
	PageSize int
}) {
	defer this.Show()

	this.Data["risks"] = nil
	this.Data["serverIp"] = params.ServerIp
	this.Data["level"] = params.Level

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk.SearchReq{}
	req.Level = params.Level
	req.ServerIp = params.ServerIp
	req.PageSize = params.PageSize
	req.PageNo = params.PageNo

	req.UserId = uint64(this.UserId(true))

	//系统漏洞数汇总
	risk, err := risk_server.SystemDistributed(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Errorf("获取系统漏洞信息失败：%v", err)
		return
	}
	for k, v := range risk.List {

		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		} else if os == nil { //无主机信息
			continue
		}
		risk.List[k]["os"] = os
	}
	this.Data["risks"] = risk.List
	this.Data["serverIp"] = params.ServerIp
	this.Data["level"] = params.Level
}

// 系统漏洞 忽略
func (this *SystemRiskAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []int
	ItemIds []string
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.Error(err.Error(), 400)
		return
	}
	req := &risk.ProcessReq{Opt: params.Opt}
	req.Req.MacCode = params.MacCode
	req.Req.RiskIds = params.RiskIds
	req.Req.ItemIds = params.ItemIds

	err = risk_server.ProcessRisk(req)

	if err != nil {
		this.Error(err.Error(), 400)
		return
	}
	this.Success()
}
