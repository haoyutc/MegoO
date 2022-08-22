package config

import (
	"fmt"
	"github.com/megoo/common/conf"
)

var (
	globalConfig *Config
)

type Config struct {
	conf ConfYaml
}

// NewConfig 创建配置对象
func NewConfig() *Config {
	if globalConfig == nil {
		globalConfig = &Config{}
		if err := globalConfig.init(); err != nil {
			panic(fmt.Errorf("init config fail err: %s", err.Error()))
		}
	}
	return globalConfig
}

//加载微服务配置
func (c *Config) init() error {
	// get the config file
	conf.Load(
		conf.Path("./common/config"),
		conf.File("config.yaml"),
		conf.Scan(&c.conf, Module),
	)
	return nil
}

// GetConf 读取配置
func (c *Config) GetConf() *ConfYaml {
	return &c.conf
}
