package targets

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	scans_server "github.com/1uLang/zhiannet-api/nessus/server/scans"

	"github.com/1uLang/zhiannet-api/nessus/model/scans"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	"github.com/iwind/TeaGo/actions"
)

//任务目标
type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) RunGet(params struct{}) {
	this.Show()
}
func (this *CreateAction) RunPost(params struct {
	Address  string
	Desc     string
	Type     int
	Username string
	Password string
	Port     int
	Os       int
	More     bool
	Must     *actions.Must
	CSRF     *actionutils.CSRF
}) {

	params.Must.
		Field("address", params.Address).
		Require("请输入目标地址").
		Match("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?", "请输入正确的目标地址")

	params.Must.
		Field("type", params.Type).
		Require("请选择类型")

	err := webscan.InitAPIServer()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if params.Type == 1 {

		req := &targets.AddReq{Address: params.Address, UserId: uint64(this.UserId(true))}
		req.Description = params.Desc
		if params.More { //高级设置 身份登录

			params.Must.
				Field("username", params.Username).
				Require("请输入用户名").
				Field("password", params.Password).
				Require("请输入密码")

			req.Username = params.Username
			req.Password = params.Password
		}

		_, err = targets_server.Add(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		// 日志
		this.CreateLogInfo("WEB漏洞扫描 - 创建任务目标:[%v]成功", params.Address)
	} else if params.Type == 2 {

		req := &scans.AddReq{}
		req.UserId = uint64(this.UserId(true))
		req.Settings.Name = params.Address
		req.Settings.Text_targets = params.Address
		req.Settings.Description = params.Desc
		if params.More { //高级设置 身份登录

			params.Must.
				Field("username", params.Username).
				Require("请输入用户名").
				Field("password", params.Password).
				Require("请输入密码").
				Field("port", params.Port).
				Require("请输入端口").
				Field("os", params.Os).
				Require("请选择系统类型")

			req.Username = params.Username
			req.Password = params.Password
			req.Port = params.Port
			req.Os = params.Os

		}

		//先忽略错误。防止页面重复提交
		go func() {
			err = scans_server.Create(req)
			if err != nil {
				this.CreateLogInfo("主机漏洞扫描 - 创建任务目标:[%v]失败:%v", params.Address, err)
			} else {
				this.CreateLogInfo("主机漏洞扫描 - 创建任务目标:[%v]成功", params.Address)
			}
			// 日志
		}()

	} else {
		this.ErrorPage(fmt.Errorf("漏洞扫描类型参数错误"))
		return
	}
	this.Success()
}
