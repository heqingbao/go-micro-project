package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	DB   string `json:"db"`
	Port int64  `json:"port"`
}

// GetMysqlFromConsul 获取 mysql 配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}