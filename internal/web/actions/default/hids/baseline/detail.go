package examine

import (
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
	baseline_server "github.com/1uLang/zhiannet-api/hids/server/baseline"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction struct {
	actionutils.ParentAction
}

func (this *DetailAction) RunGet(params struct {
	MacCode    string
	PageSize   int
	CheckCount int
	Time       string
	Must       *actions.Must
	//CSRF *actionutils.CSRF
}) {
	params.Must.
		Field("macCode", params.MacCode).
		Require("请输入机器码")

	if err := hids.InitAPIServer(); err != nil {
		this.ErrorPage(err)
		return
	}

	req := &baseline.DetailReq{}
	req.MacCode = params.MacCode
	req.PageNo = 1
	if params.PageSize > 100 {
		req.PageSize = 100 //最大值
	} else {
		req.PageSize = params.PageSize
	}
	sum := params.PageSize
	var systemSafe []map[string]interface{}
	var middleSafe []map[string]interface{}
	for req.PageSize > 0 {
		info, err := baseline_server.Detail(req)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		for _, v := range info.List {
			if v["result"].(float64) == 1 { //去掉正常项
				continue
			}
			if v["typeName"].(string) == "操作系统安全" {
				systemSafe = append(systemSafe, v)
			} else if v["typeName"].(string) == "中间件安全" {
				middleSafe = append(middleSafe, v)
			}
		}
		sum -= req.PageSize
		if req.PageSize >= sum {
			req.PageSize = sum
		}
	}

	this.Data["systemSafe"] = systemSafe
	this.Data["systemSafeCount"] = len(systemSafe)
	this.Data["middleSafeCount"] = len(middleSafe)
	this.Data["middleSafe"] = middleSafe
	this.Data["checkCount"] = params.CheckCount
	this.Data["totalCount"] = params.PageSize
	this.Data["checkTime"] = params.Time

	this.Show()
}
