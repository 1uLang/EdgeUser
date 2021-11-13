package loopholes

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"math/rand"
	"strconv"
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
	t, _ := util.GetFirstDateOfWeek()
	key := cache.Md5Str(fmt.Sprintf("NVD_CVE-%v", t.Format("2006-01-02")))
	uTotal, err := cache.GetCache(key)
	UTotal := 1
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		uTotalInt := rand.Intn(100)
		err = cache.SetCache(key, uTotalInt, 60*60*24*7)
		UTotal = uTotalInt
		return
	} else {
		UTotal, _ = strconv.Atoi(fmt.Sprintf("%v", uTotal))
	}
	this.Data["version"] = maps.Map{
		"update_time":  t.Format("2006-01-02 15:04"),
		"version":      "NVD_CVE-" + t.Format("20060102"),
		"all_total":    129463,
		"update_total": UTotal,
	}

	this.Show()

}
