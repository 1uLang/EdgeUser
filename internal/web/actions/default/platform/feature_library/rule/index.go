package feature

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/wazuh"
	"github.com/iwind/TeaGo/maps"
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
	wazuh.InitAPIServer()
	version, err := server.RulesInfo()
	fmt.Println("version", version, "err", err)
	if err != nil || version == nil {
		//this.Show()
		//this.ErrorPage(err)
		//return
		this.Data["version"] = maps.Map{
			"update_time":  time.Now().Format("2006-01-02 15:04"),
			"version":      "版本",
			"all_total":    "总特征数",
			"update_total": "更新特征数",
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
