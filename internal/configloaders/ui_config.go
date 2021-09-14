package configloaders

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/logs"
	"reflect"
	"time"
)

var sharedUIConfig *systemconfigs.UserUIConfig = nil
var HIDSType string = "safedog"

const (
	UISettingName = "userUIConfig"
)

func init() {
	// 更新任务
	// TODO 改成实时更新
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			err := reloadUIConfig()
			if err != nil {
				logs.Println("[CONFIG_LOADERS]load ui config failed: " + err.Error())
			}
		}
	}()
}

func LoadUIConfig() (*systemconfigs.UserUIConfig, error) {
	locker.Lock()
	defer locker.Unlock()

	config, err := loadUIConfig()
	if err != nil {
		return nil, err
	}

	v := reflect.Indirect(reflect.ValueOf(config)).Interface().(systemconfigs.UserUIConfig)
	return &v, nil
}

func loadUIConfig() (*systemconfigs.UserUIConfig, error) {
	if sharedUIConfig != nil {
		return sharedUIConfig, nil
	}
	var rpcClient, err = rpc.SharedRPC()
	if err != nil {
		return nil, err
	}
	resp, err := rpcClient.SysSettingRPC().ReadSysSetting(rpcClient.Context(0), &pb.ReadSysSettingRequest{
		Code: UISettingName,
	})
	if err != nil {
		return nil, err
	}
	if len(resp.ValueJSON) == 0 {
		sharedUIConfig = defaultUIConfig()
		return sharedUIConfig, nil
	}

	config := &systemconfigs.UserUIConfig{}
	err = json.Unmarshal(resp.ValueJSON, config)
	if err != nil {
		logs.Println("[UI_MANAGER]" + err.Error())
		sharedUIConfig = defaultUIConfig()
		return sharedUIConfig, nil
	}
	sharedUIConfig = config
	return sharedUIConfig, nil
}

func reloadUIConfig() error {
	var rpcClient, err = rpc.SharedRPC()
	if err != nil {
		return err
	}
	resp, err := rpcClient.SysSettingRPC().ReadSysSetting(rpcClient.Context(0), &pb.ReadSysSettingRequest{
		Code: UISettingName,
	})
	if err != nil {
		return err
	}
	if len(resp.ValueJSON) == 0 {
		return nil
	}

	config := &systemconfigs.UserUIConfig{}
	err = json.Unmarshal(resp.ValueJSON, config)
	if err != nil {
		return err
	}
	sharedUIConfig = config
	return nil
}

func defaultUIConfig() *systemconfigs.UserUIConfig {
	return &systemconfigs.UserUIConfig{
		ProductName:        "GoEdge",
		UserSystemName:     "GoEdge用户系统",
		ShowOpenSourceInfo: true,
		ShowVersion:        true,
	}
}
