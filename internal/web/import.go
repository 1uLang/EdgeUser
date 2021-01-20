package web

import (
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/acl/accesskeys"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/csrf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/dashboard"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance/bills"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/index"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/basic"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/dns"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/lb/server/settings/reverseProxy"
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
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/waf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/server/settings/websocket"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/stat"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/login"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/profile"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/ui"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf/logs"
)
