package fortcloud

import (
	jumpserver_server "github.com/1uLang/zhiannet-api/jumpserver/server"
	"fmt"
)

var ServerUrl = "" //"https://scan-web.zhiannet.com"
var Username = ""
var Password = ""

func InitAPIServer() error {

	info, err := jumpserver_server.GetFortCloud()
	if err != nil {
		return fmt.Errorf("堡垒机节点获取失败:%v", err)
	}
	Username = info.Username
	Password = info.Password
	ServerUrl = info.Addr

	return nil
}
func NewServerRequest(username, password string) ( *jumpserver_server.Request,error ){
	return jumpserver_server.NewServerRequest(ServerUrl,username,password)
}

