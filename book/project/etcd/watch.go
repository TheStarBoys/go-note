package main

import (
	"time"
	etcd_client "github.com/coreos/etcd/clientv3"
	"fmt"
	"context"
)

func main() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	fmt.Println("connect succ")

	watch(cli)
}

// 监听 etcd 中 key-value 的变化
func watch(cli *etcd_client.Client) {
	for {
		rch := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range rch {
			for _, event := range wresp.Events {
				fmt.Printf("%s %q : %q\n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}
	}
}
