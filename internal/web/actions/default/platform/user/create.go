package user

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	nc_model "github.com/1uLang/zhiannet-api/nextcloud/model"
	nc_req "github.com/1uLang/zhiannet-api/nextcloud/request"
	"github.com/TeaOSLab/EdgeUser/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/dlclark/regexp2"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/xlzd/gotp"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Username string
	Pass1    string
	Pass2    string
	Fullname string
	Mobile   string
	Tel      string
	Email    string
	Remark   string
	OtpIsOn  bool

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("username", params.Username).
		Require("请输入用户名").
		Match(`^[a-zA-Z0-9_]+$`, "用户名中只能含有英文、数字和下划线")

	exists, err := server.CheckUserUsername(&model.CheckUserNameReq{
		UserId:   0,
		Username: params.Username,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if exists {
		this.FailField("username", "此用户名已经被占用，请换一个")
	}
	params.Must.
		Field("pass1", params.Pass1).
		Require("请输入密码").
		Field("pass2", params.Pass2).
		Require("请再次输入确认密码").
		Equal(params.Pass1, "两次输入的密码不一致")

	reg, err := regexp2.Compile(
		`^(?![A-z0-9]+$)(?=.[^%&',;=?$\x22])(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,30}$`, 0)
	if err != nil {
		this.FailField("pass1", "密码格式不正确")
	}
	if match, err := reg.FindStringMatch(params.Pass1); err != nil || match == nil {
		this.FailField("pass1", "密码格式不正确")
	}

	params.Must.
		Field("fullname", params.Fullname).
		Require("请输入全名")

	if len(params.Mobile) > 0 {
		params.Must.
			Field("mobile", params.Mobile).
			Mobile("请输入正确的手机号")
	}
	if len(params.Email) > 0 {
		params.Must.
			Field("email", params.Email).
			Email("请输入正确的电子邮箱")
	}
	UseDatabackup := true
	if UseDatabackup {
		//// 创建nextcloud账号，并写入数据库
		adminToken := nc_req.GetAdminToken()
		userPwd := `adminAd#@2021`
		err = nc_req.CreateUser(adminToken, params.Username, userPwd)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		// 生成token
		gtReq := &nc_model.LoginReq{
			User:     params.Username,
			Password: userPwd,
		}
		ncToken := nc_req.GenerateToken(gtReq)
		// 写入数据库
		err = nc_model.StoreNCToken(params.Username, ncToken)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	//创建审计系统的账号
	//auditResp, auditErr := user.AddUser(&user.AddUserReq{
	//	User:        &request.UserReq{UserId: uint64(this.UserId())},
	//	Email:       params.Email,
	//	IsAdmin:     1,
	//	NickName:    params.Username,
	//	Opt:         1,
	//	Password:    params.Pass1,
	//	Phonenumber: params.Mobile,
	//	RoleIds:     []uint64{},
	//	RoleName:    "平台管理员",
	//	Sex:         1,
	//	Status:      1,
	//	UserName:    params.Username,
	//})
	//if auditErr != nil || auditResp == nil {
	//	this.ErrorPage(fmt.Errorf("创建账号失败"))
	//	return
	//}
	//if auditResp.Code != 0 {
	//	this.ErrorPage(fmt.Errorf(auditResp.Msg))
	//	return
	//}

	userId, err := server.CreateUser(&model.CreateUserReq{
		UserId:   uint64(this.UserId()),
		Username: params.Username,
		Password: params.Pass1,
		Fullname: params.Fullname,
		Mobile:   params.Mobile,
		Tel:      params.Tel,
		Email:    params.Email,
		Remark:   params.Remark,
		Source:   "user:" + numberutils.FormatInt64(this.UserId()),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	defer this.CreateLogInfo("创建用户 %d", userId)

	//关联账号
	//_, err = audit_user_relation.Add(&audit_user_relation.AuditReq{
	//	UserId:      userId,
	//	AuditUserId: uint64(auditResp.Data.Id),
	//})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	if UseDatabackup {
		// 用户账号和nextcloud账号进行关联
		// 因为用户名是唯一的，所以加入用户名字段，减少脏数据的产生
		err = nc_model.BindNCTokenAndUID(params.Username, int64(userId))
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	//otp认证
	if params.OtpIsOn {
		otpLogin := &edge_logins.EdgeLogins{
			Id:   0,
			Type: "otp",
			Params: string(maps.Map{
				"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
			}.AsJSON()),
			IsOn:    1,
			AdminId: 0,
			UserId:  userId,
			State:   1,
		}

		_, err = edge_logins_server.Save(otpLogin)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}
	this.Success()
}
