

# Kafka

## 简介

**官方描述**：

>  *Kafka is used for building* **real-time data** *pipelines and streaming apps. It is* **horizontally scalable**,**fault-tolerant**,**wicked fast**, and runs in production in thousands of companies.

大致的意思就是，这是一个实时数据处理系统，可以横向扩展、高可靠，而且还变态快，已经被很多公司使用。

**特性**：

1. 可以让你发布和订阅流式的记录。这一方面与消息队列或者企业消息系统类似。
2. 可以储存流式的记录，并且有较好的容错性。
3. 可以在流式记录产生时就进行处理。

**适用场景**：

1. 构造实时流数据管道，它可以在系统或应用之间可靠地获取数据。 (相当于message queue)
2. 构建实时流式应用程序，对这些流数据进行转换或者影响。 (就是流处理，通过kafka stream topic和topic之间内部进行变化)

3. 异步处理，把非关键流程异步化，提高系统的响应时间和健壮性
4. 应用解耦，通过消息队列
5. 流量削峰

## 文档

[kafka中文文档](http://kafka.apachecn.org/)



## 快速开始

以 **Windows** 为例：

### 第一步：下载并解压缩 kafka

### 第二步：启动服务器

1. 启动 Zookeeper 服务器：

Zookeeper 的配置文件在 conf 目录下，这个目录下有 zoo_sample.cfg 和 log4j.properties，你需要做的就是将 zoo_sample.cfg 改名为 zoo.cfg，因为 Zookeeper 在启动时会找这个文件作为默认配置文件 

```bash
$ cp ./conf/zoo_sample.cfg zoo.cfg
$ ./bin/zkServer.cmd
```

2. 启动 Kafka 服务器：

```bash
$ ./bin/windows/kafka-server-start.bat config/server.properties
```

如果出现了以下错误：

> 错误: 找不到或无法加载主类 Files\Java\jdk1.8.0_131\lib;C:\Program

编辑 `/bin/windows/kafka-run-class.bat` 修改大约在179行(notepad打开的话在179行),就是在后面%CLASSPATH%"加上英文双引号"%CLASSPATH%" 

```bat
set COMMAND=%JAVA% %KAFKA_HEAP_OPTS% %KAFKA_JVM_PERFORMANCE_OPTS% %KAFKA_JMX_OPTS% %KAFKA_LOG4J_OPTS% -cp "%CLASSPATH%" %KAFKA_OPTS% %*
```

### 第三步：使用 kafka

**创建 topic**

让我们创建一个名为“test”的topic，它有一个分区和一个副本： 

```bash
$ ./bin/windows/kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
```

**发送消息到 topic**

运行 producer，然后在控制台输入一些消息以发送到服务器。 

```bash
$ ./bin/windows/kafka-console-producer.bat --broker-list localhost:9092 --topic test
> This is a message
> This is another message
```

**启动一个 consumer**

```bash
$ ./bin/windows/kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic test --from-beginning
This is a message
This is another message
```

### 第四步：下载第三方库

```
$ go get github.com/Shopify/sarama
```

第六步：通过第三方库使用 kafka

[Code](./main.go)

