# RPC

> 详情推荐书籍《Go高级编程》中 [RPC](https://github.com/chai2010/advanced-go-programming-book/tree/master/ch4-rpc) 章节

## 简介

概念：RPC是远程过程调用（Remote Procedure Call）的缩写，通俗地说就是调用远处的一个函数，这个“远处”可以是同一文件内的不同函数，也可以是同一机器的另一个进程中的函数，也可能是远在火星的火箭中的某个秘密方法。

由于 rpc 涉及的函数可能非常远，远到它们之间说着不同的语言，所以语言就成了两边沟通的障碍。而 ProtoBuf 由于支持多种编程语言（甚至不支持的语言也可以扩展），其本身特性也非常方便描述服务的接口（也就是方法列表），因此非常适合作为 rpc 世界的接口交流语言。



## gRPC

gRPC是谷歌公司基于 ProtoBuf 开发的跨语言的开源 RPC 框架。gRPC 基于 HTTP/2 协议设计，可以基于一个 HTTP/2 链接提供多个服务，对移动设备更友好。

