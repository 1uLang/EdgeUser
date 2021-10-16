package feature

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/wazuh"
	"github.com/iwind/TeaGo/maps"
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
	wazuh.InitAPIServer()
	version, err := server.RulesInfo()
	fmt.Println("version", version, "err", err)
	if err != nil || version == nil {
		t, _ := util.GetFirstDateOfWeek()

		this.Data["version"] = maps.Map{
			"update_time":  t.Format("2006-01-02 15:04"),
			"version":      "4.2.1",
			"all_total":    "3148",
			"update_total": "251",
			"name":         "",
		}
	} else {

		this.Data["version"] = maps.Map{
			"update_time":  version.UpdateTime,
			"version":      version.Version,
			"all_total":    version.Num,
			"update_total": version.UpdateNum,
		}
	}

	this.Show()

}
