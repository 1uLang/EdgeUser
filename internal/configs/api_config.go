package configs

import (
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo/Tea"
	"io/ioutil"
)

var SharedAPIConfig *APIConfig

// API配置
type APIConfig struct {
	RPC struct {
		Endpoints []string `yaml:"endpoints"`
	} `yaml:"rpc"`
	NodeId   string `yaml:"nodeId"`
	Secret   string `yaml:"secret"`
	NumberId int64  `yaml:"numberId"`
}

// 加载API配置
func LoadAPIConfig() (*APIConfig, error) {
	if SharedAPIConfig != nil {
		return SharedAPIConfig, nil
	}

	data, err := ioutil.ReadFile(Tea.ConfigFile("api.yaml"))
	if err != nil {
		return nil, err
	}

	config := &APIConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	SharedAPIConfig = config

	return config, nil
}

// 写入API配置
func (this *APIConfig) WriteFile(path string) error {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0666)
}
