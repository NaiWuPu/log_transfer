package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log_transfer/conf"
	"log_transfer/es"
	"log_transfer/kafka"
	"log_transfer/vendor/gopkg.in/ini.v1"
)
// 将日志数据从kafka 取出来发往ES
func main()  {
	// 加载配置文件
	var cfg = new(conf.LogTransfer)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config failed:%v \n", err)
		return
	}
	fmt.Printf("cfg:%v\n",cfg)
	// 初始化ES
	// 初始化一个ES的链接 client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init Es client failed, err %v\n", err)
		return
	}
	// 初始化kafka
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err:%v\n", err)
		return
	}

	select {

	}

}
