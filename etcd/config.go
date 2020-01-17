package etcd

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	Etcd		`ini:"etcd"`
}

type Etcd struct {
	Ip 			string		`ini:"ip"`
	Port 		string		`ini:"port"`
	DialTimeout	int			`ini:"DialTimeout"`
	Key			string		`ini:"key"`
	ChanSize	int			`ini:"chanSize"`
}

var EtcdConfig = new(Config)

func LoadConfig() error {
	err := ini.MapTo(EtcdConfig, "./config/conf.ini")
	return err
}