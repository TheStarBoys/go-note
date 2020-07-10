# ProtoBuf

## 简介

官方描述：

> protocol buffers 是一种语言无关、平台无关、可扩展的序列化结构数据的方法，它可用于（数据）通信协议、数据存储等。
>
> Protocol Buffers 是一种灵活，高效，自动化机制的结构数据序列化方法－可类比 XML，但是比 XML 更小（3 ~ 10倍）、更快（20 ~ 100倍）、更为简单。
>
> 你可以定义数据的结构，然后使用特殊生成的源代码轻松的在各种数据流中使用各种语言进行编写和读取结构数据。你甚至可以更新数据结构，而不破坏由旧数据结构编译的已部署程序。

 简单来讲， ProtoBuf 是结构数据**序列化** 方法，可简单**类比于 XML**，其具有以下**特点**： 

- **语言无关、平台无关**。即 ProtoBuf 支持 Java、C++、Python 等多种语言，支持多个平台
- **高效**。即比 XML 更小（3 ~ 10倍）、更快（20 ~ 100倍）、更为简单
- **扩展性、兼容性好**。你可以更新数据结构，而不影响和破坏原有的旧程序

**缺点**：

- 编码后的数据可读性差



## 如何使用

### 基本步骤

1. 创建 `.proto` 文件定义数据结构
2. `protoc` 编译 `.proto` 文件生成读写接口
3. 调用接口实现序列化、反序列化以及读写



### 具体操作

**第一步：创建 `.proto` 文件定义数据结构，如下例1所示：**

```protobuf
syntax = "proto3"; // 表示采用 proto3 的语法
// 可以导入其他 .proto 文件中的数据结构，导入语法如下所示
// import "xxx.proto"; 可以是相对路径，也可以是绝对路径

// 例1: 在 xxx.proto 文件中定义 Example1 message
// 指定 proto文件 所处包
package main;

message Example1 {
    string stringVal = 1;
    bytes bytesVal = 2;
    message EmbeddedMessage {
        int32 int32Val = 1;
        string stringVal = 2;
    }
    EmbeddedMessage embeddedExample1 = 3;
    repeated int32 repeatedInt32Val = 4;
    repeated string repeatedStringVal = 5;
}
```

我们在上例中定义了一个名为 Example1 的 消息，语法很简单，message 关键字后跟上消息名称：

```protobuf
message xxx {

}
```

之后我们在其中定义了 message 具有的字段，形式为： 

```protobuf
message xxx {
  // 字段规则：required -> 字段只能也必须出现 1 次
  // 字段规则：repeated -> 字段可出现任意多次（包括 0）
  // 类型：int32、int64、sint32、sint64、string、32-bit ....
  // 字段编号：0 ~ 536870911（除去 19000 到 19999 之间的数字）
  字段规则 类型 名称 = 字段编号;
}
```

在上例中，我们定义了：

- 类型 string，名为 stringVal 的 optional 可选字段，字段编号为 1，此字段可出现 0 或 1 次
- 类型 bytes，名为 bytesVal 的 optional 可选字段，字段编号为 2，此字段可出现 0 或 1 次
- 类型 EmbeddedMessage（自定义的内嵌 message 类型），名为 embeddedExample1 的 optional 可选字段，字段编号为 3，此字段可出现 0 或 1 次
- 类型 int32，名为 repeatedInt32Val 的 repeated 可重复字段，字段编号为 4，此字段可出现 任意多次（包括 0）
- 类型 string，名为 repeatedStringVal 的 repeated 可重复字段，字段编号为 5，此字段可出现 任意多次（包括 0）

