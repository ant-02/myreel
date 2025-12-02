package config

import "time"

type server struct {
	Secret      string `mapstructure:"private-key"`
	PublicKey   string `mapstructure:"public-key"`
	Version     string
	Name        string
	LogLevel    string `mapstructure:"log-level"`
	IntranetUrl string `mapstructure:"intranet-url"`
}

type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type redis struct {
	Addr     string
	Password string
}

type snowflake struct {
	DatacenterID int64 `mapstructure:"datacenter-id"`
}

type config struct {
	Server    server
	Snowflake snowflake
	MySQL     mySQL
	Redis     redis
	Upyun     upyun
}

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}

type upyun struct {
	Bucket     string
	Operator   string
	Password   string
	Expiration time.Duration
	MaxSize    int64
	Domain     string
	NotifyUrl  string
}
