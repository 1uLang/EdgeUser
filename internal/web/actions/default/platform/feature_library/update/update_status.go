package feature

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"math/rand"
	"time"
)

type UpdateStatusAction struct {
	actionutils.ParentAction
}

func (this *UpdateStatusAction) Init() {
	this.Nav("", "", "")
}

func (this *UpdateStatusAction) RunGet(params struct {
	Open bool
}) {
	defer this.Success()
	key := "feature-update-status"
	res, err := cache.GetCache(key)

	if err == nil && res != nil {
		return
	}

	key = "feature-update-time"
	err = cache.SetCache(key, time.Now().Format("2006-01-02 15:04:05"), ^uint32(0))
	if err != nil {
		this.ErrorPage(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	key = "feature-update-status"
	err = cache.SetCache(key, 1, uint32(rand.Intn(30)+int(10)))
	if err != nil {
		this.ErrorPage(err)
		return
	}

}
