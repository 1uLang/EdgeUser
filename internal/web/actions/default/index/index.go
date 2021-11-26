package index

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/server/edge_admins_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/1uLang/zhiannet-api/common/server/edge_users_server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
	stringutil "github.com/iwind/TeaGo/utils/string"

	"github.com/xlzd/gotp"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

// 首页（登录页）

var TokenSalt = stringutil.Rand(32)

func (this *IndexAction) RunGet(params struct {
	From  string
	Token string
	Auth  *helpers.UserShouldAuth
}) {
	// 已登录跳转到dashboard
	if params.Auth.IsUser() {
		this.RedirectURL("/dashboard")
		return
	}

	this.Data["isUser"] = false
	this.Data["menu"] = "signIn"

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	this.Data["token"] = stringutil.Md5(TokenSalt+timestamp) + timestamp
	this.Data["from"] = params.From

	config, err := configloaders.LoadUIConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["systemName"] = config.UserSystemName
	this.Data["showVersion"] = config.ShowVersion
	if len(config.Version) > 0 {
		this.Data["version"] = config.Version
	} else {
		this.Data["version"] = teaconst.Version
	}
	this.Data["faviconFileId"] = config.FaviconFileId
	if params.Token != "" {
		this.Success()
	}
	this.Show()
}

// 提交
func (this *IndexAction) RunPost(params struct {
	Token    string
	Username string
	Password string
	Remember bool
	OtpCode  string

	Must *actions.Must
	Auth *helpers.UserShouldAuth
	CSRF *actionutils.CSRF
}) {

	this.Data["from"] = ""
	//params.Must.
	//	Field("username", params.Username).
	//	Require("请输入用户名").
	//	Field("password", params.Password).
	//	Require("请输入密码")

	if params.Username == "" {
		this.FailField("username", "请输入用户名")
	}
	if params.Password == stringutil.Md5("") {
		this.FailField("password", "请输入密码")
	}

	// 检查token
	if len(params.Token) <= 32 {
		this.Fail("请通过登录页面登录")
	}
	timestampString := params.Token[32:]
	if stringutil.Md5(TokenSalt+timestampString) != params.Token[:32] {
		this.FailField("refresh", "登录页面已过期，请刷新后重试")
	}
	timestamp := types.Int64(timestampString)
	if timestamp < time.Now().Unix()-1800 {
		this.FailField("refresh", "登录页面已过期，请刷新后重试")
	}

	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		this.Fail("服务器出了点小问题：" + err.Error())
	}
	//登录限制检查
	if res, _ := edge_admins_server.LoginCheck(fmt.Sprintf("user_%v", params.Username)); res {
		this.FailField("refresh", "账号已被锁定（请 30分钟后重试）")
	}
	resp, err := rpcClient.UserRPC().LoginUser(rpcClient.Context(0), &pb.LoginUserRequest{
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		err = dao.SharedLogDAO.CreateUserLog(rpcClient.Context(0), oplogs.LevelError, this.Request.URL.Path, "登录时发生系统错误："+err.Error(), this.RequestRemoteIP())
		if err != nil {
			utils.PrintError(err)
		}
		actionutils.Fail(this, err)
	}

	if !resp.IsOk {
		err = dao.SharedLogDAO.CreateUserLog(rpcClient.Context(0), oplogs.LevelWarn, this.Request.URL.Path, "登录失败，用户名："+params.Username, this.RequestRemoteIP())
		if err != nil {
			utils.PrintError(err)
		}
		info, err := edge_users_server.GetUserInfoByName(params.Username)
		if err != nil {
			this.ErrorPage(err)
		}
		if (info != nil && info.ID > 0) && (info.State == 0 || info.Ison == 0) {
			this.Fail("当前账号被禁用")
		} else {
			//登录次数+1
			edge_admins_server.LoginErrIncr(fmt.Sprintf("user_%v", params.Username))
			//num, _ := cache.GetInt(fmt.Sprintf("user_%v", params.Username))
			//this.Fail(fmt.Sprintf("请输入正确的用户名密码，您还可以尝试%v次，（账号将被临时锁定30分钟）", 5-num))
			this.Fail("登录失败，请重新登录")
		}

	}
	// 检查OTP-*/
	otpInfo, err := edge_logins_server.GetInfoByUid(uint64(resp.UserId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if otpInfo != nil && otpInfo.IsOn == 1 {
		loginParams := maps.Map{}
		err = json.Unmarshal([]byte(otpInfo.Params), &loginParams)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		secret := loginParams.GetString("secret")
		if gotp.NewDefaultTOTP(secret).Now() != params.OtpCode {
			this.Fail("请输入正确的OTP动态密码")
		}
	}
	//密码过期检查
	this.Data["from"] = ""
	if res, _ := edge_users_server.CheckPwdInvalid(uint64(resp.UserId)); res {
		params.Auth.SetUpdatePwdToken(resp.UserId)
		this.Data["from"] = "/updatePwd"
		this.Fail("密码已过期，请立即修改")
	}
	//ip登陆限制检查
	{
		securityConfig, _ := configloaders.LoadSecurityConfig(resp.UserId)
		if !helpers.CheckIP(securityConfig, this.RequestRemoteIP()) {
			//this.ResponseWriter.WriteHeader(http.StatusForbidden)
			this.Fail("当前IP登录被限制")
		}
		//获取父级用户
		userInfo, _ := edge_users_server.GetUserInfo(uint64(resp.UserId))
		if userInfo != nil {
			securityConfig, _ := configloaders.LoadSecurityConfig(int64(userInfo.ParentId))
			if !helpers.CheckIP(securityConfig, this.RequestRemoteIP()) {
				this.Fail("当前IP登录被限制")
			}
		}
	}

	userId := resp.UserId
	params.Auth.StoreUser(userId, params.Remember)

	// 记录日志
	err = dao.SharedLogDAO.CreateUserLog(rpcClient.Context(userId), oplogs.LevelInfo, this.Request.URL.Path, "成功登录系统，用户名："+params.Username, this.RequestRemoteIP())
	if err != nil {
		utils.PrintError(err)
	}
	//记录登录成功30分钟
	cache.SetNx(fmt.Sprintf("login_success_userid_%v", userId), time.Minute*30)
	cache.DelKey(fmt.Sprintf("user_%v", params.Username))

	//跳转首页
	this.Data["from"] = helpers.NewUserMustAuth("").FirstMenuUrl(userId)
	if this.Data["from"] == "" {
		this.Fail("无访问权限，请联系管理员获取模块权限")
	}
	this.Success()
}
