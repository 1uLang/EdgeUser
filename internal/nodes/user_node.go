package nodes

import (
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"github.com/TeaOSLab/EdgeUser/internal/configs"
	teaconst "github.com/TeaOSLab/EdgeUser/internal/const"
	"github.com/TeaOSLab/EdgeUser/internal/events"
	"github.com/TeaOSLab/EdgeUser/internal/rpc"
	"github.com/TeaOSLab/EdgeUser/internal/utils"
	_ "github.com/TeaOSLab/EdgeUser/internal/web"
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/sessions"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func NewUserNode() *UserNode {
	return &UserNode{}
}

type UserNode struct {
}

func (this *UserNode) Run() {
	// 启动用户界面
	secret := this.genSecret()
	configs.Secret = secret

	// 本地Sock
	err := this.listenSock()
	if err != nil {
		logs.Println("[USER_NODE]" + err.Error())
		return
	}

	// 检查server配置
	err = this.checkServer()
	if err != nil {
		logs.Println("[USER_NODE]" + err.Error())
		return
	}

	// 触发事件
	events.Notify(events.EventStart)

	// 拉取配置
	err = this.pullConfig()
	if err != nil {
		logs.Println("[USER_NODE]pull config failed: " + err.Error())
		return
	}

	// 启动Web服务
	TeaGo.NewServer(false).
		AccessLog(false).
		EndAll().
		Session(sessions.NewFileSessionManager(86400, secret), teaconst.CookieSID).
		ReadHeaderTimeout(3 * time.Second).
		ReadTimeout(600 * time.Second).
		Start()
}

// 实现守护进程
func (this *UserNode) Daemon() {
	path := os.TempDir() + "/edge-user.sock"
	isDebug := lists.ContainsString(os.Args, "debug")
	isDebug = true
	for {
		conn, err := net.DialTimeout("unix", path, 1*time.Second)
		if err != nil {
			if isDebug {
				log.Println("[DAEMON]starting ...")
			}

			// 尝试启动
			err = func() error {
				exe, err := os.Executable()
				if err != nil {
					return err
				}
				cmd := exec.Command(exe)
				err = cmd.Start()
				if err != nil {
					return err
				}
				err = cmd.Wait()
				if err != nil {
					return err
				}
				return nil
			}()

			if err != nil {
				if isDebug {
					log.Println("[DAEMON]", err)
				}
				time.Sleep(1 * time.Second)
			} else {
				time.Sleep(5 * time.Second)
			}
		} else {
			_ = conn.Close()
			time.Sleep(5 * time.Second)
		}
	}
}

// 安装系统服务
func (this *UserNode) InstallSystemService() error {
	shortName := teaconst.SystemdServiceName

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	manager := utils.NewServiceManager(shortName, teaconst.ProductName)
	err = manager.Install(exe, []string{})
	if err != nil {
		return err
	}
	return nil
}

// 检查Server配置
func (this *UserNode) checkServer() error {
	configFile := Tea.ConfigFile("server.yaml")
	_, err := os.Stat(configFile)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		// 创建文件
		templateFile := Tea.ConfigFile("server.template.yaml")
		data, err := ioutil.ReadFile(templateFile)
		if err == nil {
			err = ioutil.WriteFile(configFile, data, 0666)
			if err != nil {
				return errors.New("create config file failed: " + err.Error())
			}
		} else {
			templateYAML := `# environment code
env: prod

# http
http:
  "on": true
  listen: [ "0.0.0.0:7789" ]

# https
https:
  "on": false
  listen: [ "0.0.0.0:443"]
  cert: ""
  key: ""
`
			err = ioutil.WriteFile(configFile, []byte(templateYAML), 0666)
			if err != nil {
				return errors.New("create config file failed: " + err.Error())
			}
		}
	} else {
		return errors.New("can not read config from 'configs/server.yaml': " + err.Error())
	}

	return nil
}

// 生成Secret
func (this *UserNode) genSecret() string {
	tmpFile := os.TempDir() + "/edge-user-secret.tmp"
	data, err := ioutil.ReadFile(tmpFile)
	if err == nil && len(data) == 32 {
		return string(data)
	}
	secret := rands.String(32)
	_ = ioutil.WriteFile(tmpFile, []byte(secret), 0666)
	return secret
}

