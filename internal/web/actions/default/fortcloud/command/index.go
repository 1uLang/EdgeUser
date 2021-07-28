package command

import (
	"fmt"
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	commands_model "github.com/1uLang/zhiannet-api/jumpserver/model/commands"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "webscan", "index")
}

func (this *IndexAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}

	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
	//return fortcloud.NewServerRequest("admin", "21ops.com")
}
func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int

	Index     int
	Asset     string
	Input     string
	Username  string
	RiskLevel string
	DayFrom   string
	DayTo     string
}) {
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	//2021-07-25T13:13:27.753Z
	layout := "2006-01-02T15:04:05.000Z"
	now := time.Now()
	if params.Index >= 0 {
		params.DayTo = now.Format(layout)
	}
	switch params.Index {
	case 0: //最近一天
		params.DayFrom = now.AddDate(0, 0, -1).Format(layout)
	case 1: //最近一周
		params.DayFrom = now.AddDate(0, 0, -7).Format(layout)
	case 2: //最近一月
		params.DayFrom = now.AddDate(0, -1, 0).Format(layout)
	case 3: //最近三月
		params.DayFrom = now.AddDate(0, -3, 0).Format(layout)
	}
	args := &commands_model.ListReq{
		UserId: uint64(this.UserId()),
	}
	args.Asset = params.Asset
	args.Input = params.Input
	args.System_user = params.Username
	args.Risk_level = params.RiskLevel
	args.Date_from = params.DayFrom
	args.Date_to = params.DayTo

	//指定该用户下
	args.User, _ = this.UserName()

	commands, err := req.Command.List(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	assets, err := req.Assets.List(&assets_model.ListReq{
		UserId: uint64(this.UserId()),
	})
	page := this.NewPage(int64(len(commands)))
	this.Data["page"] = page.AsHTML()

	offset := page.Offset
	end := offset + page.Size
	if end > page.Total {
		end = page.Total
	}
	this.Data["commands"] = commands[offset:end]
	this.Data["assets"] = assets

	this.Data["index"] = params.Index
	if len(params.DayFrom) >= 10 {
		this.Data["dayFrom"] = params.DayFrom[:10]
	} else {
		this.Data["dayFrom"] = ""
	}
	if len(params.DayTo) >= 10 {
		this.Data["dayTo"] = params.DayTo[:10]
	} else {
		this.Data["dayTo"] = ""
	}
	this.Data["username"] = params.Username
	this.Data["riskLevel"] = params.RiskLevel
	this.Data["input"] = params.Input
	this.Data["asset"] = params.Asset

	this.Show()
}
