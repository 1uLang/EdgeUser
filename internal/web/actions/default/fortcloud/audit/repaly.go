package audit

import (
	"fmt"
	session_model "github.com/1uLang/zhiannet-api/next-terminal/model/session"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type ReplayAction struct {

	actionutils.ParentAction
}

func (this *ReplayAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	return fortcloud.NewServerRequest(fortcloud.Username, fortcloud.Password)
}

func (this *ReplayAction) RunGet(params struct {
	Id   string
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择资产")

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	args := &session_model.ReplayReq{}
	args.Id = params.Id
	buf,err := req.Session.Replay(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}// 日志
	this.CreateLogInfo("堡垒机 - 回放会话:[%v]成功", params.Id)
	this.Write(buf)
}
