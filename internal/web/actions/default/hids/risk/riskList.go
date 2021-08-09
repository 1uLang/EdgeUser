package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
)

type SystemRiskListAction struct {
	actionutils.ParentAction
}

func (this *SystemRiskListAction) Init() {
	this.FirstMenu("index")
}

// 系统漏洞列表
func (this *SystemRiskListAction) RunGet(params struct {
	Ip       string
	MacCode  string
	Page   int
	PageSize int
}) {

	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &risk.SearchReq{}
	req.PageSize = 1
	req.PageNo = params.Page
	req.UserId = uint64(this.UserId())

	//待处理
	req.ProcessState = 1
	list1, err := risk_server.SystemRiskList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//得到总数
	page := this.NewPage(int64(list1.TotalData))
	this.Data["page1"] = page.AsHTML()

	req.PageSize = int(page.Size)
	req.PageNo = params.Page

	list1, err = risk_server.SystemRiskList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//已处理
	req.ProcessState = 2
	req.PageSize = 1
	req.PageNo = 1
	list2, err := risk_server.SystemRiskList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//得到总数
	page2 := this.NewPage(int64(list2.TotalData))
	this.Data["page2"] = page2.AsHTML()

	req.PageSize = int(page2.Size)
	req.PageNo = params.Page
	list2, err = risk_server.SystemRiskList(req)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//漏洞列表
	this.Data["risks1"] = list1.SystemRiskInfoList
	this.Data["risks2"] = list2.SystemRiskInfoList

	this.Data["total1"] = list1.TotalData
	this.Data["total2"] = list2.TotalData

	this.Data["ip"] = params.Ip
	this.Data["macCode"] = params.MacCode
	//os
	os, err := server.Info(params.Ip)
	if err != nil {
		this.ErrorPage(err)
	} else if os == nil{	//无主机信息
		this.ErrorPage(fmt.Errorf("无主机信息"))
	}
	this.Data["os"] = os["osType"]
	//最后扫描时间
	this.Data["lastScanTime"] = os["lastUpdateTime"]

	this.Show()
}
