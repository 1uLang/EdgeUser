package waf

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/server/logs_statistics_server"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"time"
)

type DdosAction struct {
	actionutils.ParentAction
}

func (this *DdosAction) RunGet(params struct {
	Address    string
	NodeId     uint64
	StartTime  string
	EndTime    string
	AttackType string
	Status     int
	Report     int
}) {
	defer this.Show()

	this.Data["attacks"] = nil
	this.Data["ddos"] = nil
	this.Data["nodeId"] = params.NodeId
	//2006-01-02
	this.Data["startTime"] = params.StartTime
	this.Data["endTime"] = params.EndTime
	this.Data["address"] = params.Address
	this.Data["attackType"] = params.AttackType
	this.Data["status"] = params.Status

	//ddos节点
	ddos, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(ddos) == 0 {
		this.ErrorPage(fmt.Errorf("未配置DDoS防火墙节点"))
		return
	}
	if params.NodeId == 0 {
		params.NodeId = ddos[0].Id
	}

	req := &logs_server.AttackLogReq{
		NodeId:     params.NodeId,
		Addr:       params.Address,
		AttackType: params.AttackType,
		Status:     params.Status,
	}
	if params.StartTime != "" && params.EndTime != "" {
		sT, err := time.ParseInLocation("2006-01-02", params.StartTime, time.Local)
		if err != nil {
			this.ErrorPage(fmt.Errorf("起始时间参数错误"))
			return
		}
		eT, err := time.ParseInLocation("2006-01-02", params.EndTime, time.Local)
		if err != nil {
			this.ErrorPage(fmt.Errorf("结束时间参数错误"))
			return
		}
		req.StartTime = sT
		req.EndTime = eT
	}

	list, err := logs_server.GetAttackLogList(req)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取统计日志列表失败：%v", err))
		return
	}

	page := this.NewPage(int64(len(list.Report)))
	this.Data["page"] = page.AsHTML()
	offset := page.Offset
	if offset > int64(len(list.Report)) {
		offset = 0
	}
	end := offset + page.Size
	if end > int64(len(list.Report)) {
		end = int64(len(list.Report))
	}
	this.Data["attacks"] = list.Report[offset:end]
	this.Data["ddos"] = ddos
	this.Data["nodeId"] = params.NodeId

	this.Data["startTime"] = params.StartTime

	this.Data["endTime"] = params.EndTime

	this.Data["address"] = params.Address
	this.Data["attackType"] = params.AttackType
	this.Data["status"] = params.Status

	//周报 日报
	reportList := maps.Map{
		"lineValue": []interface{}{},
		"lineData":  []interface{}{},
	}
	reportLists, _ := logs_statistics_server.GetWafStatistics([]int64{int64(params.NodeId)}, params.Report, 1)
	if len(reportLists) > 0 {
		lineValue := []interface{}{}
		lineData := []interface{}{}
		for _, v := range reportLists {
			lineValue = append(lineValue, v.Time)
			lineData = append(lineData, v.Value)
		}
		reportList["lineValue"] = lineValue
		reportList["lineData"] = lineData
	}
	this.Data["detailTableData"] = reportList
}
