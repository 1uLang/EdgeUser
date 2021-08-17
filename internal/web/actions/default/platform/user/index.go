package user

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/user/userutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/tidwall/gjson"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "user", "index")
}
func (this *IndexAction) RunGet() {
	this.Data["users"] = []map[string]interface{}{}
	count, err := server.CountAllEnabledUsers(&model.GetNumReq{UserId: uint64(this.UserId())})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	list, err := server.ListEnabledUsers(&model.ListReq{
		UserId: uint64(this.UserId()),
		Offset: int(page.Offset),
		Size:   int(page.Size),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(list) > 0 {
		userMaps := []maps.Map{}
		for _, user := range list {
			dom := gjson.Parse(user.OtpParams)
			userMaps = append(userMaps, maps.Map{
				"id":          user.Id,
				"username":    user.Username,
				"isOn":        user.IsOn,
				"fullname":    user.Name,
				"email":       user.Email,
				"mobile":      user.Mobile,
				"tel":         user.Tel,
				"remark":      user.Remark,
				"otpIsOn":     user.OtpOn == 1,
				"createdTime": timeutil.FormatTime("Y-m-d H:i:s", user.CreatedAt),
				"otpParams":   dom.Get("secret").String(),
			})
		}
		this.Data["users"] = userMaps
	}
	//判断用户权限
	this.Data["createIsOn"],err = userutils.CheckCreateIsOn(this.UserId(),"platform.userCreate")
	this.Show()
}
