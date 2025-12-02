package config

import (
	"errors"

	"myreel/pkg/constants"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/spf13/viper"
)

var (
	Server       *server
	Mysql        *mySQL
	Snowflake    *snowflake
	Redis        *redis
	Service      *service
	Upyun        *upyun
	runtimeViper = viper.New()
)

const (
	configPath     = "./config"
	configFileName = "config"
	configFileType = "yaml"
)

func Init(service string) {
	runtimeViper.AddConfigPath(configPath)
	runtimeViper.SetConfigName(configFileName)
	runtimeViper.SetConfigType(configFileType)
	if err := runtimeViper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Fatal("config.Init: could not find config files")
		}
		logger.Fatalf("config.Init: read config error: %v", err)
	}
	configMapping(service)
}

func configMapping(srv string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		// 由于这个函数会在配置重载时被再次触发，所以需要判断日志记录方式
		logger.Fatalf("config.configMapping: config: unmarshal error: %v", err)
	}
	Snowflake = &c.Snowflake
	Server = &c.Server
	Mysql = &c.MySQL
	Redis = &c.Redis
	Upyun = &c.Upyun
	Service = getService(srv)
}

func getService(name string) *service {
	addrList := runtimeViper.GetStringSlice("services." + name + ".addr")

	return &service{
		Name:     runtimeViper.GetString("services." + name + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + name + ".load-balance"),
	}
}

func GetDataCenterID() int64 {
	if Snowflake == nil {
		return constants.DefaultDataCenterID
	}
	return Snowflake.DatacenterID
}

func GetLoggerLevel() string {
	if Server == nil {
		return constants.DefaultLogLevel
	}
	return Server.LogLevel
}
