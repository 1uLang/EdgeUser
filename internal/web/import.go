package web

import (
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl/accesskeys"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/csrf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/dashboard"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance/bills"

	//堡垒机
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud/assets"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud/audit"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud/cert"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/fortcloud/sessions"

	//主机安全防护
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/agent"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/baseline"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/bwlist"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/examine"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/abnormalAccount"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/abnormalLogin"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/abnormalProcess"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/logDelete"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/reboundShell"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/systemCmd"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/virus"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/invade/webShell"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/risk"

	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/index"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/basic"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/dns"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/reverseProxy"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/tcp"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/tls"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/logout"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/messages"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/cache"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/certs"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/components/waf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/log"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/accessLog"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/cache"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/charset"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/conds"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/dns"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/headers"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/http"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/https"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/origins"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/redirects"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/reverseProxy"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/serverNames"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/stat"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/waf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/websocket"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/stat"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/stat"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/login"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/profile"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/ui"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf"
	//_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf/logs"

	//web漏洞扫描
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/reports"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/scans"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/targets"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/vulnerabilities"

	//安全审计
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/agent"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/app"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/db"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/host"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/logs"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/audit/report"

	//数据备份
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/databackup"

	//ddos攻击
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/ddos"

	//ddos攻击
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/nfw/acl"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/nfw/conversation"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/nfw/ips"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/nfw/virus"

	//平台管理
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/ip_white"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/logs"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/strategy"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/platform/user"
	//APT
	//_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/apt"

	//主机列表
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hostlist"

	//主机安全防护
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/wazuh"
)
