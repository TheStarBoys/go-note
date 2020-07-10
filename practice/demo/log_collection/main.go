package main

import (
	"path"
	"github.com/astaxie/beego/logs"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/tailf"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/kafka"
)

func main() {
	filename := path.Join(conf.GetWd(), "/conf/logagent.conf")
	conf.LoadConf("ini", filename)

	initLogger()
	logs.Debug("load conf success, config: %#v", conf.GetConfig())

	tailf.InitTail(conf.GetConfig().CollectConf, conf.GetConfig().ChanSize)

	kafka.InitKafka(conf.GetConfig().KafkaAddr)

	logs.Debug("initialize success")

	//go func() {
	//	count := 0
	//	for {
	//		count++
	//		logs.Debug("test for logger: %d", count)
	//		time.Sleep(400 * time.Millisecond)
	//	}
	//}()

	Run()

	logs.Info("program exist")
}

func Run() {
	for {
		msg, ok := tailf.GetLine()
		if !ok {
			logs.Warn("tailf.GetLine failed")
			continue
		}
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err: %s", err)
			continue
		}
	}
}

func sendToKafka(msg *tailf.TextMsg) error {
	//fmt.Printf("read msg: %#v\n", msg)
	return kafka.Send(msg.Msg, msg.Topic)
}