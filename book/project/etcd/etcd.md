# etcd介绍与使用

## 介绍

**概念**：高可用的分布式key-value存储，可以用于配置共享和服务发现

**类似项目**：zookeeper 和 consul

**开发语言**：Go

**接口**：提供RESTful的HTTP接口，使用简单

**实现算法**：基于raft算法的强一致性，高可用的服务存储目录

**应用场景**：

- 服务发现和服务注册
- 配置中心
- 分布式锁
- master 选举



## etcd搭建

第一步：到官网 [release](https://github.com/etcd-io/etcd/releases) 处下载，如果 release 下载缓慢，[参考此处](./github.md#Release 下载缓慢解决方案)。

例如 Windows 用户直接下载其中的  [etcd-v3.4.9-windows-amd64.zip](https://github.com/etcd-io/etcd/releases/download/v3.4.9/etcd-v3.4.9-windows-amd64.zip)

第二步：解压



## etcd使用

第一步：启动 etcd

双击 `etcd.exe` 即可启动

第二步：使用 etcdctl 进行配置



第三步：下载第三方库

