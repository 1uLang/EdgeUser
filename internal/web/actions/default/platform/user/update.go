package user

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
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
		Id:        params.UserId,
		Name:      params.Fullname,
		Password:      params.Pass1,
		Mobile:        params.Mobile,
		Email:         params.Email,
		Remark:        params.Remark,
		IsOn:          params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
