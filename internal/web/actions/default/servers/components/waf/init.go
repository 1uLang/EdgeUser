package waf

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Helper(NewHelper()).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "waf").
			Prefix("/servers/components/waf").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Get("/policy", new(PolicyAction)).
			Get("/groups", new(GroupsAction)).
			Get("/group", new(GroupAction)).
			Get("/log", new(LogAction)).
			GetPost("/update", new(UpdateAction)).
			GetPost("/test", new(TestAction)).
			GetPost("/export", new(ExportAction)).
			Get("/exportDownload", new(ExportDownloadAction)).
			GetPost("/import", new(ImportAction)).
			Post("/updateGroupOn", new(UpdateGroupOnAction)).
			Post("/deleteGroup", new(DeleteGroupAction)).
			GetPost("/createGroupPopup", new(CreateGroupPopupAction)).
			Post("/sortGroups", new(SortGroupsAction)).
			GetPost("/updateGroupPopup", new(UpdateGroupPopupAction)).
			GetPost("/createSetPopup", new(CreateSetPopupAction)).
			GetPost("/createRulePopup", new(CreateRulePopupAction)).
			Post("/sortSets", new(SortSetsAction)).
			Post("/updateSetOn", new(UpdateSetOnAction)).
			Post("/deleteSet", new(DeleteSetAction)).
			GetPost("/updateSetPopup", new(UpdateSetPopupAction)).
			Post("/count", new(CountAction)).
			Get("/selectPopup", new(SelectPopupAction)).


			EndAll()
	})
}
