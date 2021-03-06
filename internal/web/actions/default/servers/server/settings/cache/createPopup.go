package cache

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	CacheRefJSON []byte

	Must *actions.Must
}) {
	cacheRef := &serverconfigs.HTTPCacheRef{}
	err := json.Unmarshal(params.CacheRefJSON, cacheRef)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if len(cacheRef.Key) == 0 {
		this.Fail("请输入缓存Key")
	}

	err = cacheRef.Init()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["cacheRef"] = cacheRef

	this.Success()
}
