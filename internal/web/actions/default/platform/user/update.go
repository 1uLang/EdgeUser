package user

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/dlclark/regexp2"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/xlzd/gotp"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunPost(params struct {
	UserId   uint64
	Pass1    string
	Pass2    string
	Fullname string
	Mobile   string
	Email    string
	Remark   string
	IsOn     uint8
	OtpIsOn  bool

	Must *actions.Must
}) {
	defer this.CreateLogInfo("修改用户 %d", params.UserId)

	params.Must.
		Field("userId", params.UserId).
		Require("请选择用户")

	if len(params.Pass1) > 0 {
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

	err := server.UpdateUser(&model.UpdateUserReq{
		Id:       params.UserId,
		Name:     params.Fullname,
		Password: params.Pass1,
		Mobile:   params.Mobile,
		Email:    params.Email,
		Remark:   params.Remark,
		IsOn:     params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//otp认证
	otpLogin, err := edge_logins_server.GetInfoByUid(uint64(params.UserId))
	if err != nil {
		this.ErrorPage(err)
		return
	}
	fmt.Println("otpLogin.is", params.OtpIsOn)
	fmt.Println("otpLogin", otpLogin)
	if params.OtpIsOn {
		if otpLogin == nil || otpLogin.Id == 0 {
			otpLogin = &edge_logins.EdgeLogins{
				Id:   0,
				Type: "otp",
				Params: string(maps.Map{
					"secret": gotp.RandomSecret(16), // TODO 改成可以设置secret长度
				}.AsJSON()),
				IsOn:    1,
				AdminId: 0,
				UserId:  uint64(params.UserId),
				State:   1,
			}
		} else {
			// 如果已经有了，就覆盖，这样可以保留既有的参数
			otpLogin.IsOn = 1
		}

		_, err = edge_logins_server.Save(otpLogin)
		if err != nil {
			this.ErrorPage(err)
			return
		}
	} else {
		fmt.Println("otp=====", otpLogin)
		if otpLogin != nil && otpLogin.Id > 0 {
			_, err = edge_logins_server.UpdateOpt(uint64(otpLogin.Id), 0)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}

	}
	this.Success()
}
