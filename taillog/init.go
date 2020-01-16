package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func InitTail(config *Config) *tail.Tail {
	t, err := tail.TailFile(config.Tail.Path, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("tail file err: ", err)
		return nil
	}
	fmt.Println("====init tail config success===")
	return t
}