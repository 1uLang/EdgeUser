package remotelogs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/configs"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/iwind/TeaGo/logs"
	"time"
)

var logChan = make(chan *pb.NodeLog, 1024)

func init() {
	// 定期上传日志
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for range ticker.C {
			err := uploadLogs()
			if err != nil {
				logs.Println("[LOG]" + err.Error())
			}
		}
	}()
}

// 打印普通信息
func Println(tag string, description string) {
	logs.Println("[" + tag + "]" + description)

	nodeConfig, _ := configs.LoadAPIConfig()
	if nodeConfig == nil {
		return
	}

	select {
	case logChan <- &pb.NodeLog{
		Role:        teaconst.Role,
		Tag:         tag,
		Description: description,
		Level:       "info",
		NodeId:      nodeConfig.NumberId,
		CreatedAt:   time.Now().Unix(),
	}:
	default:

	}
}

// 打印警告信息
func Warn(tag string, description string) {
	logs.Println("[" + tag + "]" + description)

	nodeConfig, _ := configs.LoadAPIConfig()
	if nodeConfig == nil {
		return
	}

	select {
	case logChan <- &pb.NodeLog{
		Role:        teaconst.Role,
		Tag:         tag,
		Description: description,
		Level:       "warning",
		NodeId:      nodeConfig.NumberId,
		CreatedAt:   time.Now().Unix(),
	}:
	default:

	}
}

// 打印错误信息
func Error(tag string, description string) {
	logs.Println("[" + tag + "]" + description)

	nodeConfig, _ := configs.LoadAPIConfig()
	if nodeConfig == nil {
		return
	}

	select {
	case logChan <- &pb.NodeLog{
		Role:        teaconst.Role,
		Tag:         tag,
		Description: description,
		Level:       "error",
		NodeId:      nodeConfig.NumberId,
		CreatedAt:   time.Now().Unix(),
	}:
	default:

	}
}

// 上传日志
func uploadLogs() error {
	logList := []*pb.NodeLog{}
Loop:
	for {
		select {
		case log := <-logChan:
			logList = append(logList, log)
		default:
			break Loop
		}
	}
	if len(logList) == 0 {
		return nil
	}
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return err
	}
	_, err = rpcClient.NodeLogRPC().CreateNodeLogs(rpcClient.Context(0), &pb.CreateNodeLogsRequest{NodeLogs: logList})
	return err
}
