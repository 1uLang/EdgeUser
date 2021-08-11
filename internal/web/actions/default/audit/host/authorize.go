package host

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server"
	"github.com/1uLang/zhiannet-api/audit/server/audit_host"
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

	list, err := audit_host.GetAuthEmail(&server.AuthReq{
		Id: params.Id,
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
	})
	//var email string
	allUsers, authUsers := make([]UserList, 0), make([]UserList, 0)
	if err != nil || list == nil {

	} else {
		for _, v := range list.Data.UserList {
			if v.IsOn {

				if v.My {
					allUsers = append(allUsers, UserList{
						Id:   v.Id,
						Name: v.Name,
						My:   v.My,
						IsOn: v.IsOn,
					})
				} else {
					authUsers = append(authUsers, UserList{
						Id:   v.Id,
						Name: v.Name,
						My:   v.My,
						IsOn: v.IsOn,
					})
				}
			} else {
				allUsers = append(allUsers, UserList{
					Id:   v.Id,
					Name: v.Name,
					My:   v.My,
					IsOn: v.IsOn,
				})
			}

		}
	}
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

	params.Must.
		Field("id", params.Id).
		Require("请选择主机")

	res, err := audit_host.AuthHost(&server.AuthReq{
		User: &request.UserReq{
			UserId: uint64(this.UserId()),
		},
		Id:  params.Id,
		Ids: params.Users,
		//Email: emails,
	})
	if err != nil || res.Code != 0 {
		if err == nil {
			err = fmt.Errorf(res.Msg)
		}
		this.ErrorPage(err)
		return
	}
	defer this.CreateLogInfo("修改 安全审计-主机 -授权 %v", res.Msg)
	this.Success()
}
