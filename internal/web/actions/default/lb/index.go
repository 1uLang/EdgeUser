package lb

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "index")
}

func (this *IndexAction) RunGet(params struct{}) {
	if !this.ValidateFeature("lb-tcp") {
		return
	}

	countResp, err := this.RPC().ServerRPC().CountAllEnabledServersMatch(this.UserContext(), &pb.CountAllEnabledServersMatchRequest{
		ServerGroupId:  0,
		Keyword:        "",
		UserId:         this.UserId(true),
		ProtocolFamily: "tcp",
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	count := countResp.Count
	page := this.NewPage(count)
	this.Data["page"] = page.AsHTML()

	serversResp, err := this.RPC().ServerRPC().ListEnabledServersMatch(this.UserContext(), &pb.ListEnabledServersMatchRequest{
		Offset:         page.Offset,
		Size:           page.Size,
		ServerGroupId:  0,
		Keyword:        "",
		ProtocolFamily: "tcp",
		UserId:         this.UserId(true),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverMaps := []maps.Map{}
	for _, server := range serversResp.Servers {
		// CNAME
		cname := ""
		if server.NodeCluster != nil {
			clusterId := server.NodeCluster.Id
			if clusterId > 0 {
				dnsInfoResp, err := this.RPC().NodeClusterRPC().FindEnabledNodeClusterDNS(this.UserContext(), &pb.FindEnabledNodeClusterDNSRequest{NodeClusterId: clusterId})
				if err != nil {
					this.ErrorPage(err)
					return
				}
				if dnsInfoResp.Domain != nil {
					cname = server.DnsName + "." + dnsInfoResp.Domain.Name + "."
				}
			}
		}

		// TCP
		tcpPorts := []string{}
		if len(server.TcpJSON) > 0 {
			config, err := serverconfigs.NewTCPProtocolConfigFromJSON(server.TcpJSON)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			if config.IsOn {
				for _, listen := range config.Listen {
					tcpPorts = append(tcpPorts, listen.PortRange)
				}
			}
		}

		// TLS
		tlsPorts := []string{}
		if len(server.TlsJSON) > 0 {
			config, err := serverconfigs.NewTLSProtocolConfigFromJSON(server.TlsJSON)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			if config.IsOn {
				for _, listen := range config.Listen {
					tlsPorts = append(tlsPorts, listen.PortRange)
				}
			}
		}

		serverMaps = append(serverMaps, maps.Map{
			"id":       server.Id,
			"name":     server.Name,
			"cname":    cname,
			"tcpPorts": tcpPorts,
			"tlsPorts": tlsPorts,
			"isOn":     server.IsOn,
		})
	}
	this.Data["servers"] = serverMaps

	this.Show()
}
