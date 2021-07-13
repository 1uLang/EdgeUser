package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	examine_server "github.com/1uLang/zhiannet-api/hids/server/examine"
	"github.com/1uLang/zhiannet-api/hids/server/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"strings"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.FirstMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	PageNo       int
	PageSize     int
	State        int
	Score        int
	Type         int
	StartTime    string //体检开始时间
	EndTime      string //体检结束时间
	ExamineItems string //体检项目集合

}) {
	defer this.Show()

	this.Data["datas"] = nil
	this.Data["state"] = 0
	this.Data["Type"] = 0
	this.Data["score"] = 0
	this.Data["examineItems"] = ""
	this.Data["startTime"] = params.StartTime
	this.Data["endTime"] = params.EndTime

	err := hids.InitAPIServer()
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	req := &examine.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取当前用户信息失败：%v", err)
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize

	req.State = params.State - 1
	req.Score = params.Score - 1
	if params.Type == 0 {
		req.Type = -1
	} else {
		req.Type = params.Type
	}

	req.StartTime = params.StartTime
	req.EndTime = params.EndTime
	list, err := examine_server.List(req)
	if err != nil {
		this.Data["errorMessage"] = fmt.Sprintf("获取主机体检信息列表失败：%v", err)
		return
	}
	datas := make([]map[string]interface{}, 0)
	for k, v := range list.ServerExamineResultInfoList {
		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverExamineResultInfo"].(map[string]interface{})["serverIp"].(string), req.UserName)
		if err != nil {
			this.Data["errorMessage"] = fmt.Sprintf("获取主机信息失败：%v", err)
			return
		}
		list.ServerExamineResultInfoList[k]["os"] = os

		if req.State != -1 {
			list.ServerExamineResultInfoList[k]["state"] = req.State
		}

		//看当前主机是否保存了主机体检状态
		_, isExist := gl_examine_scan_maps.Load(v["macCode"].(string) + v["userName"].(string))
		if isExist {
			if v["serverExamineResultInfo"].(map[string]interface{})["state"].(float64) == 0 { //未检测  - 延迟 设置成体检中
				list.ServerExamineResultInfoList[k]["serverExamineResultInfo"].(map[string]interface{})["state"] = 1
			} else { //已生效 删除该记录
				gl_examine_scan_maps.Delete(v["macCode"].(string) + v["userName"].(string))
			}
		}
		datas = append(datas, list.ServerExamineResultInfoList[k])
	}
	this.Data["datas"] = datas
	this.Data["state"] = params.State
	this.Data["Type"] = params.Type
	this.Data["score"] = params.Score

	if len(params.StartTime) > 0 {
		this.Data["startTime"] = strings.ReplaceAll(params.StartTime, " ", "T")
	}

	if len(params.EndTime) > 0 {
		this.Data["endTime"] = strings.ReplaceAll(params.EndTime, " ", "T")
	}
}
func (this *IndexAction) RunPost(params struct {
	PageNo       int
	PageSize     int
	State        int
	Score        int
	Type         int
	StartTime    string //体检开始时间
	EndTime      string //体检结束时间
	ExamineItems string //体检项目集合

}) {
	err := hids.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	req := &examine.SearchReq{}
	req.UserName, err = this.UserName()
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取当前用户信息失败：%v", err))
		return
	}
	req.PageNo = params.PageNo
	req.PageSize = params.PageSize

	req.State = params.State - 1
	req.Score = params.Score - 1
	if params.Type == 0 {
		req.Type = -1
	} else {
		req.Type = params.Type
	}

	req.StartTime = params.StartTime
	req.EndTime = params.EndTime
	list, err := examine_server.List(req)
	if err != nil {
		this.ErrorPage(fmt.Errorf("获取主机体检信息列表失败：%v", err))
		return
	}
	datas := make([]map[string]interface{}, 0)
	for k, v := range list.ServerExamineResultInfoList {
		if v["userName"] != req.UserName {
			continue
		}
		os, err := server.Info(v["serverExamineResultInfo"].(map[string]interface{})["serverIp"].(string), req.UserName)
		if err != nil {
			this.ErrorPage(fmt.Errorf("获取主机信息失败：%v", err))
			return
		}
		list.ServerExamineResultInfoList[k]["os"] = os

		if req.State != -1 {
			list.ServerExamineResultInfoList[k]["state"] = req.State
		}
		//看当前主机是否保存了主机体检状态
		_, isExist := gl_examine_scan_maps.Load(v["macCode"].(string) + v["userName"].(string))
		if isExist {
			if list.ServerExamineResultInfoList[k]["state"].(float64) == 0 { //未检测  - 延迟 设置成体检中
				list.ServerExamineResultInfoList[k]["state"] = 1
			} else { //已生效 删除该记录
				gl_examine_scan_maps.Delete(v["macCode"].(string) + req.UserName)
			}
		}

		datas = append(datas, list.ServerExamineResultInfoList[k])
	}
	this.Data["datas"] = datas
	this.Data["state"] = params.State
	this.Data["Type"] = params.Type
	this.Data["score"] = params.Score

	if len(params.StartTime) > 0 {
		this.Data["startTime"] = strings.ReplaceAll(params.StartTime, " ", "T")
	}

	if len(params.EndTime) > 0 {
		this.Data["endTime"] = strings.ReplaceAll(params.EndTime, " ", "T")
	}
	this.Success()
	return
}
