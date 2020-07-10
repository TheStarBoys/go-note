package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	"path"
)

func initLogger() {
	config := make(map[string]interface{})
	filename := path.Join(conf.GetWd(), conf.GetConfig().LogPath)

	config["filename"] = filename
	config["level"] = str2LogLevel(conf.GetConfig().LogLevel)

	js, err := json.Marshal(config)
	if err != nil {
		panic(fmt.Sprintf("initLogger failed, marshal err: %s", err))
	}

	logs.SetLogger(logs.AdapterFile, string(js))

}

func str2LogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}