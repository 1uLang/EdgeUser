package stat

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "stat", "")
	this.SecondMenu("index")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
	Month    string
}) {
	month := params.Month
	if len(month) != 6 {
		month = timeutil.Format("Ym")
	}
	this.Data["month"] = month

	serverTypeResp, err := this.RPC().ServerRPC().FindEnabledServerType(this.UserContext(), &pb.FindEnabledServerTypeRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverType := serverTypeResp.Type

	statIsOn := false

	// 是否已开启
	if serverconfigs.IsHTTPServerType(serverType) {
		webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.UserContext(), params.ServerId)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if webConfig != nil && webConfig.StatRef != nil {
			statIsOn = webConfig.StatRef.IsOn
		}
	} else {
		this.WriteString("此类型服务暂不支持统计")
		return
	}

	this.Data["statIsOn"] = statIsOn

	// 统计数据
	countryStatMaps := []maps.Map{}
	provinceStatMaps := []maps.Map{}
	cityStatMaps := []maps.Map{}

	if statIsOn {
		// 地区
		{
			resp, err := this.RPC().ServerRegionCountryMonthlyStatRPC().FindTopServerRegionCountryMonthlyStats(this.UserContext(), &pb.FindTopServerRegionCountryMonthlyStatsRequest{
				Month:    month,
				ServerId: params.ServerId,
				Offset:   0,
				Size:     10,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			for _, stat := range resp.Stats {
				countryStatMaps = append(countryStatMaps, maps.Map{
					"count": stat.Count,
					"country": maps.Map{
						"id":   stat.RegionCountry.Id,
						"name": stat.RegionCountry.Name,
					},
				})
			}
		}

		// 省份
		{
			resp, err := this.RPC().ServerRegionProvinceMonthlyStatRPC().FindTopServerRegionProvinceMonthlyStats(this.UserContext(), &pb.FindTopServerRegionProvinceMonthlyStatsRequest{
				Month:    month,
				ServerId: params.ServerId,
				Offset:   0,
				Size:     10,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			for _, stat := range resp.Stats {
				provinceStatMaps = append(provinceStatMaps, maps.Map{
					"count": stat.Count,
					"country": maps.Map{
						"id":   stat.RegionCountry.Id,
						"name": stat.RegionCountry.Name,
					},
					"province": maps.Map{
						"id":   stat.RegionProvince.Id,
						"name": stat.RegionProvince.Name,
					},
				})
			}
		}

		// 城市
		{
			resp, err := this.RPC().ServerRegionCityMonthlyStatRPC().FindTopServerRegionCityMonthlyStats(this.UserContext(), &pb.FindTopServerRegionCityMonthlyStatsRequest{
				Month:    month,
				ServerId: params.ServerId,
				Offset:   0,
				Size:     10,
			})
			if err != nil {
				this.ErrorPage(err)
				return
			}
			for _, stat := range resp.Stats {
				cityStatMaps = append(cityStatMaps, maps.Map{
					"count": stat.Count,
					"country": maps.Map{
						"id":   stat.RegionCountry.Id,
						"name": stat.RegionCountry.Name,
					},
					"province": maps.Map{
						"id":   stat.RegionProvince.Id,
						"name": stat.RegionProvince.Name,
					},
					"city": maps.Map{
						"id":   stat.RegionCity.Id,
						"name": stat.RegionCity.Name,
					},
				})
			}
		}
	}

	this.Data["countryStats"] = countryStatMaps
	this.Data["provinceStats"] = provinceStatMaps
	this.Data["cityStats"] = cityStatMaps

	this.Show()
}
