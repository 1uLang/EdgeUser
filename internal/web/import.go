package web

import (
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/cache"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/csrf"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/dashboard"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/finance/bills"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/index"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/logout"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/messages"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/servers/certs"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/login"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/settings/profile"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/ui"
	_ "github.com/TeaOSLab/EdgeUser/internal/web/actions/default/waf"
)