关于 proto2 定义 message 消息的更多语法细节，例如具有支持哪些类型，字段编号分配、import
 导入定义，reserved 保留字段等知识请参阅 [[翻译\] ProtoBuf 官方文档（二）- 语法指引（proto2）](https://www.jianshu.com/p/6f68fb2c7d19)。

关于定义时的一些规范请参阅 [[翻译\] ProtoBuf 官方文档（四）- 规范指引](https://www.jianshu.com/p/8c55fb0a09b5)



**第二步：`protoc` 编译 `.proto` 文件生成读写接口**

我们在 `.proto` 文件中定义了数据结构，这些数据结构是面向开发者和业务程序的，并不面向存储和传输。

当需要把这些数据进行存储或传输时，就需要将这些结构数据进行序列化、反序列化以及读写。那么如何实现呢？不用担心， ProtoBuf 将会为我们提供相应的接口代码。如何提供？答案就是通过 `protoc` 这个编译器。

可通过如下命令生成相应的接口代码（目标语言以Go为例），详情看后面的编译部分：

```
$ protoc --go_out=. example1.proto
```

这里只生成了一个 example1.pb.go 文件，其主要内容如下：

```go
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 例1: 在 xxx.proto 文件中定义 Example1 message
type Example1 struct {
	StringVal            string
	BytesVal             []byte
	EmbeddedExample1     *Example1_EmbeddedMessage
	RepeatedInt32Val     []int32
	RepeatedStringVal    []string
}
func (m *Example1) Reset()         { *m = Example1{} }
func (m *Example1) String() string { return proto.CompactTextString(m) }
func (*Example1) ProtoMessage()    {}
func (m *Example1) GetStringVal() string {
	if m != nil {
		return m.StringVal
	}
	return ""
}

func (m *Example1) GetBytesVal() []byte {
	if m != nil {
		return m.BytesVal
	}
	return nil
}

func (m *Example1) GetEmbeddedExample1() *Example1_EmbeddedMessage {
	if m != nil {
		return m.EmbeddedExample1
	}
	return nil
}

func (m *Example1) GetRepeatedInt32Val() []int32 {
	if m != nil {
		return m.RepeatedInt32Val
	}
	return nil
}

func (m *Example1) GetRepeatedStringVal() []string {
	if m != nil {
		return m.RepeatedStringVal
	}
	return nil
}

type Example1_EmbeddedMessage struct {
	Int32Val             int32
	StringVal            string
}
func (m *Example1_EmbeddedMessage) Reset()         { *m = Example1_EmbeddedMessage{} }
func (m *Example1_EmbeddedMessage) String() string { return proto.CompactTextString(m) }
func (*Example1_EmbeddedMessage) ProtoMessage()    {}
func (m *Example1_EmbeddedMessage) GetInt32Val() int32 {
	if m != nil {
		return m.Int32Val
	}
	return 0
}

func (m *Example1_EmbeddedMessage) GetStringVal() string {
	if m != nil {
		return m.StringVal
	}
	return ""
}
```

可以看到编译后得到的文件中，生成了对应的结构体类型 `Example1` ，并且自动生成了一组方法。

其中 `ProtoMessage()` 方法表示这是一个实现了 `proto.Message` 接口的方法。

此外，ProtoBuf 还为结构体中的每个成员生成了一个 `Get` 方法，不仅可以用于处理空指针类型，而且可以和ProtoBuf 第2版的方法保持一致。

**第三步：调用接口实现序列化、反序列化以及读写**

```go
import (
	"testing"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func TestExample1(t *testing.T) {
	pb := Example1{
		StringVal: "1",
		BytesVal: []byte{},
		EmbeddedExample1: &Example1_EmbeddedMessage{
			Int32Val: 2,
			StringVal: "3",
		},
		RepeatedInt32Val: []int32{1, 2, 3},
		RepeatedStringVal: []string{"1", "2", "3"},
	}
	buf, err := proto.Marshal(&pb) // 序列化
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", buf)
	// 0a01311a05080212013322030102032a01312a01322a0133

	pb1 := Example1{}
	proto.Unmarshal(buf, &pb1) // 反序列化
	fmt.Printf("%#v\n", pb1)
	// main.Example1{
	//     StringVal:"1",
	//     BytesVal:[]uint8(nil),
	//     EmbeddedExample1:(*main.Example1_EmbeddedMessage)(0xc000106380),
	//     RepeatedInt32Val:[]int32{1, 2, 3},
	//     RepeatedStringVal:[]string{"1", "2", "3"}
	// }

	// 访问字段的方式获取数据
	fmt.Println("pb.EmbeddedExample1.Int32Val:", pb.EmbeddedExample1.Int32Val)
	// 使用 Get 方法获取数据
	fmt.Println("pb.EmbeddedExample1.Int32Val:", pb.GetEmbeddedExample1().GetInt32Val())
	// pb.EmbeddedExample1.Int32Val: 2
}
```



## 编译ProtoBuf

第一步：安装官方的 `protoc` 工具，可以从其 GitHub 官网的 [release](https://github.com/protocolbuffers/protobuf/releases) 下载：

以 Windows 操作系统为例，找到 `protoc-3.12.3-win64.zip` 链接，点击即可下载，解压后会得到一个可执行文件 `protoc.exe` 直接放到你的环境变量里面去，推荐直接加到 `$GOPATH/bin` 下，把这个路径加入到环境变量。

**注意**：如果 release 下载缓慢，[参考此处](./github.md#Release 下载缓慢解决方案)。

第二步：在官方的 `protoc` 编译期中并不支持 Go 语言，因此想基于 `.proto` 文件生成相应的 Go 代码，需要安装相应的插件。

二者任选其一：

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
```

或

```
$ go install github.com/golang/protobuf/protoc-gen-go
```

第三步：生成 Go 代码

命令介绍：

```
// $SRC_DIR: .proto 所在的源目录
// --go_out: 生成 Go 代码
// $DST_DIR: 生成代码的目标目录
// xxx.proto: 要针对哪个 proto 文件生成接口代码

$ protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/xxx.proto
```

示例：

```
$ protoc --go_out=. hello.proto
```

最终生成 hello.pb.go 文件

## ProtoBuf详细语法



## 代码生成插件

