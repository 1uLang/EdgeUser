package hostlist

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) RunGet(params struct{}) {

	// 文件列表（不包含目录）
	//list, err := host_server.HostList(&host.HostListReq{})
	//if err != nil {
	//	this.ErrorPage(err)
	//	return
	//}
	//s, _ := json.Marshal(list.Inventories)
	//dom := gjson.ParseBytes(s)
	//lists := make([]map[string]interface{}, 0)
	//if len(dom.Array()) > 0 {
	//	for _, v := range dom.Array() {
	//		memorySize := v.Get("memorySize").Int()
	//		momory := memorySize / 1024 / 1024 / 1024
	//		//cTime, _ := time.ParseInLocation("Mon 02, 2006 03:04:05", v.Get("createDate").String(), time.Local)
	//		row := map[string]interface{}{
	//			"name":       v.Get("name").String(),
	//			"cpu":        v.Get("cpuNum").Int(),
	//			"memory":     momory,
	//			"ip":         v.Get("vmNics.0.ip").String(),
	//			"state":      v.Get("state").String(),
	//			"createDate": v.Get("createDate").String(),
	//			//"createDate":cTime.Format("2006-01-02 15:04:05"),
	//		}
	//		lists = append(lists, row)
	//	}
	//}
	//this.Data["list"] = lists

	this.Show()
}
