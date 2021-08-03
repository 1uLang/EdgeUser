package admins

import (
	"fmt"
	admin_users_model "github.com/1uLang/zhiannet-api/jumpserver/model/admin_users"
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud"
	"github.com/iwind/TeaGo/actions"
)

type AssetsListAction struct {
	actionutils.ParentAction
}

func (this *AssetsListAction) checkAndNewServerRequest() (*jumpserver_server.Request, error) {
	if fortcloud.ServerUrl == "" {
		err := fortcloud.InitAPIServer()
		if err != nil {
			return nil, err
		}
	}
	username, _ := this.UserName()
	return fortcloud.NewServerRequest(username, "dengbao-"+username)
}
func (this *AssetsListAction) Init() {
	this.Nav("", "fortcloud", "index")
}
func (this *AssetsListAction) RunPost(params struct {
	Limit  int
	Offset int
	Id     string

	Must *actions.Must
}) {

	req, err := this.checkAndNewServerRequest()
	if err != nil {
		this.ErrorPage(fmt.Errorf("堡垒机组件错误:" + err.Error()))
		return
	}

	args := &admin_users_model.AssetsListReq{}
	args.Limit = params.Limit
	args.Offset = params.Offset
	args.Systemuser = params.Id
	args.UserId = uint64(this.UserId())
	list, err := req.AdminUser.AssetsList(args)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["assetsList"] = list
	// 日志
	this.CreateLogInfo("堡垒机 - 管理用户资产列表:[%v]成功", params.Id)
	this.Success()
}
