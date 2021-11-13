package servers

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
	"regexp"
)

type AddPortPopupAction struct {
	actionutils.ParentAction
}

func (this *AddPortPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *AddPortPopupAction) RunGet(params struct {
	ServerType string
	Protocol   string
}) {
	protocols := serverconfigs.AllServerProtocolsForType(params.ServerType)
	if len(params.Protocol) > 0 {
		result := []maps.Map{}
		for _, p := range protocols {
			if p.GetString("code") == params.Protocol {
				result = append(result, p)
			}
		}
		protocols = result
	}
	this.Data["protocols"] = protocols

	this.Show()
}

func (this *AddPortPopupAction) RunPost(params struct {
	Protocol string
	Address  string

	Must *actions.Must
}) {
	// 校验地址
	addr := maps.Map{
		"protocol":  params.Protocol,
		"host":      "",
		"portRange": "",
	}

	digitRegexp := regexp.MustCompile(`^\d+$`)
	if !digitRegexp.MatchString(params.Address) {
		this.Fail("端口号只能是一个数字")
	}

	port := types.Int32(params.Address)
	if port < 1024 || port > 65534 {
		this.Fail("端口范围错误")
	}

	addr["portRange"] = params.Address

	this.Data["address"] = addr
	this.Success()
}
