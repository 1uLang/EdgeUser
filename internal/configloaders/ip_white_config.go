package configloaders

import (
	"encoding/json"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/logs"
	"reflect"
)

const (
	SecuritySettingName = "userIPWhiteConfig"

	FrameNone       = ""
	FrameDeny       = "DENY"
	FrameSameOrigin = "SAMEORIGIN"
)

var sharedSecurityConfig *systemconfigs.SecurityConfig = nil

func LoadSecurityConfig(userId int64) (*systemconfigs.SecurityConfig, error) {
	locker.Lock()
	defer locker.Unlock()
	fmt.Println("userid----", userId)
	config, err := loadSecurityConfig(userId)
	fmt.Println("config-----", *config)
	if err != nil {
		return nil, err
	}

	v := reflect.Indirect(reflect.ValueOf(config)).Interface().(systemconfigs.SecurityConfig)
	fmt.Println("config2-----", v)
	return &v, nil
}

func UpdateSecurityConfig(securityConfig *systemconfigs.SecurityConfig, userId int64) error {
	locker.Lock()
	defer locker.Unlock()

	var rpcClient, err = rpc.SharedRPC()
	if err != nil {
		return err
	}
	valueJSON, err := json.Marshal(securityConfig)
	if err != nil {
		return err
	}
	_, err = rpcClient.SysSettingRPC().UpdateSysSetting(rpcClient.Context(userId), &pb.UpdateSysSettingRequest{
		Code:      fmt.Sprintf("%s_user%v", SecuritySettingName, userId),
		ValueJSON: valueJSON,
	})
	if err != nil {
		return err
	}
	err = securityConfig.Init()
	if err != nil {
		return err
	}
	sharedSecurityConfig = securityConfig

	// 通知更新
	//events.Notify(events.EventSecurityConfigChanged)

	return nil
}

func loadSecurityConfig(userId int64) (*systemconfigs.SecurityConfig, error) {
	//if sharedSecurityConfig != nil {
	//	return sharedSecurityConfig, nil
	//}
	var rpcClient, err = rpc.SharedRPC()
	if err != nil {
		return nil, err
	}
	resp, err := rpcClient.SysSettingRPC().ReadSysSetting(rpcClient.Context(userId), &pb.ReadSysSettingRequest{
		Code: fmt.Sprintf("%s_user%v", SecuritySettingName, userId),
	})
	fmt.Println("resp------1------", resp)
	if err != nil {
		return nil, err
	}
	if len(resp.ValueJSON) == 0 {
		sharedSecurityConfig = defaultSecurityConfig()
		return sharedSecurityConfig, nil
	}

	config := &systemconfigs.SecurityConfig{}
	err = json.Unmarshal(resp.ValueJSON, config)
	if err != nil {
		logs.Println("[SECURITY_MANAGER]" + err.Error())
		sharedSecurityConfig = defaultSecurityConfig()
		return sharedSecurityConfig, nil
	}
	err = config.Init()
	if err != nil {
		return nil, err
	}
	sharedSecurityConfig = config
	return sharedSecurityConfig, nil
}

func defaultSecurityConfig() *systemconfigs.SecurityConfig {
	return &systemconfigs.SecurityConfig{
		Frame:      FrameSameOrigin,
		AllowLocal: true,
	}
}
