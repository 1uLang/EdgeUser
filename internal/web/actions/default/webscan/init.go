package webscan

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	nessus_request "github.com/1uLang/zhiannet-api/nessus/request"
	nessus_server "github.com/1uLang/zhiannet-api/nessus/server"
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
var NessusServerUrl = ""
func InitAPIServer() error {

	info, err := server.GetWebScan()
	if err != nil {
		return fmt.Errorf("web漏扫节点获取失败:%v", err)
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

	nessus_info, err := nessus_server.GetNessus()
	if err != nil {
		return fmt.Errorf("主机漏扫节点获取失败:%v", err)
	}
	err = nessus_server.SetUrl(nessus_info.Addr)
	NessusServerUrl = nessus_info.Addr
	if err != nil {
		return err
	}
	err = nessus_server.SetAPIKeys(&nessus_request.APIKeys{
		Access: nessus_info.Access,
		Secret: nessus_info.Secret,
	})
	if err != nil {
		return err
	}
	return nil
}
