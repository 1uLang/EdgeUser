package user

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/user/userutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type FeaturesAction struct {
	actionutils.ParentAction
}

func (this *FeaturesAction) RunGet(params struct {
	UserId int64
}) {
	allFeatures := userutils.FindAllUserFeatures()

	features, err := server.FindUserFeatures(&model.FindUserFeaturesReq{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userFeatureCodes := []string{}
	for _, userFeature := range features {
		userFeatureCodes = append(userFeatureCodes, userFeature)
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

	this.Success()
}

func (this *FeaturesAction) RunPost(params struct {
	UserId int64
	Codes  string

	Must *actions.Must
}) {
	defer this.CreateLogInfo("设置用户 %d 的功能列表", params.UserId)

	err := server.UpdateUserFeatures(&model.UpdateUserFeaturesReq{
		UserId:   params.UserId,
		Features: params.Codes,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Success()
}
