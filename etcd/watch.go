package etcd

// 监控数据的变化，并将数据发往通道  TailConfigChan
//func Watch() {
//	watchChan := Client.Watch(context.Background(), EtcdConfig.Key)
//	for {
//		for event := range watchChan {
//			for _, ev := range event.Events {
//				fmt.Printf("type:%s,  key:%s, value:%s \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
//				ValueChan <- string(ev.Kv.Value)
//				//config := string(ev.Kv.Value)
//				//t := taillog.InitTail(config)
//				//tailConfig :=  &taillog.TailAndConfig{t, config}
//				//TailConfigChan <- tailConfig
//			}
//		}
//	}
//}


// 从通道取数据,当能取出数据时，开启新的tailObj和配置新的config，关闭原来的tailObj
//func GetValue() string {
//	for {
//		select {
//		case value := <- ValueChan:
//			taillog.InitTail()
//			return value
//		default:
//			time.Sleep(time.Microsecond)
//		}
//	}
//}