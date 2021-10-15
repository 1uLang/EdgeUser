package feature

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateAction) RunGet(params struct {
	Open bool
}) {

	key := "feature-auth-update"
	err := cache.SetCache(key, params.Open, ^uint32(0))
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()

}
