package userutils


var (
	// 所有功能列表，注意千万不能在运行时进行修改
	allUserFeatures = []*UserFeature{
		{
			Name:        "记录访问日志",
			Code:        "server.accessLog",
			Description: "用户可以开启服务的访问日志",
		},
		{
			Name:        "查看访问日志",
			Code:        "server.viewAccessLog",
			Description: "用户可以查看服务的访问日志",
		},
		{
			Name:        "转发访问日志",
			Code:        "server.accessLog.forward",
			Description: "用户可以配置访问日志转发到自定义的API",
		},
		{
			Name:        "负载均衡",
			Code:        "server.tcp",
			Description: "用户可以添加TCP/TLS负载均衡服务",
		},
		{
			Name:        "自定义负载均衡端口",
			Code:        "server.tcp.port",
			Description: "用户可以自定义TCP端口",
		},
		{
			Name:        "开启WAF",
			Code:        "server.waf",
			Description: "用户可以开启WAF功能并可以设置黑白名单等",
		},
		{
			Name:        "费用账单",
			Code:        "finance",
			Description: "开启费用账单相关功能",
		},
		{
			Name:        "主机防护",
			Code:        "hids",
			Description: "开启主机防护组件功能",
		},
		{
			Name:        "漏洞扫描",
			Code:        "webscan",
			Description: "开启漏洞扫描组件功能",
		},
		{
			Name:        "堡垒机",
			Code:        "fort",
			Description: "开启堡垒机组件功能",
		},
		{
			Name:        "安全审计",
			Code:        "audit",
			Description: "开启安全审计组件功能",
		},
		{
			Name:        "数据备份",
			Code:        "databackup",
			Description: "开启数据备份组件功能",
		},
	}
)

// 用户功能
type UserFeature struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// 所有功能列表
func FindAllUserFeatures() []*UserFeature {
	return allUserFeatures
}

// 查询单个功能
func FindUserFeature(code string) *UserFeature {
	for _, feature := range allUserFeatures {
		if feature.Code == code {
			return feature
		}
	}
	return nil
}

