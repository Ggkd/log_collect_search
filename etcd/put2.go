package etcd

import (
	"context"
	"fmt"
	"time"
)

func Put2()  {
	// put
	defer Client.Close()
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Client.Put(ctx, EtcdConfig.Key, `{"path":"/Users/gongjiabao/Documents/gitdemo/log_collect/test_log/mysql.log","topic":"mysql_log"}`)
	if err != nil {
		fmt.Printf("put %s err: %v\n", EtcdConfig.Key, err)
		return
	}
	fmt.Println("=====put "+EtcdConfig.Key +" success=====")
}

