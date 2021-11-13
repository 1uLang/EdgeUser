package user

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/user/userutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"strings"
)

type FeaturesAction struct {
	actionutils.ParentAction
}

func (this *FeaturesAction) RunGet(params struct {
	UserId int64
}) {
	allFeatures := userutils.FindAllUserFeatures()

	parentFeatureCodes, err := server.FindUserFeatures(&model.FindUserFeaturesReq{UserId: this.UserId()})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	userFeatureCodes, err := server.FindUserFeatures(&model.FindUserFeaturesReq{UserId: params.UserId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	codeMaps := map[string]bool{}
	for _, v := range parentFeatureCodes {
		v := strings.Split(v, ".")[0]
		codeMaps[v] = true
	}
	nl := []string{}
	for k := range codeMaps {
		nl = append(nl, k)
	}

	featureMaps := []maps.Map{}
	idx := 0
	var checkStr []string
	var parent bool
	for _, feature := range allFeatures {

		//当前用户无权限跳过
		if strings.Contains(feature.Code, ".") {
			checkStr = parentFeatureCodes
			parent = false
		} else {
			parent = true
			checkStr = nl
		}
		if !lists.ContainsString(checkStr, feature.Code) {
			continue
		}
		if parent {
			feature.Code = strings.Split(feature.Code, ".")[0]
		}

		item := maps.Map{
			"name":        feature.Name,
			"code":        feature.Code,
			"bShowChild":  false,
			"description": feature.Description,
			"children":    []maps.Map{},
		}
		//子菜单
		if codes := strings.Split(feature.Code, "."); len(codes) == 2 {

			subItems := featureMaps[idx-1]["children"].([]maps.Map)
			subItems = append(subItems, item)
			featureMaps[idx-1]["children"] = subItems
		} else {
			featureMaps = append(featureMaps, item)
			idx++
		}
	}

	this.Data["features"] = featureMaps
	this.Data["selectList"] = userFeatureCodes
	this.Data["userId"] = params.UserId
	this.Success()
}

func (this *FeaturesAction) RunPost(params struct {
	UserId int64
	Codes  string

	Must *actions.Must
}) {
	if this.UserId() == params.UserId{
		this.Fail("无权限")
		return
	}

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
