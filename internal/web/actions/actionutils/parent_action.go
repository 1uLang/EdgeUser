package actionutils

import (
	"context"
	"errors"
	"fmt"
	edge_user_model "github.com/1uLang/zhiannet-api/edgeUsers/model"
	edge_user_server "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/dao"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/oplogs"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"net/http"
	"strconv"
)

type ParentAction struct {
	actions.ActionObject
	userName 	string
	parentId 	uint64
	rpcClient *rpc.RPCClient
}

// 可以调用自身的一个简便方法
func (this *ParentAction) Parent() *ParentAction {
	return this
}

func (this *ParentAction) ErrorPage(err error) {
	if err == nil {
		return
	}

	// 日志
	this.CreateLog(oplogs.LevelError, "系统发生错误：%s", err.Error())

	if this.Request.Method == http.MethodGet {
		FailPage(this, err)
	} else {
		Fail(this, err)
	}
}

func (this *ParentAction) ErrorText(err string) {
	this.ErrorPage(errors.New(err))
}

func (this *ParentAction) NotFound(name string, itemId int64) {
	this.ErrorPage(errors.New(name + " id: '" + strconv.FormatInt(itemId, 10) + "' is not found"))
}

func (this *ParentAction) NewPage(total int64, size ...int64) *Page {
	if len(size) > 0 {
		return NewActionPage(this, total, size[0])
	}
	return NewActionPage(this, total, 20)
}

func (this *ParentAction) Nav(mainMenu string, tab string, firstMenu string) {
	this.Data["mainMenu"] = mainMenu
	this.Data["mainTab"] = tab
	this.Data["firstMenuItem"] = firstMenu
}

func (this *ParentAction) FirstMenu(menuItem string) {
	this.Data["firstMenuItem"] = menuItem
}

func (this *ParentAction) SecondMenu(menuItem string) {
	this.Data["secondMenuItem"] = menuItem
}

func (this *ParentAction) TinyMenu(menuItem string) {
	this.Data["tinyMenuItem"] = menuItem
}

func (this *ParentAction) UserId(parent ...bool) int64 {
	userId := this.Context.GetInt64("userId")
	if len(parent) > 0 && parent[0]{
		parentId,err := this.ParentId()
		if err != nil {
			logs.Error(err)
		}
		if parentId != 0 {
			return int64(parentId)
		}
	}

	return userId
}
func (this *ParentAction) ParentId() (uint64, error) {
	if this.parentId != 0 {
		return this.parentId,nil
	}
	parentId,err := edge_user_server.GetParentId(&edge_user_model.GetParentIdReq{UserId: uint64(this.UserId())})
	if err != nil {
		return 0, err
	}
	this.parentId = parentId
	return this.parentId,nil
}
func (this *ParentAction) UserName() (string, error) {

	if this.userName != "" {
		return this.userName, nil
	}
	userResp, err := this.RPC().UserRPC().FindEnabledUser(this.UserContext(), &pb.FindEnabledUserRequest{UserId: this.UserId()})
	if err != nil {
		return "", err
	}
	user := userResp.User
	if user == nil {
		return "", fmt.Errorf("无效的用户id")
	}
	this.userName = user.Username
	return user.Username, nil
}

func (this *ParentAction) CreateLog(level string, description string, args ...interface{}) {
	desc := fmt.Sprintf(description, args...)
	if level == oplogs.LevelInfo {
		if this.Code != 200 {
			level = oplogs.LevelWarn
			if len(this.Message) > 0 {
				desc += " 失败：" + this.Message
			}
		}
	}
	err := dao.SharedLogDAO.CreateUserLog(this.UserContext(), level, this.Request.URL.Path, desc, this.RequestRemoteIP())
	if err != nil {
		utils.PrintError(err)
	}
}

func (this *ParentAction) CreateLogInfo(description string, args ...interface{}) {
	this.CreateLog(oplogs.LevelInfo, description, args...)
}

// 获取RPC
func (this *ParentAction) RPC() *rpc.RPCClient {
	if this.rpcClient != nil {
		return this.rpcClient
	}

	// 所有集群
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		logs.Fatal(err)
		return nil
	}
	this.rpcClient = rpcClient

	return rpcClient
}

// 获取Context
func (this *ParentAction) UserContext() context.Context {
	if this.rpcClient == nil {
		rpcClient, err := rpc.SharedRPC()
		if err != nil {
			logs.Fatal(err)
			return nil
		}
		this.rpcClient = rpcClient
	}
	return this.rpcClient.Context(this.UserId(true))
}

// 校验Feature
func (this *ParentAction) ValidateFeature(feature string) bool {
	// 用户功能
	userFeatureResp, err := this.RPC().UserRPC().FindUserFeatures(this.UserContext(), &pb.FindUserFeaturesRequest{UserId: this.UserId()})
	if err != nil {
		logs.Error(err)
		return false
	}
	userFeatureCodes := []string{}
	for _, feature := range userFeatureResp.Features {
		userFeatureCodes = append(userFeatureCodes, feature.Code)
	}

	return lists.ContainsString(userFeatureCodes, feature)
}
