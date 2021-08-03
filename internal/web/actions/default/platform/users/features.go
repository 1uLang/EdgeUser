package users

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/users/userutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type FeaturesAction struct {
	actionutils.ParentAction
}

func (this *FeaturesAction) Init() {
	this.Nav("", "", "feature")
}

func (this *FeaturesAction) RunGet(params struct {
	UserId int64
}) {
	err := userutils.InitUser(this.Parent(), params.UserId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	featuresResp, err := this.RPC().UserRPC().FindAllUserFeatureDefinitions(this.UserContext(), &pb.FindAllUserFeatureDefinitionsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	allFeatures := featuresResp.Features

	userFeaturesResp, err := this.RPC().UserRPC().FindUserFeatures(this.UserContext(), &pb.FindUserFeaturesRequest{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userFeatureCodes := []string{}
	for _, userFeature := range userFeaturesResp.Features {
		userFeatureCodes = append(userFeatureCodes, userFeature.Code)
	}

	featureMaps := []maps.Map{}
	for _, feature := range allFeatures {
		featureMaps = append(featureMaps, maps.Map{
			"name":        feature.Name,
			"code":        feature.Code,
			"description": feature.Description,
			"isChecked":   lists.ContainsString(userFeatureCodes, feature.Code),
		})
	}

	this.Data["features"] = featureMaps

	this.Show()
}

func (this *FeaturesAction) RunPost(params struct {
	UserId int64
	Codes  []string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo("设置用户 %d 的功能列表", params.UserId)

	_, err := this.RPC().UserRPC().UpdateUserFeatures(this.UserContext(), &pb.UpdateUserFeaturesRequest{
		UserId:       params.UserId,
		FeatureCodes: params.Codes,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	moduleCodes := map[string]bool{}
	for _, code := range params.Codes {
		moduleCodes[code] = true
	}
	this.Success()
}
