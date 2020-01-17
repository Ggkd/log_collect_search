package taillog

import (
	"encoding/json"
	"fmt"
	"github.com/Ggkd/log_collect/etcd"
	"github.com/hpcloud/tail"
)


// tail初始化
func InitTail(config string) *tail.Tail {
	var putValue = new(etcd.PutValue)
	json.Unmarshal([]byte(config), putValue)
	t, err := tail.TailFile(putValue.Path, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("tail file err: ", err)
		return nil
	}
	fmt.Println("====init tail config success====")
	return t
}