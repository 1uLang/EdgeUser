package rpc

import (
	"github.com/TeaOSLab/EdgeUser/internal/configs"
	"sync"
)

var sharedRPC *RPCClient = nil
var locker = &sync.Mutex{}

// SharedRPC 获取GRPC客户端
func SharedRPC() (*RPCClient, error) {
	locker.Lock()
	defer locker.Unlock()

	if sharedRPC != nil {
		return sharedRPC, nil
	}

	config, err := configs.LoadAPIConfig()
	if err != nil {
		return nil, err
	}
	client, err := NewRPCClient(config)
	if err != nil {
		return nil, err
	}

	sharedRPC = client
	return sharedRPC, nil
}
