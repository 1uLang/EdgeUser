// 主机防护使用wazuh组件

package wazuh

import (
	"fmt"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
	"strings"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "baseline")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Data["group"] = fmt.Sprintf("user_%v", this.UserId(true))

	addr := strings.Replace(serverAddr, "http://", "", 1)
	addr = strings.Replace(addr, "https://", "", 1)

	this.Data["server"] = addr

	this.Data["commands"] = maps.Map{
		"1": maps.Map{
			"1": maps.Map{
				"1": "zhianhids-agent.el5.i386.rpm",
				"2": "zhianhids-agent.el5.x86_64.rpm",
			},
			"2": maps.Map{
				"1": "zhianhids-agent.i386.rpm",
				"2": "zhianhids-agent.x86_64.rpm",
				"3": "zhianhids-agent.armv7hl.rpm",
				"4": "zhianhids-agent.aarch64.rpm",
			},
		},
		"2": maps.Map{
			"1": "zhianhids-agent_i386.deb",
			"2": "zhianhids-agent_amd64.deb",
			"3": "zhianhids-agent_armhf.deb",
			"4": "zhianhids-agent_arm64.deb",
		},
		"3": "zhianhids-agent.msi",
		"4": "zhianhids-agent.pkg",
	}
	this.Show()
}
