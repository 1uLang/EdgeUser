package web

import (
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl/accesskeys"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/csrf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/dashboard"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance/bills"

	//主机安全防护
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/agent"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/hids/baseline"
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
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf/logs"

	//web漏洞扫描
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/reports"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/scans"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/targets"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/webscan/vulnerabilities"
)
