package fortcloud

import (
	"fmt"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
)

var ServerUrl = "" //"https://scan-web.zhiannet.com"
var Username = ""
var Password = ""

func InitAPIServer() error {

	info, err := next_terminal_server.GetFortCloud()
	if err != nil {
		return fmt.Errorf("堡垒机节点获取失败:%v", err)
	}
	Username = info.Username
	Password = info.Password
	ServerUrl = info.Addr
	return nil
}
func NewServerRequest(username, password string) (*next_terminal_server.Request, error) {
	return next_terminal_server.NewServerRequest(ServerUrl, username, password)
}
func AssetCount(uid int64) (uint64,error) {
	err := InitAPIServer()
	if err != nil {
		return 0, err
	}
	req,err := NewServerRequest(Username,Password)
	if err != nil {
		return 0, err
	}

	_,total,err := req.Assets.List(&asset_model.ListReq{UserId: uid})

	return total,err

}