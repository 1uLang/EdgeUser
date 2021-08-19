package users

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	timeutil "github.com/iwind/TeaGo/utils/time"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "user")
}

func (this *IndexAction) RunGet(params struct{}) {
	countResp, err := this.RPC().ACMEUserRPC().CountACMEUsers(this.UserContext(), &pb.CountAcmeUsersRequest{
		AdminId: 0,
		UserId:  this.UserId(true),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	usersResp, err := this.RPC().ACMEUserRPC().ListACMEUsers(this.UserContext(), &pb.ListACMEUsersRequest{
		AdminId: 0,
		UserId:  this.UserId(true),
		Offset:  page.Offset,
		Size:    page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userMaps := []maps.Map{}
	for _, user := range usersResp.AcmeUsers {
		userMaps = append(userMaps, maps.Map{
			"id":          user.Id,
			"email":       user.Email,
			"description": user.Description,
			"createdTime": timeutil.FormatTime("Y-m-d H:i:s", user.CreatedAt),
		})
	}
	this.Data["users"] = userMaps

	this.Show()
}
