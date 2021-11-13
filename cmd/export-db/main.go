package main

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_logs"
	"github.com/1uLang/zhiannet-api/common/server"
	"github.com/1uLang/zhiannet-api/common/server/edge_logs_server"
	"github.com/TeaOSLab/EdgeUser/internal/ip2region"
	"github.com/iwind/TeaGo/lists"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/tealeg/xlsx/v3"
	"strings"
)

func main() {

	//初始化 第三方包的配置文件
	server.SetApiDbPath("api_db.yaml")
	server.InitMysqlLink()
	cache.ApiDbPath = "api_db.yaml"

	list, err := edge_logs_server.GetAll()
	if err != nil {
		panic(err)
	}
	if err = writeXlsx(list, "./userLogin-20211027.xlsx"); err != nil {
		panic(err)
	}
}

type loginInfo struct {
	Account string `json:"account"`
	IP      string `json:"ip"`
	Addr    string `json:"addr"`
	Num     int    `json:"numt"`
}

func writeXlsx(ips []*edge_logs.UserLogResp, filename string) error {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("list")
	if err != nil {
		return err
	}
	row = sheet.AddRow()
	row.SetHeightCM(0.8)
	cell = row.AddCell()
	cell.Value = "用户账号"
	cell = row.AddCell()
	cell.Value = "登录IP"
	cell = row.AddCell()
	cell.Value = "IP所属地"
	cell = row.AddCell()
	cell.Value = "登录时间"
	ipStatistics := map[string]loginInfo{}
	for _, i := range ips {
		var row1 *xlsx.Row

		if !strings.Contains(i.Description, "成功登录") || i.UserName == ""{
			continue
		}

		row1 = sheet.AddRow()
		row1.SetHeightCM(0.8)

		cell = row1.AddCell()
		cell.Value = i.UserName
		cell = row1.AddCell()
		cell.Value = i.Ip
		cell = row1.AddCell()
		info, err := ip2region.DB.MemorySearch(i.Ip)
		addr := "-"
		if err == nil {
			pieces := []string{}
			if len(info.Country) > 0 && info.Country != "0" {
				pieces = append(pieces, info.Country)
			}
			if len(info.Province) > 0 && info.Province != "0" && !lists.ContainsString(pieces, info.Province) {
				pieces = append(pieces, info.Province)
			}
			if len(info.City) > 0 && info.City != "0" && !lists.ContainsString(pieces, info.City) {
				pieces = append(pieces, info.City)
			}
			addr = strings.Join(pieces, " ")
		}
		cell.Value = addr
		cell = row1.AddCell()
		cell.Value = timeutil.FormatTime("Y-m-d H:i:s", int64(i.CreatedAt))

		ipinfo, isExist := ipStatistics[i.UserName+addr]
		if isExist {
			ipinfo.Num = ipinfo.Num + 1
		} else {
			ipinfo = loginInfo{
				IP:      i.Ip,
				Addr:    addr,
				Num:     1,
				Account: i.UserName,
			}
		}
		ipStatistics[i.UserName+addr] = ipinfo
	}

	sheet, err = file.AddSheet("statistics")
	if err != nil {
		return err
	}

	row = sheet.AddRow()
	row.SetHeightCM(0.8)
	cell = row.AddCell()
	cell.Value = "用户账号"
	cell = row.AddCell()
	cell.Value = "登录IP"
	cell = row.AddCell()
	cell.Value = "IP所属地"
	cell = row.AddCell()
	cell.Value = "登录次数"

	for _, i := range ipStatistics {
		var row1 *xlsx.Row

		row1 = sheet.AddRow()
		row1.SetHeightCM(0.8)

		cell = row1.AddCell()
		cell.Value = i.Account
		cell = row1.AddCell()
		cell.Value = i.IP
		cell = row1.AddCell()
		cell.Value = i.Addr
		cell = row1.AddCell()
		cell.SetInt(i.Num)
	}
	return file.Save(filename)
}
