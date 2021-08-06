package index

import (
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Prefix("/").
			GetPost("", new(IndexAction)).
			GetPost("updatePwd", new(UpdatePwdAction)).
			GetPost("checkOTP", new(CheckOTPAction)).
			EndAll()
	})
}
