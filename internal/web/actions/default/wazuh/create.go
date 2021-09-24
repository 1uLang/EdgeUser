// 主机防护使用wazuh组件

package wazuh

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/server"
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

	userIdGroup := fmt.Sprintf("user_%v", this.UserId(true))
	_ = server.CreateGroup(userIdGroup)

	this.Data["group"] = userIdGroup

	addr := strings.Replace(serverAddr, "http://", "", 1)
	addr = strings.Replace(addr, "https://", "", 1)

	this.Data["server"] = addr

	this.Data["installs"] = maps.Map{
		"1": maps.Map{
			"1": maps.Map{
				"1": "hids-agent.el5.i386.rpm",
				"2": "hids-agent.el5.x86_64.rpm",
			},
			"2": maps.Map{
				"1": "hids-agent.i386.rpm",
				"2": "hids-agent.x86_64.rpm",
				"3": "hids-agent.armv7hl.rpm",
				"4": "hids-agent.aarch64.rpm",
			},
		},
		"2": maps.Map{
			"1": "hids-agent_i386.deb",
			"2": "hids-agent_amd64.deb",
			"3": "hids-agent_armhf.deb",
			"4": "hids-agent_arm64.deb",
		},
		"3": "hids-agent.msi",
		"4": "hids-agent.pkg",
	}
	this.Data["commands"] = maps.Map{
		"1": "sudo systemctl daemon-reload\nsudo systemctl enable wazuh-agent\nsudo systemctl start wazuh-agent",
		"2": "sudo systemctl daemon-reload\nsudo systemctl enable wazuh-agent\nsudo systemctl start wazuh-agent",
		"3": "",
		"4": "sudo /Library/Ossec/bin/wazuh-control start",
	}
	this.Data["uninstalls"] = maps.Map{
		"1": "yum remove wazuh-agent",
		"2": "dpkg -P wazuh-agent",
		"3": "msiexec.exe /x hids-agent.msi /qn",
		"4": "",
	}
	this.Show()
}
