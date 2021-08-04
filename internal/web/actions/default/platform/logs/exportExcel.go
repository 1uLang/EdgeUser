package logs

import (
	"bytes"
	"github.com/1uLang/zhiannet-api/common/model/edge_logs"
	edge_logs_server "github.com/1uLang/zhiannet-api/common/server/edge_logs_server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/lists"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/tealeg/xlsx/v3"
	"strconv"
	"strings"
	"time"
)

type ExportExcelAction struct {
	actionutils.ParentAction
}

func (this *ExportExcelAction) Init() {
	this.Nav("", "", "")
}

func (this *ExportExcelAction) RunGet(params struct {
	DayFrom  string
	DayTo    string
	Keyword  string
	UserType string
}) {
	//logsResp, err := this.RPC().LogRPC().ListLogs(this.UserContext(), &pb.ListLogsRequest{
	//	Offset:   0,
	//	Size:     1000, // 日志最大导出1000条，TODO 将来可以配置
	//	DayFrom:  params.DayFrom,
	//	DayTo:    params.DayTo,
	//	Keyword:  params.Keyword,
	//	UserType: params.UserType,
	//})
	var err error
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
	logsResp := make([]*edge_logs.UserLogResp, 0)
	logsResp, _, err = edge_logs_server.GetLogList(&edge_logs.UserLogReq{
		UserId:    uint64(this.UserId()),
		StartTime: startTime,
		EndTime:   endTime,
		Keyword:   params.Keyword,
		PageNum:   0,
		PageSize:  1000,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("default")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 头部
	{
		row := sheet.AddRow()
		row.SetHeight(25)
		row.AddCell().SetString("ID")
		row.AddCell().SetString("日期")
		row.AddCell().SetString("用户")
		row.AddCell().SetString("描述")
		row.AddCell().SetString("IP")
		row.AddCell().SetString("区域")
		row.AddCell().SetString("运营商")
		row.AddCell().SetString("页面地址")
	}

	// 数据
	for _, log := range logsResp {
		regionName := ""
		ispName := ""
		regionResp, err := this.RPC().IPLibraryRPC().LookupIPRegion(this.UserContext(), &pb.LookupIPRegionRequest{Ip: log.Ip})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if regionResp.IpRegion != nil {
			pieces := []string{}
			if len(regionResp.IpRegion.Country) > 0 {
				pieces = append(pieces, regionResp.IpRegion.Country)
			}
			if len(regionResp.IpRegion.Province) > 0 && !lists.ContainsString(pieces, regionResp.IpRegion.Province) {
				pieces = append(pieces, regionResp.IpRegion.Province)
			}
			if len(regionResp.IpRegion.City) > 0 && !lists.ContainsString(pieces, regionResp.IpRegion.City) && !lists.ContainsString(pieces, strings.TrimSuffix(regionResp.IpRegion.Province, "市")) {
				pieces = append(pieces, regionResp.IpRegion.City)
			}
			regionName = strings.Join(pieces, " ")

			if len(regionResp.IpRegion.Isp) > 0 {
				ispName = regionResp.IpRegion.Isp
			}
		}

		row := sheet.AddRow()
		row.SetHeight(25)
		row.AddCell().SetInt64(int64(log.Id))
		row.AddCell().SetString(timeutil.FormatTime("Y-m-d H:i:s", int64(log.CreatedAt)))
		if log.UserId > 0 {
			row.AddCell().SetString("用户 | " + log.UserName)
		} else {
			row.AddCell().SetString(log.UserName)
		}
		row.AddCell().SetString(log.Description)
		row.AddCell().SetString(log.Ip)
		row.AddCell().SetString(regionName)
		row.AddCell().SetString(ispName)
		row.AddCell().SetString(log.Action)
	}

	this.AddHeader("Content-Type", "application/vnd.ms-excel")
	this.AddHeader("Content-Disposition", "attachment; filename=\"LOG-"+timeutil.Format("YmdHis")+".xlsx\"")
	this.AddHeader("Cache-Control", "max-age=0")

	buf := bytes.NewBuffer([]byte{})
	err = wb.Write(buf)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.AddHeader("Content-Length", strconv.Itoa(buf.Len()))
	this.Write(buf.Bytes())
}
