package systemCmd

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

// 异常进程 相关主机
func (this *IndexAction) RunGet(params struct {
	ServerIp string
	PageNo   int
	pageSize int
}) {

	defer this.Show()

	this.Data["datas"] = nil
	this.Data["serverIp"] = params.ServerIp

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &risk.RiskSearchReq{}
	req.ServerIp = params.ServerIp
	req.PageSize = params.pageSize
	req.PageNo = params.PageNo
	req.UserId = uint64(this.UserId())
	list, err := risk_server.SystemCmdList(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取系统命令篡改入侵威胁列表失败：%v", err)
		return
	}
	for k, v := range list.SystemCmdInfoList {

		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverIp"].(string))
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}
		list.SystemCmdInfoList[k]["os"] = os
	}
	this.Data["datas"] = list.SystemCmdInfoList
	this.Data["serverIp"] = params.ServerIp
}

// 异常进程 忽略
func (this *IndexAction) RunPost(params struct {
	Opt     string
	MacCode string
	RiskIds []int
	ItemIds []string
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.ProcessReq{Opt: params.Opt}
	req.Req.MacCode = params.MacCode
	req.Req.RiskIds = params.RiskIds
	req.Req.ItemIds = params.ItemIds

	err = risk_server.ProcessSystemCmd(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) Init() {
	this.FirstMenu("index")
}

// 异常进程列表
func (this *DetailAction) RunGet(params struct {
	MacCode   string
	RiskId    string
	IsProcess bool
	Must      *actions.Must
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")
	params.Must.
		Field("riskId", params.MacCode).
		Require("请输入异常进程id")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)

	}

	info, err := risk_server.SystemCmdDetail(params.MacCode, params.RiskId, params.IsProcess)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["details"] = info

	this.Show()
}

type DetailListAction struct {
	actionutils.ParentAction
}

func (this *DetailListAction) Init() {
	this.FirstMenu("index")
}

// 异常进程列表
func (this *DetailListAction) RunGet(params struct {
	Ip       string
	MacCode  string
	PageNo   int
	PageSize int
	Must     *actions.Must
}) {

	params.Must.Field("macCode", params.MacCode).Require("请输入机器码")
	params.Must.Field("ip", params.Ip).Require("请输入主机ip")

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
	}
	req := &risk.DetailReq{}
	req.Req.PageSize = params.PageSize
	req.Req.PageNo = params.PageNo
	req.MacCode = params.MacCode

	var list1,list2,list3,list4 risk.DetailResp

	details, err := risk_server.SystemCmdDetailList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	for _,v := range details.SystemCmdInfoList {
		if v["state"].(float64) == 1 || v["state"].(float64) == -1{
			list2.SystemCmdInfoList = append(list2.SystemCmdInfoList, v)
		}else if v["state"].(float64) == 2 || v["state"].(float64) == -2{
			list3.SystemCmdInfoList = append(list3.SystemCmdInfoList, v)
		}else if v["state"].(float64) == 3 || v["state"].(float64) == -3{
			list4.SystemCmdInfoList = append(list4.SystemCmdInfoList, v)
		}else{
			list1.SystemCmdInfoList = append(list1.SystemCmdInfoList, v)
		}
	}
	//漏洞列表
	this.Data["datas1"] = list1.SystemCmdInfoList
	this.Data["datas2"] = list2.SystemCmdInfoList
	this.Data["datas3"] = list3.SystemCmdInfoList
	this.Data["datas4"] = list4.SystemCmdInfoList

	this.Data["total1"] = len(list1.SystemCmdInfoList)
	this.Data["total2"] = len(list2.SystemCmdInfoList)
	this.Data["total3"] = len(list3.SystemCmdInfoList)
	this.Data["total4"] = len(list4.SystemCmdInfoList)

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode
	//os
	os, err := server.Info(params.Ip)
	if err != nil {
		this.ErrorPage(err)
	}
	this.Data["os"] = os["osType"]
	//最后扫描时间
	this.Data["lastUpdateTime"] = os["lastUpdateTime"]

	this.Show()
}
