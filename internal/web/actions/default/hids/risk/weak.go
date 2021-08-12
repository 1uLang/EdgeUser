package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
)

type WeakAction struct {
	actionutils.ParentAction
}

func (this *WeakAction) Init() {
	this.FirstMenu("index")
}

// 弱口令 相关主机
func (this *WeakAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	PageSize int
}) {

	defer this.Show()

	this.Data["weaks"] = nil
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
	list, err := risk_server.WeakList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Errorf("获取弱口令信息失败：%v", err)
		return
	}
	req.ProcessState = 2
	list2, err := risk_server.WeakList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Errorf("获取弱口令信息失败：%v", err)
		return
	}
	list.List = append(list.List, list2.List...)
	for k, v := range list.List {
		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v",err)
			return
		}else if os == nil {
			continue
		}
		list.List[k]["os"] = os
	}
	this.Data["weaks"] = list.List
	this.Data["serverIp"] = params.ServerIp
}

// 弱口令 忽略
func (this *WeakAction) RunPost(params struct {
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
	err = risk_server.ProcessWeak(req)
	if err != nil {
		this.Error(err.Error(),400)
		return
	}
	this.Success()
}

type WeakListAction struct {
	actionutils.ParentAction
}

func (this *WeakListAction) Init() {
	this.FirstMenu("index")
}

// 弱口令列表
func (this *WeakListAction) RunGet(params struct {
	Ip             string
	MacCode        string
	Os             string
	LastUpdateTime string
	PageNo         int
	PageSize       int
}) {
	defer this.Show()

	this.Data["weak1"] = nil
	this.Data["weak2"] = nil

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
	list1, err := risk_server.WeakDetailList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取弱口令详细列表信息失败：%v",err)
		return
	}
	//已处理
	req.Req.ProcessState = 2
	list2, err := risk_server.WeakDetailList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取弱口令详细列表信息失败：%v",err)
		return
	}
	//漏洞列表
	this.Data["weak1"] = list1.WeakInfoList
	this.Data["weak2"] = list2.WeakInfoList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode

	this.Data["os"] = params.Os
	//最后扫描时间
	this.Data["lastUpdateTime"] = params.LastUpdateTime

}
