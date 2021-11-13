package ip2region

import (
	"github.com/iwind/TeaGo/logs"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var DB *ip2region.Ip2Region

func init() {
	db, err := ip2region.New("ip2region.db")
	if err != nil {
		logs.Println("[ERROR]" + err.Error())
	}
	DB = db
}
