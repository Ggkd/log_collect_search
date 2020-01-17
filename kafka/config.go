package kafka

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	Kafka	`ini:"kafka"`
}

type Kafka struct {
	Ip 			string 		`ini:"ip"`
	Port 		string		`ini:"port"`
	ChanSize	int			`ini:"chanSize"`
}

func LoadConfig() *Config {
	config := new(Config)
	err := ini.MapTo(config, "./config/conf.ini")
	if err != nil {
		fmt.Println("ini config err:", err)
		return nil
	}
	fmt.Println("====load kafka config success====")
	return config
}