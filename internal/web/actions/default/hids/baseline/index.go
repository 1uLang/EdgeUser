package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/1uLang/zhiannet-api/hids/server/server"
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
	PageNo      int
	PageSize    int
	State       int
	ResultState int

	StartTime string
	EndTime   string
	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	defer this.Show()
	this.Data["baselines"] = nil
	this.Data["State"] = params.State
	this.Data["ResultState"] = params.ResultState

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &baseline.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取用户信息失败：%v", err)
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ResultState = params.ResultState

	state := 0
	if params.State > 0 {
		state = params.State - 1
		req.State = &state
	}
	list, err := baseline_server.List(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取主机合规基线列表失败：%v", err)
		return
	}
	baselines := make([]map[string]interface{}, 0)
	for k, v := range list.List {
		if v["userName"] != req.UserName {
			continue
		}
		overTime, isExist := v["overTime"].(string)
		if isExist && overTime != "" && params.StartTime != "" && params.EndTime != "" && (overTime < params.StartTime || overTime > params.EndTime) {
			continue
		}
		os, err := server.Info(v["serverIp"].(string), req.UserName)
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}

		//看当前主机是否保存了合规基线状态
		_, isExist = gl_baseline_check_maps.Load(v["macCode"].(string) + v["userName"].(string))
		if isExist {
			if v["state"].(float64) == 0 { //未检测  - 延迟 设置成检测中
				list.List[k]["state"] = 1
			} else { //已生效 删除该记录
				gl_baseline_check_maps.Delete(v["macCode"].(string) + req.UserName)
			}
		}

		list.List[k]["os"] = os

		baselines = append(baselines, list.List[k])
	}
	this.Data["baselines"] = baselines
	this.Data["State"] = params.State
	this.Data["ResultState"] = params.ResultState

	if len(params.StartTime) > 0 {
		this.Data["startTime"] = strings.ReplaceAll(params.StartTime, " ", "T")
	}

	if len(params.EndTime) > 0 {
		this.Data["endTime"] = strings.ReplaceAll(params.EndTime, " ", "T")
	}
}

func (this *IndexAction) RunPost(params struct {
	PageNo      int
	PageSize    int
	State       int
	ResultState int

	StartTime string
	EndTime   string
	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &baseline.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取用户信息失败：%v", err))
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize
	req.ResultState = params.ResultState

	state := 0
	if params.State > 0 {
		state = params.State - 1
		req.State = &state
	}
	list, err := baseline_server.List(req)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取主机合规基线列表失败：%v", err))
		return
	}
	for k, v := range list.List {
		if v["userName"] != req.UserName {
			continue
		}
		overTime, isExist := v["overTime"].(string)
		if isExist && overTime != "" && params.StartTime != "" && params.EndTime != "" && overTime > params.StartTime && overTime < params.EndTime {
			continue
		}
		os, err := server.Info(v["serverIp"].(string), req.UserName)
		if err != nil {
			this.ErrorPage(fmt.Errorf("获取主机信息失败：%v", err))
			return
		}
		list.List[k]["os"] = os
	}
	this.Data["baselines"] = list.List
	this.Data["State"] = params.State
	this.Data["ResultState"] = params.ResultState

	if len(params.StartTime) > 0 {
		this.Data["startTime"] = strings.ReplaceAll(params.StartTime, " ", "T")
	}

	if len(params.EndTime) > 0 {
		this.Data["endTime"] = strings.ReplaceAll(params.EndTime, " ", "T")
	}
}
