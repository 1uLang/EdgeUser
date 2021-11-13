package userutils

import (
	"github.com/1uLang/zhiannet-api/edgeUsers/model"
	"github.com/1uLang/zhiannet-api/edgeUsers/server"
	"github.com/iwind/TeaGo/lists"
)

var (
	// 所有功能列表，注意千万不能在运行时进行修改
	allUserFeatures = []*UserFeature{
		{
			Name:        "业务概览",
			Code:        "dashboard",
			Description: "当前用户用户平台所有业务数据汇总",
		},
		{
			Name:        "态势感知",
			Code:        "waf",
			Description: "用户平台整体地洞悉安全风险展示",
		},
		{
			Name:        "安全概览",
			Code:        "waf.waf",
			Description: "用户可以查看安全概览统计数据",
		},
		{
			Name:        "WAF日志",
			Code:        "waf.wafLogs",
			Description: "用户可以查看用户平台拦截日志",
		},
		{
			Name:        "DDoS日志",
			Code:        "waf.ddos",
			Description: "用户可以查看平台检测到的DDoS攻击记录",
		},
		{
			Name:        "IPS日志",
			Code:        "waf.alarm",
			Description: "用户可以查看云防火墙的安全监控记录",
		},
		{
			Name:        "DDoS防护",
			Code:        "ddos",
			Description: "用户开启DDoS防护功能",
		},
		{
			Name:        "主机状态",
			Code:        "ddos.host",
			Description: "用户可以查看DDoS防护的主机状态",
		},
		{
			Name:        "连接监控",
			Code:        "ddos.link",
			Description: "用户可以查看DDoS防护的连接监控",
		},
		{
			Name:        "屏蔽列表",
			Code:        "ddos.shield",
			Description: "用户可以查看DDoS防护的屏蔽列表",
		}, {
			Name:        "云防火墙",
			Code:        "nfw",
			Description: "用户开启云防火墙功能",
		},
		{
			Name:        "ACL规则",
			Code:        "nfw.acl",
			Description: "用户可以查阅云防火墙ACL规则",
		},
		{
			Name:        "IPS规则",
			Code:        "nfw.ips",
			Description: "用户可以查阅云防火墙IPS规则",
		},
		{
			Name:        "病毒库",
			Code:        "nfw.virus",
			Description: "用户可以查阅云防火墙的病毒库信息",
		},
		{
			Name:        "会话列表",
			Code:        "nfw.conversation",
			Description: "用户可以查阅云防火墙会话列表",
		},
		{
			Name:        "WAF服务",
			Code:        "servers",
			Description: "对自定义域名以及证书进行管理和预热",
		},
		{
			Name:        "域名管理",
			Code:        "servers.servers",
			Description: "用户可以新增、管理、删除域名",
		},
		{
			Name:        "证书管理",
			Code:        "servers.certs",
			Description: "用户可以上传各类的证书",
		},
		{
			Name:        "刷新预热",
			Code:        "servers.cache",
			Description: "用户可以配置URL刷新预热缓存",
		},
		{
			Name:        "证书服务",
			Code:        "certs",
			Description: "用户可以开启证书服务功能",
		},
		{
			Name:        "负载均衡",
			Code:        "lb-tcp-port",
			Description: "用户可以添加TCP/TLS负载均衡服务",
		},
		{
			Name:        "自定义负载均衡端口",
			Code:        "lb-tcp-port",
			Description: "用户可以自定义TCP端口",
		},
		{
			Name:        "主机防护",
			Code:        "hids",
			Description: "开启主机防护组件功能",
		},
		{
			Name:        "主机体检",
			Code:        "hids.examine",
			Description: "开启主机防护主机体检功能",
		},
		{
			Name:        "漏洞风险",
			Code:        "hids.risk",
			Description: "用户可以查看主机体检后可能检测出来的漏洞风险",
		},
		{
			Name:        "入侵威胁",
			Code:        "hids.invade",
			Description: "用户可以查看主机体检后可能检测出来的入侵威胁",
		},
		{
			Name:        "合规基线",
			Code:        "hids.baseline",
			Description: "用户可以开启主机防护等保合规基线检测功能",
		},
		{
			Name:        "Agent管理",
			Code:        "hids.agent",
			Description: "用户可以安装、添加agent主机功能",
		},
		{
			Name:        "端点防护",
			Code:        "nhids",
			Description: "开启端点防护组件功能",
		},
		{
			Name:        "资产管理",
			Code:        "nhids.agents",
			Description: "开启端点防护主机资产管理",
		},
		{
			Name:        "安全事件",
			Code:        "nhids.attck",
			Description: "用户可以查看检测出来的安全事件记录",
		},
		//{
		//	Name:        "漏洞风险",
		//	Code:        "nhids.vulnerability",
		//	Description: "用户可以查看检测出来的漏洞风险记录",
		//},
		{
			Name:        "病毒检测",
			Code:        "nhids.virus",
			Description: "用户可以查看检测出来的病毒检测记录",
		},
		{
			Name:        "合规基线",
			Code:        "nhids.baseLine",
			Description: "用户可以查看检测出来的合规基线记录",
		},
		{
			Name:        "文件监控",
			Code:        "nhids.syscheck",
			Description: "用户可以查看主机文件列表以及文件监控记录",
		},
		{
			Name:        "漏洞扫描",
			Code:        "webscan",
			Description: "开启漏洞扫描组件功能",
		},
		{
			Name:        "扫描目标",
			Code:        "webscan.examine",
			Description: "用户可以创建web、主机漏洞扫描目标功能",
		},
		{
			Name:        "扫描任务",
			Code:        "webscan.risk",
			Description: "用户可以对web、主机漏洞扫描目标开启行扫描已经生成报告功能",
		},
		{
			Name:        "扫描报告",
			Code:        "webscan.reports",
			Description: "用户可以下载漏洞扫描报告文件",
		},
		{
			Name:        "云堡垒机",
			Code:        "fortcloud",
			Description: "开启堡垒机组件功能",
		},
		{
			Name:        "资产管理",
			Code:        "fortcloud.assets",
			Description: "用户可以创建授权连接资产",
		},
		{
			Name:        "授权凭证",
			Code:        "fortcloud.cert",
			Description: "用户可以创建授权登录资产的授权凭证",
		},
		{
			Name:        "会话管理",
			Code:        "fortcloud.sessions",
			Description: "用户可以主动断开或监控资产连接在线会话",
		},
		{
			Name:        "运维审计",
			Code:        "fortcloud.audit",
			Description: "用户可以回放资产连接历史会话",
		},
		{
			Name:        "安全审计",
			Code:        "audit",
			Description: "开启安全审计组件功能",
		},
		{
			Name:        "数据库管理",
			Code:        "audit.db",
			Description: "用户拥有针对数据库的安全审计功能",
		},
		{
			Name:        "主机管理",
			Code:        "audit.host",
			Description: "用户拥有针对主机的安全审计功能",
		},
		{
			Name:        "应用管理",
			Code:        "audit.app",
			Description: "用户拥有针对应用的安全审计功能",
		},
		{
			Name:        "审计日志",
			Code:        "audit.logs",
			Description: "用户可以查看该用户下所有资产的审计日志",
		},
		{
			Name:        "订阅报告",
			Code:        "audit.report",
			Description: "用户可以配置资产安全审计的订阅报告",
		},
		{
			Name:        "Agent管理",
			Code:        "audit.agent",
			Description: "用户可以下载各类型的安全审计Agent",
		},
		{
			Name:        "数据备份",
			Code:        "databackup",
			Description: "开启数据备份组件功能",
		},
		{
			Name:        "平台管理",
			Code:        "platform",
			Description: "配置子账号以及平台的安全策略和查看用户以及子账号的操作日志",
		},
		{
			Name:        "账号管理",
			Code:        "platform.user",
			Description: "用户可以新增、删除子账号以及配置其权限",
		},
		{
			Name:        "创建子账号",
			Code:        "platform.userCreate",
			Description: "用户可以新增子账号",
		},
		{
			Name:        "操作日志",
			Code:        "platform.logs",
			Description: "用户可以查看用户及其子账号的操作日志",
		},
		{
			Name:        "安全策略",
			Code:        "platform.strategy",
			Description: "用户可以设置用户平台的安全策略",
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

func CheckCreateIsOn(userId int64, checkCode string) (bool, error) {

	userFeatureCodes, err := server.FindUserFeatures(&model.FindUserFeaturesReq{UserId: userId})
	if err != nil {
		return false, err
	}
	return lists.ContainsString(userFeatureCodes, checkCode), nil
}
