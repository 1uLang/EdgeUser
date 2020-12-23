package webutils

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

// 根据ServerId查找Web配置
func FindWebConfigWithServerId(parentAction *actionutils.ParentAction, serverId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := parentAction.RPC().ServerRPC().FindAndInitServerWebConfig(parentAction.UserContext(), &pb.FindAndInitServerWebConfigRequest{ServerId: serverId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 根据LocationId查找Web配置
func FindWebConfigWithLocationId(parentAction *actionutils.ParentAction, locationId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := parentAction.RPC().HTTPLocationRPC().FindAndInitHTTPLocationWebConfig(parentAction.UserContext(), &pb.FindAndInitHTTPLocationWebConfigRequest{LocationId: locationId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 根据WebId查找Web配置
func FindWebConfigWithId(parentAction *actionutils.ParentAction, webId int64) (*serverconfigs.HTTPWebConfig, error) {
	resp, err := parentAction.RPC().HTTPWebRPC().FindEnabledHTTPWebConfig(parentAction.UserContext(), &pb.FindEnabledHTTPWebConfigRequest{WebId: webId})
	if err != nil {
		return nil, err
	}
	config := &serverconfigs.HTTPWebConfig{}
	err = json.Unmarshal(resp.WebJSON, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
