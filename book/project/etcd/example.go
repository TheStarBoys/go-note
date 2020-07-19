package main

import (
	"time"
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"context"
	"encoding/json"
)

type EtcdClient struct {
	client *etcd_client.Client
}

var (
	etcdClient *EtcdClient
)

func main() {
	SetLogConf2Etcd()
	//EtcdExample()

}

func EtcdExample() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

const (
	// ip 换成自己的ip
	EtcdKey = "/oldboy/backend/logagent/config/192.168.2.194"
)

type LogConf struct {
	Path string `json:"log_path"`
	Topic string `json:"topic"`
}

func SetLogConf2Etcd() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	logConfArr := []LogConf{
		{
			Path:  "D:/project/nginx/logs/access.log",
			Topic: "nginx_log",
		},
		{
			Path:  "D:/project/nginx/logs/error.log",
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