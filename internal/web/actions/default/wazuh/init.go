// 主机防护使用wazuh组件

package wazuh

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//主机防护

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "hids").
			Prefix("/hids").
			Data("teaMenu", "agents").
			Get("/agents", new(AgentsAction)).
			Get("/agents/create", new(CreateAction)).
			Post("/agents/delete", new(AgentsAction)).
			Data("teaMenu", "vulnerability").
			Get("/vulnerability", new(VulnerabilityAction)).
			Data("teaMenu", "virus").
			Get("/virus", new(VirusAction)).
			Data("teaMenu", "baseLine").
			Get("/baseLine", new(BaseLineAction)).
			Get("/baseLineDetails", new(BaseLineDetailsAction)).
			Data("teaMenu", "attck").
			Get("/attck", new(AttckAction)).
			Data("teaMenu", "invades").
			Get("/invades", new(InvadeAction)).
			Data("teaMenu", "syscheck").
			Get("/syscheck", new(SysCheckAction)).
			EndAll()
	})
}

var serverAddr string

func InitAPIServer() error {

	info, err := server.GetWazuhInfo()
	if err != nil {
		return fmt.Errorf("主机防护节点获取失败:%v", err)
	}
	serverAddr = info.Addr
	err = server.SetUrl(info.Addr)
	if err != nil {
		return err
	}
	err = server.InitToken(info.Username, info.Password)
	if err != nil {
		return err
	}

	return nil
}
