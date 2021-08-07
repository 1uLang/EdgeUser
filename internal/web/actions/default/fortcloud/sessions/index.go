package sessions

import (
	"fmt"
	session_model "github.com/1uLang/zhiannet-api/next-terminal/model/session"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) checkAndNewServerRequest() (*next_terminal_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	return fortcloud.NewServerRequest(fortcloud.Username,fortcloud.Password)
}
func (this *IndexAction) RunGet(params struct {
	PageSize int
	PageNo   int
}) {
	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}
	online, err := req.Session.List(&session_model.ListReq{
		Status: "connected",
		UserId: uint64(this.UserId()),
		PageSize: 999,
		PageIndex: 1,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	fmt.Println(online)
	this.Data["online"] = online
	this.Show()
}
