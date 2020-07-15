## kafka 消费端日志 推送入ES
### 未使用 influshDB 原因是 influshDB在一定规模之后必须要使用企业版
### ElasticSearch 及 kibana 安装 https://www.elastic.co/cn/downloads/

######### 目录结构描述
``` 
├─.idea
├─conf                          // 配置文件
├─es                            // 初始化Es 暴露写入Es 接口
├─kafka                         // 初始化kafka 并开启暴露
└─vendor                        // 三方包
```

#### 运行 go run main.go