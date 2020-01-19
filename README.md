## log_collect_search
使用etcd、kafka、tail、elastic 实现日志配置中心、日志收集、日志搜索监控

通过etcd获取日志的配置参数(日志的path和kafka的topic)，并通过watch实时监控最新的配置参数；

ectd将获取的配置发往tail，tail根据配置打开对应的日志文件，且一直获取日志文件的每一行数据；

tail将获取的日志数据发往kafka，kafka的生产者将会在后台生产日志数据；

kafka的消费者将会消费当前topic的数据，将数据发往es；

es收到日志数据后，把数据写入存储，可通过kibana搜索查询数据;

可使用Prometheus进行监控;
