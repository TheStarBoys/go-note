package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/conf"
	"github.com/TheStarBoys/go-note/practice/demo/log_collection/tailf"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr, key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintln("connect failed, err:", err))
	}

	logs.Debug("etcd connect success")

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key += "/"
	}
}

func GetEtcdLogConf(key string) (collectConf []conf.CollectConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	for _, ip := range localIPs {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		logs.Debug("etcd find ip: %s", ip)
		resp, err := etcdClient.client.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd: %v", err)
			continue
		}

		for i, v := range resp.Kvs {
			logs.Debug("etcd cli get, i: %v, v: %v", i, v)
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("json unmarshal err: %v", err)
					continue
				}

				logs.Debug("log config is: %v", collectConf)
			}
		}
	}
	initEtcdWatch()

	logs.Debug("etcd cli load success")

	return
}

func initEtcdWatch() {
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

// 监听 etcd 中 key-value 的变化
func watchKey(key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()

	logs.Debug("etcd watch key: %s", key)
	for {
		rch := cli.Watch(context.Background(), key)

		for wresp := range rch {
			var (
				confs []conf.CollectConf
				isGetConfSucc = true
			)
			for _, event := range wresp.Events {
				if event.Type == mvccpb.DELETE {
					logs.Warn("key[%s]'s config deleted", key)
					continue
				}
				if event.Type == mvccpb.PUT && string(event.Kv.Key) == key {
					err = json.Unmarshal(event.Kv.Value, &confs)
					if err != nil {
						logs.Error("json unmarshal err: %v", err)
						isGetConfSucc = false
						continue
					}
				}
				logs.Debug("etcd watched key changed. %s %q : %q\n", event.Type, event.Kv.Key, event.Kv.Value)
			}

			if isGetConfSucc {
				tailf.UpdateConfig(confs)
			}
		}
	}
}