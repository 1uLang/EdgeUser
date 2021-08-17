package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_logs"
	edge_logs_server "github.com/1uLang/zhiannet-api/common/server/edge_logs_server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"strings"
	"time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "platform", "index")
}

func (this *IndexAction) RunGet(params struct {
	//PageNum  int
	//PageSize int
	DayFrom  string
	DayTo    string
	Keyword  string
	UserType string
}) {

	this.Data["dayFrom"] = params.DayFrom
	this.Data["dayTo"] = params.DayTo
	this.Data["keyword"] = params.Keyword
	this.Data["userType"] = params.UserType
	var sTime, eTime time.Time
	if params.DayFrom != "" {
		sTime, _ = time.Parse("2006-01-02", params.DayFrom)
	}
	if params.DayTo != "" {
		eTime, _ = time.Parse("2006-01-02", params.DayTo)
	}
	var startTime, endTime uint64
	if sTime.Unix() > 0 {
		startTime = uint64(sTime.Unix())
	}
	if eTime.Unix() > 0 {
		endTime = uint64(eTime.Unix())
	}

	//先获取总数
	total, err := edge_logs_server.GetLogNum(&edge_logs.UserLogReq{
		UserId:    uint64(this.UserId(true)),
		StartTime: startTime,
		EndTime:   endTime,
		Keyword:   params.Keyword,
	})
	//fmt.Println("params===", total)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	fmt.Println(total)
	page := this.NewPage(total)
	this.Data["page"] = page.AsHTML()
	//fmt.Println("page====", page.Offset, page.Size,int(page.Offset / page.Size)+1)
	list := make([]*edge_logs.UserLogResp, 0)
	if total > 0 {
		list, _, err = edge_logs_server.GetLogList(&edge_logs.UserLogReq{
			UserId:    uint64(this.UserId(true)),
			StartTime: startTime,
			EndTime:   endTime,
			Keyword:   params.Keyword,
			PageNum:   int(page.Offset/page.Size) + 1,
			PageSize:  int(page.Size),
		})
	}

	if err != nil {
		this.ErrorPage(err)
		return
	}
	//fmt.Println("params===", len(list))
	logMaps := []maps.Map{}
	for _, log := range list {
		regionName := ""
		regionResp, err := this.RPC().IPLibraryRPC().LookupIPRegion(this.UserContext(), &pb.LookupIPRegionRequest{Ip: log.Ip})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if regionResp.IpRegion != nil {
			pieces := []string{regionResp.IpRegion.Summary}
			if len(regionResp.IpRegion.Isp) > 0 && !lists.ContainsString(pieces, regionResp.IpRegion.Isp) {
				pieces = append(pieces, "| "+regionResp.IpRegion.Isp)
			}
			regionName = strings.Join(pieces, " ")
		}

		logMaps = append(logMaps, maps.Map{
			"id":          log.Id,
			"adminId":     log.AdminId,
			"userId":      log.UserId,
			"description": log.Description,
			"userName":    log.UserName,
			"createdTime": timeutil.FormatTime("Y-m-d H:i:s", int64(log.CreatedAt)),
			"level":       log.Level,
			"type":        log.Type,
			"ip":          log.Ip,
			"region":      regionName,
			"action":      log.Action,
		})
	}
	this.Data["logs"] = logMaps

	this.Show()
}
