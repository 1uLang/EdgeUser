package user

import (
	"encoding/json"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configloaders"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"
)

type OtpQrcodeAction struct {
	actionutils.ParentAction
}

func (this *OtpQrcodeAction) Init() {
	this.Nav("", "", "")
}

func (this *OtpQrcodeAction) RunGet(params struct {
	UserId uint64
}) {

	otpInfo, err := edge_logins_server.GetInfoByUid(params.UserId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if otpInfo == nil || otpInfo.IsOn == 0 {
		this.NotFound("userLogin", int64(params.UserId))
		return
	}
	loginParams := maps.Map{}
	err = json.Unmarshal([]byte(otpInfo.Params), &loginParams)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	secret := loginParams.GetString("secret")

	// 当前用户信息
	userResp, err := this.RPC().UserRPC().FindEnabledUser(this.UserContext(), &pb.FindEnabledUserRequest{UserId: int64(params.UserId)})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	admin := userResp.User
	if admin == nil {
		this.NotFound("userid", int64(params.UserId))
		return
	}

	uiConfig, err := configloaders.LoadUIConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	url := gotp.NewDefaultTOTP(secret).ProvisioningUri(admin.Username, uiConfig.UserSystemName)
	data, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.AddHeader("Content-Type", "image/png")
	this.Write(data)
}
