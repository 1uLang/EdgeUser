package feature

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	NodeId uint64
}) {

	key := "feature-auth-update"
	res, err := cache.GetCache(key)
	authUpdate := false
	if err == nil && res != nil {
		if fmt.Sprintf("%v", res) == "true" {
			authUpdate = true
		}

	}

	this.Data["authUpdate"] = authUpdate

	key = "feature-update-status"
	res, err = cache.GetCache(key)
	status := "立即更新"
	if err == nil && res != nil {
		status = "更新中.."
	}
	this.Data["status"] = status

	key = "feature-update-time"
	res, err = cache.GetCache(key)
	t, _ := util.GetFirstDateOfWeek()
	if err == nil && res != nil {
		t, _ = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%v", res), time.Local)
	}
	this.Data["update_time"] = t.Format("2006-01-02 15:04:05")
	this.Show()

}
