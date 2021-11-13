package logout

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction actions.Action

// 退出登录
func (this *IndexAction) Run(params struct {
	Auth *helpers.UserShouldAuth
}) {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		this.Fail("服务器出了点小问题：" + err.Error())
	}
	userResp, err := rpcClient.UserRPC().FindEnabledUser(rpcClient.Context(int64(params.Auth.UserId())), &pb.FindEnabledUserRequest{UserId: int64(params.Auth.UserId())})
	if err != nil {
		this.Fail("获取用户信息失败")
		return
	}
	// 记录日志
	err = dao.SharedLogDAO.CreateAdminLog(rpcClient.Context(int64(params.Auth.UserId())), oplogs.LevelInfo, this.Request.URL.Path, "退出系统，用户名："+userResp.User.Username, this.RequestRemoteIP())
	if err != nil {
		utils.PrintError(err)
	}
	params.Auth.Logout()
	this.RedirectURL("/")
}
