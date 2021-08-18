package host

import (
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"github.com/1uLang/zhiannet-api/common/server/audit_assets_relation_server"
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	user_server "github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type AuthorizeAction struct {
	actionutils.ParentAction
}

type UserList struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	IsOn bool   `json:"is_on"`
	My   bool   `json:"my"`
}

func (this *AuthorizeAction) RunGet(params struct {
	Id   uint64
	Must *actions.Must
}) {

	params.Must.
		Field("id", params.Id).
		Require("请选择主机")

	//获取用户端用户
	list, err := user_server.ListEnabledUsers(&model.ListReq{
		UserId: uint64(this.UserId()),
		Offset: 0,
		Size:   999,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	//获取授权中的列表
	authUse, _, err := audit_assets_relation_server.GetList(&audit_assets_relation.ListReq{
		//UserId: uint64(this.UserId()),
		AssetsId:   []uint64{params.Id},
		AssetsType: 1,
		PageSize:   999,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	authMap := map[uint64]uint64{}
	if len(authUse) > 0 {
		for _, v := range authUse {
			authMap[v.UserId] = v.UserId
		}
	}
	allUsers, authUsers := make([]UserList, 0), make([]UserList, 0)
	if len(list) > 0 {
		for _, v := range list {
			var isOn bool
			if _, ok := authMap[v.Id]; ok {
				isOn = true
			}
			if isOn {
				authUsers = append(authUsers, UserList{
					Id:   v.Id,
					Name: v.Username,
					IsOn: isOn,
					My:   v.Id == uint64(this.UserId()),
				})

			} else {
				allUsers = append(allUsers, UserList{
					Id:   v.Id,
					Name: v.Username,
					IsOn: isOn,
					My:   v.Id == uint64(this.UserId()),
				})
			}

		}
	}
	userResp, err := this.RPC().UserRPC().FindEnabledUser(this.UserContext(), &pb.FindEnabledUserRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	user := userResp.User
	if user == nil {
		this.NotFound("user", this.UserId())
		return
	}
	allUsers = append(allUsers, UserList{
		Id:   uint64(user.Id),
		Name: user.Username,
		IsOn: true,
		My:   true,
	})
	//list, err := audit_host.GetAuthEmail(&server.AuthReq{
	//	Id: params.Id,
	//	User: &request.UserReq{
	//		UserId: uint64(this.UserId()),
	//	},
	//})
	////var email string
	//allUsers, authUsers := make([]UserList, 0), make([]UserList, 0)
	//if err != nil || list == nil {
	//
	//} else {
	//	for _, v := range list.Data.UserList {
	//		if v.IsOn {
	//
	//			if v.My {
	//				allUsers = append(allUsers, UserList{
	//					Id:   v.Id,
	//					Name: v.Name,
	//					My:   v.My,
	//					IsOn: v.IsOn,
	//				})
	//			} else {
	//				authUsers = append(authUsers, UserList{
	//					Id:   v.Id,
	//					Name: v.Name,
	//					My:   v.My,
	//					IsOn: v.IsOn,
	//				})
	//			}
	//		} else {
	//			allUsers = append(allUsers, UserList{
	//				Id:   v.Id,
	//				Name: v.Name,
	//				My:   v.My,
	//				IsOn: v.IsOn,
	//			})
	//		}
	//
	//	}
	//}
	//email = strings.TrimSpace(email)
	this.Data["allUsers"] = allUsers
	this.Data["authUsers"] = authUsers
	this.Success()
}
func (this *AuthorizeAction) RunPost(params struct {
	Id    uint64
	Users []uint64
	Must  *actions.Must
}) {
	if len(params.Users) == 0 {
		params.Users = []uint64{uint64(this.UserId())}
	}
	params.Must.
		Field("id", params.Id).
		Require("请选择主机")

	//res, err := audit_host.AuthHost(&server.AuthReq{
	//	User: &request.UserReq{
	//		UserId: uint64(this.UserId()),
	//	},
	//	Id:  params.Id,
	//	Ids: params.Users,
	//	//Email: emails,
	//})
	//if err != nil || res.Code != 0 {
	//	if err == nil {
	//		err = fmt.Errorf(res.Msg)
	//	}
	//	this.ErrorPage(err)
	//	return
	//}
	users := append([]uint64{uint64(this.UserId())}, params.Users...)
	res := audit_assets_relation_server.Reset(&audit_assets_relation.AddReq{
		UserId:     users,
		AssetsId:   params.Id,
		AssetsType: 1,
	})
	defer this.CreateLogInfo("修改 安全审计-主机 -授权 %v", res)
	if res != nil {
		this.ErrorPage(res)
		return
	}
	this.Success()
}
