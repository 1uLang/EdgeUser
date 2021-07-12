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

type ConfigDefectAction struct {
	actionutils.ParentAction
}

// 配置缺陷 相关主机
func (this *ConfigDefectAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	PageSize int
}) {
	defer this.Show()

	this.Data["configDefects"] = nil
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

	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}
	list, err := risk_server.ConfigDefectList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取缺陷配置详细列表信息失败：%v", err)
		return
	}
	for k, v := range list.List {

		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverIp"].(string), req.UserName)
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}
		list.List[k]["os"] = os
	}
	this.Data["configDefects"] = list.List
	this.Data["serverIp"] = params.ServerIp

}

// 配置缺陷 忽略
func (this *ConfigDefectAction) RunPost(params struct {
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

	err = risk_server.ProcessConfigDefect(req)
	if err != nil {
		this.Error(err.Error(), 400)
		return
	}
	this.Success()
}

type ConfigDefectListAction struct {
	actionutils.ParentAction
}

// 配置缺陷列表
func (this *ConfigDefectListAction) RunGet(params struct {
	Ip             string
	MacCode        string
	Os             string
	LastUpdateTime string
	PageNo         int
	PageSize       int

	Must *actions.Must
}) {

	params.Must.Field("macCode", params.MacCode).Require("请输入机器码")

	defer this.Show()

	this.Data["configDefect1"] = nil
	this.Data["configDefect2"] = nil

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
	req.Req.UserName,err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}

	//待处理
	req.Req.ProcessState = 1
	list1, err := risk_server.ConfigDefectDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		this.Data["errorMessage"] = fmt.Sprintf("获取缺陷配置详细列表失败：%v", err)
		return
	}
	//已处理
	req.Req.ProcessState = 2
	list2, err := risk_server.ConfigDefectDetailList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取缺陷配置详细列表失败：%v", err)
		return
	}
	//漏洞列表
	this.Data["configDefect1"] = list1.ConfigDefectList
	this.Data["configDefect2"] = list2.ConfigDefectList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

}

type ConfigDefectDetailAction struct {
	actionutils.ParentAction
}

func (this *ConfigDefectDetailAction) Init() {
	this.FirstMenu("index")
}

// 弱口令详情
func (this *ConfigDefectDetailAction) RunGet(params struct {
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

	this.Data["ConfigDefectDetails"] = nil

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}

	info, err := risk_server.ConfigDefectDetail(params.MacCode, params.RiskId, params.ProcessState == 2)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取缺陷配置详情信息失败：%v", err)
		return
	}
	this.Data["ConfigDefectDetails"] = info

}
