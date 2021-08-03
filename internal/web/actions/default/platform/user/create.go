package user

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "platform", "create")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	Name        string
	Protocols   []string
	CertIdsJSON []byte
	OriginsJSON []byte
	TcpPorts    []int
	TlsPorts    []int

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	// 检查用户所在集群

	//defer this.CreateLogInfo("创建TCP负载均衡服务 %d", serverId)

	this.Success()
}
