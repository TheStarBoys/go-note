package conf

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"os"
	"path"
)

var (
	appConfig *Config
	currWd string
)

func init() {
	dir, _ := os.Getwd()
	currWd = path.Join(dir, "practice/demo/log_collection")
}

type Config struct {
	LogLevel    string
	LogPath     string
	ChanSize 	int
	KafkaAddr   string
	CollectConf []CollectConf
}

type CollectConf struct {
	LogPath string
	Topic   string
}

func GetConfig() *Config {
	return appConfig
}

func GetWd() string {
	return currWd
}

func LoadConf(confType, filename string) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		panic(fmt.Sprintf("new config failed, err: %s", err))
	}

	appConfig = &Config{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}

	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "logs/logagent.log"
	}

	appConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.KafkaAddr) == 0 {
		panic(fmt.Sprintf("invalid kafka addr"))
	}

	appConfig.ChanSize, err = conf.Int("collection::chan_size")
	if err != nil {
		appConfig.ChanSize = 100
	}

	loadCollectionConf(conf)

	return
}

func loadCollectionConf(conf config.Configer) {
	var cc CollectConf
	cc.LogPath = conf.String("collection::log_path")
	if len(cc.LogPath) == 0 {
		panic(fmt.Sprintf("invalid collection::log_path"))
	}
	cc.LogPath = path.Join(GetWd(), cc.LogPath)

	cc.Topic = conf.String("collection::Topic")
	if len(cc.Topic) == 0 {
		panic(fmt.Sprintf("invalid collection::Topic"))
	}

	appConfig.CollectConf = append(appConfig.CollectConf, cc)
}