package configloaders

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/systemconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/logs"
	"reflect"
)

var sharedUIConfig *systemconfigs.UserUIConfig = nil

const (
	UISettingName = "userUIConfig"
)

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

func UpdateUIConfig(uiConfig *systemconfigs.UserUIConfig) error {
	locker.Lock()
	defer locker.Unlock()

	var rpcClient, err = rpc.SharedRPC()
	if err != nil {
		return err
	}
	valueJSON, err := json.Marshal(uiConfig)
	if err != nil {
		return err
	}
	_, err = rpcClient.SysSettingRPC().UpdateSysSetting(rpcClient.Context(0), &pb.UpdateSysSettingRequest{
		Code:      UISettingName,
		ValueJSON: valueJSON,
	})
	if err != nil {
		return err
	}
	sharedUIConfig = uiConfig

	return nil
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

func defaultUIConfig() *systemconfigs.UserUIConfig {
	return &systemconfigs.UserUIConfig{
		ProductName:        "GoEdge",
		UserSystemName:     "GoEdge用户系统",
		ShowOpenSourceInfo: true,
		ShowVersion:        true,
	}
}