// 拉取配置
func (this *UserNode) pullConfig() error {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		return err
	}
	nodeResp, err := rpcClient.UserNodeRPC().FindCurrentUserNode(rpcClient.Context(0), &pb.FindCurrentUserNodeRequest{})
	if err != nil {
		return err
	}
	node := nodeResp.Node
	if node == nil {
		return errors.New("invalid 'nodeId' or 'secret'")
	}

	// 读取Web服务配置
	serverConfig := &TeaGo.ServerConfig{
		Env: Tea.EnvProd,
	}
	if Tea.IsTesting() {
		serverConfig.Env = Tea.EnvDev
	}

	// HTTP
	httpConfig, err := this.decodeHTTP(node)
	if err != nil {
		return errors.New("decode http config failed: " + err.Error())
	}
	if httpConfig != nil && httpConfig.IsOn && len(httpConfig.Listen) > 0 {
		serverConfig.Http.On = true

		listens := []string{}
		for _, listen := range httpConfig.Listen {
			listens = append(listens, listen.Addresses()...)
		}
		serverConfig.Http.Listen = listens
	}

	// HTTPS
	httpsConfig, err := this.DecodeHTTPS(node)
	if err != nil {
		return errors.New("decode https config failed: " + err.Error())
	}
	if httpsConfig != nil && httpsConfig.IsOn && len(httpsConfig.Listen) > 0 {
		serverConfig.Https.On = true
		serverConfig.Https.Cert = "configs/https.cert.pem"
		serverConfig.Https.Key = "configs/https.key.pem"

		listens := []string{}
		for _, listen := range httpsConfig.Listen {
			listens = append(listens, listen.Addresses()...)
		}
		serverConfig.Https.Listen = listens
	}

	// 保存到文件
	serverYAML, err := yaml.Marshal(serverConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Tea.ConfigFile("server.yaml"), serverYAML, 0666)
	if err != nil {
		return err
	}

	return nil
}

// 解析HTTP配置
func (this *UserNode) decodeHTTP(node *pb.UserNode) (*serverconfigs.HTTPProtocolConfig, error) {
	if len(node.HttpJSON) == 0 {
		return nil, nil
	}
	config := &serverconfigs.HTTPProtocolConfig{}
	err := json.Unmarshal(node.HttpJSON, config)
	if err != nil {
		return nil, err
	}

	err = config.Init()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// 解析HTTPS配置
func (this *UserNode) DecodeHTTPS(node *pb.UserNode) (*serverconfigs.HTTPSProtocolConfig, error) {
	if len(node.HttpsJSON) == 0 {
		return nil, nil
	}
	config := &serverconfigs.HTTPSProtocolConfig{}
	err := json.Unmarshal(node.HttpsJSON, config)
	if err != nil {
		return nil, err
	}

	err = config.Init()
	if err != nil {
		return nil, err
	}

	if config.SSLPolicyRef != nil {
		policyId := config.SSLPolicyRef.SSLPolicyId
		if policyId > 0 {
			rpcClient, err := rpc.SharedRPC()
			if err != nil {
				return nil, err
			}
			policyConfigResp, err := rpcClient.SSLPolicyRPC().FindEnabledSSLPolicyConfig(rpcClient.Context(0), &pb.FindEnabledSSLPolicyConfigRequest{SslPolicyId: policyId})
			if err != nil {
				return nil, err
			}
			if len(policyConfigResp.SslPolicyJSON) > 0 {
				policyConfig := &sslconfigs.SSLPolicy{}
				err = json.Unmarshal(policyConfigResp.SslPolicyJSON, policyConfig)
				if err != nil {
					return nil, err
				}
				if len(policyConfig.Certs) > 0 {
					err = ioutil.WriteFile(Tea.ConfigFile("https.cert.pem"), policyConfig.Certs[0].CertData, 0666)
					if err != nil {
						return nil, err
					}

					err = ioutil.WriteFile(Tea.ConfigFile("https.key.pem"), policyConfig.Certs[0].KeyData, 0666)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	err = config.Init()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// 监听本地sock
func (this *UserNode) listenSock() error {
	path := os.TempDir() + "/edge-user.sock"

	// 检查是否已经存在
	_, err := os.Stat(path)
	if err == nil {
		conn, err := net.Dial("unix", path)
		if err != nil {
			_ = os.Remove(path)
		} else {
			_ = conn.Close()
		}
	}

	// 新的监听任务
	listener, err := net.Listen("unix", path)
	if err != nil {
		return err
	}
	events.On(events.EventQuit, func() {
		logs.Println("USER_NODE", "quit unix sock")
		_ = listener.Close()
	})

	go func() {
		for {
			_, err := listener.Accept()
			if err != nil {
				return
			}
		}
	}()

	return nil
}
