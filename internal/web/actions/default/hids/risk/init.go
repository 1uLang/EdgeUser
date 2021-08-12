package risk

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

//漏洞风险

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth("")).
			Data("teaMenu", "risk").
			Prefix("/hids/risk").
			Get("", new(IndexAction)).
			//系统漏洞
			GetPost("/systemRisk", new(SystemRiskAction)).
			Get("/systemRiskList", new(SystemRiskListAction)).
			Get("/riskDetail", new(RiskDetailAction)).
			//弱口令
			GetPost("/weak", new(WeakAction)).
			Get("/weakDetail", new(WeakDetailAction)).
			Get("/weakList", new(WeakListAction)).
			//风险账号
			GetPost("/dangerAccount", new(DangerAccountAction)).
			Get("/dangerAccountDetail", new(DangerAccountDetailAction)).
			Get("/dangerAccountList", new(DangerAccountListAction)).
			//配置缺陷
			GetPost("/configDefect", new(ConfigDefectAction)).
			Get("/configDefectDetail", new(ConfigDefectDetailAction)).
			Get("/configDefectList", new(ConfigDefectListAction)).
			EndAll()
	})
}
