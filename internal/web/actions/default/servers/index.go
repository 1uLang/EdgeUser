package servers

import (
	"encoding/json"
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
	countResp, err := this.RPC().ServerRPC().CountAllEnabledServersMatch(this.UserContext(), &pb.CountAllEnabledServersMatchRequest{
		ServerGroupId:        0,
		Keyword:        "",
		UserId:         this.UserId(),
		ProtocolFamily: "http",
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
		ServerGroupId:        0,
		Keyword:        "",
		ProtocolFamily: "http",
		UserId:         this.UserId(),
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	serverMaps := []maps.Map{}
	for _, server := range serversResp.Servers {
		// 域名列表
		serverNames := []*serverconfigs.ServerNameConfig{}
		if server.IsAuditing || (server.AuditingResult != nil && !server.AuditingResult.IsOk) {
			server.ServerNamesJSON = server.AuditingServerNamesJSON
		}
		if len(server.ServerNamesJSON) > 0 {
			err = json.Unmarshal(server.ServerNamesJSON, &serverNames)
			if err != nil {
				this.ErrorPage(err)
				return
			}
		}
		countServerNames := 0
		for _, serverName := range serverNames {
			if len(serverName.SubNames) == 0 {
				countServerNames++
			} else {
				countServerNames += len(serverName.SubNames)
			}
		}

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

		// HTTP
		httpIsOn := false
		if len(server.HttpJSON) > 0 {
			httpConfig, err := serverconfigs.NewHTTPProtocolConfigFromJSON(server.HttpJSON)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			httpIsOn = httpConfig.IsOn && len(httpConfig.Listen) > 0
		}

		// HTTPS
		httpsIsOn := false
		if len(server.HttpsJSON) > 0 {
			httpsConfig, err := serverconfigs.NewHTTPSProtocolConfigFromJSON(server.HttpsJSON)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			httpsIsOn = httpsConfig.IsOn && len(httpsConfig.Listen) > 0
		}

		serverMaps = append(serverMaps, maps.Map{
			"id":               server.Id,
			"serverNames":      serverNames,
			"countServerNames": countServerNames,
			"isAuditing":       false,
			"cname":            cname,
			"httpIsOn":         httpIsOn,
			"httpsIsOn":        httpsIsOn,
			"status": maps.Map{
				"isOk":    false,
				"message": "",
				"type":    "",
			},
			"isOn": server.IsOn,
		})
	}
	this.Data["servers"] = serverMaps

	this.Show()
}
