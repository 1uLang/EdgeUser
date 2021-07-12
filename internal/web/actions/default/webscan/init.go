package webscan

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//web漏洞扫描
func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "webscan").
			Prefix("/webscan").
			Get("", new(IndexAction)).
			EndAll()
	})
}

var ServerUrl = "" //"https://scan-web.zhiannet.com"
var Key = ""

func InitAPIServer() error {

	info, err := server.GetWebScan()
	if err != nil {
		return fmt.Errorf("漏扫节点获取失败:%v", err)
	}
	Key = info.Key
	ServerUrl = info.Addr

	err = server.SetUrl(ServerUrl)
	if err != nil {
		return err
	}
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		XAuth: Key,
	})
	if err != nil {
		return err
	}
	return nil
}
