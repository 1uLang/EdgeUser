package accesskeys

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	accessKeysResp, err := this.RPC().UserAccessKeyRPC().FindAllEnabledUserAccessKeys(this.UserContext(), &pb.FindAllEnabledUserAccessKeysRequest{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	accessKeyMaps := []maps.Map{}
	for _, accessKey := range accessKeysResp.UserAccessKeys {
		accessKeyMaps = append(accessKeyMaps, maps.Map{
			"id":          accessKey.Id,
			"isOn":        accessKey.IsOn,
			"uniqueId":    accessKey.UniqueId,
			"secret":      accessKey.Secret,
			"description": accessKey.Description,
		})
	}
	this.Data["accessKeys"] = accessKeyMaps

	this.Show()
}
