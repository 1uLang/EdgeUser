package hids

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
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
			Get("", new(IndexAction)).
			EndAll()
	})
}
var ServerAddr string
func InitAPIServer() error {

	info, err := server.GetHideInfo()
	if err != nil {
		return fmt.Errorf("主机防护节点获取失败:%v", err)
	}

	err = server.SetUrl(info.Addr)
	if err != nil {
		return err
	}
	ServerAddr = info.Addr
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		AppId:  info.AppId,  //"39rkz",
		Secret: info.Secret, //"tkvgpvjuht2625mo",
	})
	if err != nil {
		return err
	}
	//统计各用户下的入侵威胁
	//目前只统计系统用户下的 主机防护 入侵威胁

	return nil
}
