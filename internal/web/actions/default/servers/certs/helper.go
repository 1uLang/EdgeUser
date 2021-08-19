package certs

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"net/http"
)

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

func (this *Helper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["leftMenuItems"] = []maps.Map{
		{
			"name":     "证书",
			"url":      "/servers/certs",
			"isActive": action.Data.GetString("leftMenuItem") == "cert",
		},
		// TODO 暂时不开放
		{
			"name":     "申请证书",
			"url":      "/servers/certs/acme",
			"isActive": action.Data.GetString("leftMenuItem") == "acme",
		},
	}
}
