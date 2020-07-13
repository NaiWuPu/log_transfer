package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// 初始化ES，准备接受kafka那边发来的数据
var (
	client *elastic.Client
	ch     chan *LogData
)

// Init ...
func Init(address string, chanSize, num int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}

	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < num; i++ {
		go Send2Es()
	}
	return err
}

func Send2EsChan(msg *LogData) {
	ch <- msg
}

func Send2Es() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Printf("send2ES err %v \n", err)
				continue
			}
			fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
