package audit

import (
	"fmt"
	sessions_model "github.com/1uLang/zhiannet-api/jumpserver/model/sessions"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
)

type IndexAction struct {
	actionutils.ParentAction
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

	offline, err := req.Session.List(&sessions_model.ListReq{
		Is_finished:"1",
		UserId: uint64(this.UserId()),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["offline"] = offline
	this.Show()
}
