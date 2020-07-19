package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	etcd_client "github.com/coreos/etcd/clientv3"
	"path/filepath"
	"time"
)

const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.2.194"
)

func main() {
	SetLogConf2Etcd()
}

func SetLogConf2Etcd() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	logConfArr := []conf.CollectConf{
		{
			LogPath: filepath.Join("D:/Development/Go/src/github.com/TheStarBoys/go-note/practice/demo/log_collection/", "logs/logagent.log"),
			Topic: "nginx_log",
		},
		{
			LogPath:  filepath.Join("D:/Development/Go/src/github.com/TheStarBoys/go-note/practice/demo/log_collection/", "logs/collection.log"),
			Topic: "nginx_log_err",
		},
	}

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
