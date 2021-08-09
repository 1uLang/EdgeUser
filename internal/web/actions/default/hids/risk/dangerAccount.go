package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DangerAccountAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountAction) Init() {
	this.FirstMenu("index")
}

// 危险账号 相关主机
func (this *DangerAccountAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	PageSize int
}) {
	defer this.Show()

	this.Data["dangerAccounts"] = nil
	this.Data["serverIp"] = params.ServerIp

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk.SearchReq{}
	req.ServerIp = params.ServerIp
	req.PageSize = params.PageSize
	req.PageNo = params.PageNo

	req.UserId = uint64(this.UserId())
	list, err := risk_server.DangerAccountList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号信息失败：%v", err)
		return
	}
	req.ProcessState = 2
	list2, err := risk_server.DangerAccountList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号信息失败：%v", err)
		return
	}
	fmt.Println(list.List)
	list.List = append(list.List, list2.List...)
	for k, v := range list.List {
		fmt.Println(v["serverIp"])
		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}else if os == nil{	//无主机信息
			continue
		}
		list.List[k]["os"] = os
	}
	this.Data["dangerAccounts"] = list.List
	this.Data["serverIp"] = params.ServerIp
}

// 危险账号 忽略
func (this *DangerAccountAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []int
	ItemIds []string
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.Error(err.Error(),400)
		return
	}
	req := &risk.ProcessReq{Opt: params.Opt}
	req.Req.MacCode = params.MacCode
	req.Req.RiskIds = params.RiskIds
	req.Req.ItemIds = params.ItemIds

	err = risk_server.ProcessDangerAccount(req)
	if err != nil {
		this.Error(err.Error(),400)
		return
	}
	this.Success()
}

type DangerAccountListAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountListAction) Init() {
	this.FirstMenu("index")
}

// 危险账号列表
func (this *DangerAccountListAction) RunGet(params struct {
	Ip             string
	MacCode        string
	Os             string
	LastUpdateTime string
	PageNo         int
	PageSize       int
}) {
	defer this.Show()

	this.Data["dangerAccount1"] = nil
	this.Data["dangerAccount2"] = nil

	this.Data["total1"] = 0
	this.Data["total2"] = 0

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk.DetailReq{}
	req.MacCode = params.MacCode
	req.Req.PageSize = params.PageSize
	req.Req.PageNo = params.PageNo

	//待处理
	req.Req.ProcessState = 1
	list1, err := risk_server.DangerAccountDetailList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号详细列表失败：%v",err)
		return
	}
	//已处理
	req.Req.ProcessState = 2
	list2, err := risk_server.DangerAccountDetailList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号详细列表失败：%v",err)
		return
	}
	//漏洞列表
	this.Data["dangerAccount1"] = list1.DangerAccountList
	this.Data["dangerAccount2"] = list2.DangerAccountList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

}

type DangerAccountDetailAction struct {
	actionutils.ParentAction
}

func (this *DangerAccountDetailAction) Init() {
	this.FirstMenu("index")
}

// 弱口令详情
func (this *DangerAccountDetailAction) RunGet(params struct {
	MacCode      string
	RiskId       string
	ProcessState int

	Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")
	defer this.Show()

	this.Data["DangerAccountDetails"] = nil

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	info, err := risk_server.DangerAccountDetail(params.MacCode, params.RiskId, params.ProcessState == 2)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取风险账号详情信息失败：%v",err)
		return
	}
	this.Data["DangerAccountDetails"] = info
}
