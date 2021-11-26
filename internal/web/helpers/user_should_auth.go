package helpers

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/utils/numberutils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
)

type UserShouldAuth struct {
	action *actions.ActionObject
}

func (this *UserShouldAuth) BeforeAction(actionPtr actions.ActionWrapper, paramName string) (goNext bool) {
	this.action = actionPtr.Object()

	var action = this.action
	action.AddHeader("X-Frame-Options", "SAMEORIGIN")
	action.AddHeader("Content-Security-Policy", "default-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'")

	return true
}

// 存储用户名到SESSION
func (this *UserShouldAuth) StoreUser(userId int64, remember bool) {
	// 修改sid的时间
	if remember {
		cookie := &http.Cookie{
			Name:     teaconst.CookieSID,
			Value:    this.action.Session().Sid,
			Path:     "/",
			MaxAge:   14 * 86400,
			HttpOnly: true,
		}
		if this.action.Request.TLS != nil {
			cookie.SameSite = http.SameSiteStrictMode
			cookie.Secure = true
		}
		this.action.AddCookie(cookie)
	} else {
		cookie := &http.Cookie{
			Name:     teaconst.CookieSID,
			Value:    this.action.Session().Sid,
			Path:     "/",
			MaxAge:   0,
			HttpOnly: true,
		}
		if this.action.Request.TLS != nil {
			cookie.SameSite = http.SameSiteStrictMode
			cookie.Secure = true
		}
		this.action.AddCookie(cookie)
	}
	this.action.Session().Write("userId", numberutils.FormatInt64(userId))
}

func (this *UserShouldAuth) IsUser() bool {
	return this.action.Session().GetInt("userId") > 0
}

func (this *UserShouldAuth) UserId() int64 {
	return this.action.Session().GetInt64("userId")
}

func (this *UserShouldAuth) Logout() {
	cache.DelKey(fmt.Sprintf("login_success_userid_%v", this.UserId()))
	this.action.Session().Delete()
}
