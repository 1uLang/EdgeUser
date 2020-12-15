package dashboard

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/utils/numberutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"math"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	dashboardResp, err := this.RPC().UserRPC().ComposeUserDashboard(this.UserContext(), &pb.ComposeUserDashboardRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["dashboard"] = maps.Map{
		"countServers":            dashboardResp.CountServers,
		"monthlyTrafficBytes":     numberutils.HumanBytes1000(dashboardResp.MonthlyTrafficBytes * 8),
		"monthlyPeekTrafficBytes": numberutils.HumanBytes1000(dashboardResp.MonthlyPeekTrafficBytes * 8),
		"dailyTrafficBytes":       numberutils.HumanBytes1000(dashboardResp.DailyTrafficBytes * 8),
		"dailyPeekTrafficBytes":   numberutils.HumanBytes1000(dashboardResp.DailyPeekTrafficBytes * 8),
	}

	// 每日流量统计
	{
		statMaps := []maps.Map{}
		for _, stat := range dashboardResp.DailyTrafficStats {
			statMaps = append(statMaps, maps.Map{
				"count": math.Ceil((float64(stat.Count)*8/1000/1000)*100) / 100,
				"day":   stat.Day[4:6] + "月" + stat.Day[6:] + "日",
			})
		}
		this.Data["dailyTrafficStats"] = statMaps
	}

	// 每日峰值带宽统计
	{
		statMaps := []maps.Map{}
		for _, stat := range dashboardResp.DailyPeekTrafficStats {
			statMaps = append(statMaps, maps.Map{
				"count": math.Ceil((float64(stat.Count)*8/1000/1000/1000)*100) / 100,
				"day":   stat.Day[4:6] + "月" + stat.Day[6:] + "日",
			})
		}
		this.Data["dailyPeekTrafficStats"] = statMaps
	}

	this.Show()
}
