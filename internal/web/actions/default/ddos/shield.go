package host

import (
	"fmt"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
)

type ShieldAction struct {
	actionutils.ParentAction
}

func (this *ShieldAction) Init() {
	this.Nav("", "", "")
}

func (this *ShieldAction) RunGet(params struct {
	Addr   string
	NodeId uint64
}) {
	if params.NodeId == 0 {
		this.ErrorPage(fmt.Errorf("请选择节点"))
		return
	}
	if len(params.Addr) == 0 {
		this.ErrorPage(fmt.Errorf("请选择释放屏蔽的主机"))
		return
	}
	list, err := host_status_server.GetHostShieldList(&host_status_server.ShieldReq{
		Addr:   params.Addr,
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["shield"] = list.Fblink
	// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "释放屏蔽主机 %v", params.Addr)

	this.Success()
}

func (this *ShieldAction) RunPost(params struct {
	Addr   []string
	NodeId uint64
}) {
	if params.NodeId == 0 {
		this.ErrorPage(fmt.Errorf("请选择节点"))
		return
	}
	if len(params.Addr) == 0 {
		this.ErrorPage(fmt.Errorf("请选择释放屏蔽的主机"))
		return
	}
	//GetHostShieldList
	err := host_status_server.ReleaseShield(&host_status_server.ReleaseShieldReq{
		Addr:   params.Addr,
		NodeId: params.NodeId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建日志
	//defer this.CreateLog(oplogs.LevelInfo, "释放屏蔽主机 %v", params.Addr)

	this.Success()
}
