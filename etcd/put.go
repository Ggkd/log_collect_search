package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// 移到包外测试

func Put2(Client *clientv3.Client)  {
	// put
	defer Client.Close()
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Client.Put(ctx, "collect_log", `[{"path":"/Users/gongjiabao/Documents/gitdemo/log_collect/test_log/mysql.log","topic":"mysql_log"}]`)
	if err != nil {
		fmt.Printf("put %s err: %v\n", "log_path", err)
		return
	}
	fmt.Println("=====put log_path success=====")
}

func Put1(Client *clientv3.Client)  {
	// put
	defer Client.Close()
	ctx,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Client.Put(ctx, "collect_log", `[{"path":"/Users/gongjiabao/Documents/gitdemo/log_collect/test_log/web.log","topic":"web_log"}]`)
	if err != nil {
		fmt.Printf("put %s err: %v\n", "log_path", err)
		return
	}
	fmt.Println("=====put log_path success=====")
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*3,
	})
	if err != nil {
		fmt.Println("new client err :", err)
		return
	}
	fmt.Println("conn etcd success")

	//Put1(client)
	Put2(client)
}