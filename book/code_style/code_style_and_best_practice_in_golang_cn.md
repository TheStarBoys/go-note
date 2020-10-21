# Go 代码风格与最佳实践

## 目录

- [代码风格](#代码风格)
  - [核心观点](#核心观点)
  - [风格](#风格)
  - [工具](#工具)
- [最佳实践](#最佳实践)
  - [指导原则](#指导原则)
  - [编程陷阱](#编程陷阱)
  - [测试](#测试)
  - [设计模式](#设计模式)
- [参考资料](#参考资料)



## 代码风格

整洁的代码简单直接。整洁的代码如同优美的散文。整洁的代码从不隐藏设计者的意图，充满了干净利落的抽象和直截了当的控制语句。



### 核心观点

- 良好风格的代码阅读起来就像看书一样，望文生义，由浅入深
- 团队内部保持相同的代码风格，以至于看不出来具体是谁写的哪部分代码
- 你不是在编程，你是在进行艺术创作



### 风格

#### 包

##### import 分组

导入应该分为两组：

- 标准库
- 其他库

默认情况下，这是 goimports 应用的分组。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
import (
  "fmt"
  "os"
  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```

</td><td>

```go
import (
  "fmt"
  "os"

  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```

</td></tr>
</tbody></table>



##### 相似声明放一起

Go 语言支持将相似的声明放在一个组内。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
import "a"
import "b"
```

</td><td>

```go
import (
  "a"
  "b"
)
```

</td></tr>
</tbody></table>

这同样适用于常量、变量和类型声明：

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
const a = 1
const b = 2

var a = 1
var b = 2

type Area float64
type Volume float64
```

</td><td>

```go
const (
  a = 1
  b = 2
)

var (
  a = 1
  b = 2
)

type (
  Area float64
  Volume float64
)
```

</td></tr>
</tbody></table>

仅将相关的声明放在一组。不要将不相关的声明放在一组。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
  ENV_VAR = "MY_ENV"
)
```

</td><td>

```go
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
)

const ENV_VAR = "MY_ENV"
```

</td></tr>
</tbody></table>

分组使用的位置没有限制，例如：你可以在函数内部使用它们：

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
func f() string {
  var red = color.New(0xff0000)
  var green = color.New(0x00ff00)
  var blue = color.New(0x0000ff)

  ...
}
```

</td><td>

```go
func f() string {
  var (
    red   = color.New(0xff0000)
    green = color.New(0x00ff00)
    blue  = color.New(0x0000ff)
  )

  ...
}
```

</td></tr>
</tbody></table>



##### 包名

当命名包时，请按下面规则选择一个名称：

- 全部小写。没有大写或下划线。
- 大多数使用命名导入的情况下，不需要重命名。
- 简短而简洁。请记住，在每个使用的地方都完整标识了该名称。
- 不用复数。例如`net/url`，而不是`net/urls`。
- 不要用“common”，“util”，“shared”或“lib”。这些是不好的，信息量不足的名称。



#### 命名



##### 避免参数、变量语义不明确

函数调用中的`意义不明确的参数`可能会损害可读性。当参数名称的含义不明显时，请为参数添加 C 样式注释 (`/* ... */`)

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true, true)
```

</td><td>

```go
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true /* isLocal */, true /* done */)
```

</td></tr>
</tbody></table>

对于上面的示例代码，还有一种更好的处理方式是将上面的 `bool` 类型换成自定义类型。将来，该参数可以支持不仅仅局限于两个状态（true/false）。

```go
type Region int

const (
  UnknownRegion Region = iota
  Local
)

type Status int

const (
  StatusReady Status= iota + 1
  StatusDone
  // Maybe we will have a StatusInProgress in the future.
)

func printInfo(name string, region Region, status Status)
```



#### 函数

检查函数是否满足以下条件：

- 短小
- 只做一件事
- 每个函数一个抽象层级
- 使用描述性的函数名
- 别重复自己（在编程时，如果写了过多重复代码，那么考虑封装成函数）



可导出的函数加以详细的注释，阐述清楚函数的意图。

注意阅读 `GetAndSaveDataFromURL()` 函数时的顺序，是很自然的从上往下阅读，由浅入深的层级关系。

每个函数只做了一件事，并且函数足够简短，非导出函数可以从函数名就能清晰的知道其意图。

```go
// GetAndSaveDataFromURL save data into local disk
// after get it from url.
func GetAndSaveDataFromURL(url string) {
	data := getSpecialContentFromURL(url)
	saveDataToDisk(data)
}

func getSpecialContentFromURL(url string) string {
	resp := getRespFromURL(url)

	return getSpecialContentFromResp(resp)
}

func saveDataToDisk(data string) {
	// save data
}

func getRespFromURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle err
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
		return ""
	}

	return string(body)
}

func getSpecialContentFromResp(resp string) string {
	return ""
}
```



#### 注释

**对于注释的态度**：能通过代码阐述清楚的话，尽量避免使用注释

**坏注释**

- 多余的注释（看注释不比看代码省事）

```go
// waitForClose wait some time if resource hasn't been closed.
// The program will panic if resource hasn't been closed after timeout.
func waitForClose(timeout time.Duration) {
	if !isClose() {
		time.Sleep(timeout)
	}

	if !isClose() {
		panic("some resource could not be closed.")
	}
}

var closed bool

func isClose() bool {
	return closed
}
```



- 废话注释

```go
// dayOfMonth The day of the month
var dayOfMonth int

// getDayOfMonth return the day of the month
func getDayOfMonth() int {
	return dayOfMonth
}
```



### 工具

代码应该使用一下工具进行审查

- goimports
- golint
- go vet



## 最佳实践

### 指导原则

#### 处理错误

##### 错误类型

Go 中有多种声明错误（Error) 的选项：

- [`errors.New`] 对于简单静态字符串的错误
- [`fmt.Errorf`] 用于格式化的错误字符串
- 实现 `Error()` 方法的自定义类型
- 用 [`"pkg/errors".Wrap`] 的 Wrapped errors

返回错误时，请考虑以下因素以确定最佳选择：

- 这是一个不需要额外信息的简单错误吗？如果是这样，[`errors.New`] 足够了。

- 客户需要检测并处理此错误吗？如果是这样，则应使用自定义类型并实现该 `Error()` 方法。

- 您是否正在传播下游函数返回的错误？如果是这样，请查看本文后面有关错误包装 [section on error wrapping](#错误包装 (Error-Wrapping)) 部分的内容。

- 否则 [`fmt.Errorf`] 就可以了。

  [`errors.New`]: https://golang.org/pkg/errors/#New
  [`fmt.Errorf`]: https://golang.org/pkg/fmt/#Errorf
  [`"pkg/errors".Wrap`]: https://godoc.org/github.com/pkg/errors#Wrap

如果客户端需要检测错误，并且您已使用创建了一个简单的错误 [`errors.New`]，请使用一个错误变量。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// package foo

func Open() error {
  return errors.New("could not open")
}

// package bar

func use() {
  if err := foo.Open(); err != nil {
    if err.Error() == "could not open" {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td><td>

```go
// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
  return ErrCouldNotOpen
}

// package bar

if err := foo.Open(); err != nil {
  if err == foo.ErrCouldNotOpen {
    // handle
  } else {
    panic("unknown error")
  }
}
```

</td></tr>
</tbody></table>

如果您有可能需要客户端检测的错误，并且想向其中添加更多信息（例如，它不是静态字符串），则应使用自定义类型。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
func open(file string) error {
  return fmt.Errorf("file %q not found", file)
}

func use() {
  if err := open("testfile.txt"); err != nil {
    if strings.Contains(err.Error(), "not found") {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td><td>

```go
type errNotFound struct {
  file string
}

func (e errNotFound) Error() string {
  return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
  return errNotFound{file: file}
}

func use() {
  if err := open("testfile.txt"); err != nil {
    if _, ok := err.(errNotFound); ok {
      // handle
    } else {
      panic("unknown error")
    }
  }
}
```

</td></tr>
</tbody></table>

直接导出自定义错误类型时要小心，因为它们已成为程序包公共 API 的一部分。最好公开匹配器功能以检查错误。

```go
// package foo

type errNotFound struct {
  file string
}

func (e errNotFound) Error() string {
  return fmt.Sprintf("file %q not found", e.file)
}

func IsNotFoundError(err error) bool {
  _, ok := err.(errNotFound)
  return ok
}

func Open(file string) error {
  return errNotFound{file: file}
}

// package bar

if err := foo.Open("foo"); err != nil {
  if foo.IsNotFoundError(err) {
    // handle
  } else {
    panic("unknown error")
  }
}
```

<!-- TODO: Exposing the information to callers with accessor functions. -->



##### 错误包装 (Error Wrapping)

一个（函数/方法）调用失败时，有三种主要的错误传播方式：

- 如果没有要添加的其他上下文，并且您想要维护原始错误类型，则返回原始错误。
- 添加上下文，使用 [`"pkg/errors".Wrap`] 以便错误消息提供更多上下文 ,[`"pkg/errors".Cause`] 可用于提取原始错误。
- 如果调用者不需要检测或处理的特定错误情况，使用 [`fmt.Errorf`]。

建议在可能的地方添加上下文，以使您获得诸如“调用服务 foo：连接被拒绝”之类的更有用的错误，而不是诸如“连接被拒绝”之类的模糊错误。

在将上下文添加到返回的错误时，请避免使用“failed to”之类的短语以保持上下文简洁，这些短语会陈述明显的内容，并随着错误在堆栈中的渗透而逐渐堆积：

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "failed to create new store: %s", err)
}
```

</td><td>

```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "new store: %s", err)
}
```

<tr><td>

```
failed to x: failed to y: failed to create new store: the error
```

</td><td>

```
x: y: new store: the error
```

</td></tr>
</tbody></table>

但是，一旦将错误发送到另一个系统，就应该明确消息是错误消息（例如使用`err`标记，或在日志中以”Failed”为前缀）。

另请参见 [Don't just check errors, handle them gracefully]. 不要只是检查错误，要优雅地处理错误

[`"pkg/errors".Cause`]: https://godoc.org/github.com/pkg/errors#Cause
[Don't just check errors, handle them gracefully]: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully



##### 不要 panic

在生产环境中运行的代码必须避免出现 panic。panic 是 [cascading failures] 级联失败的主要根源 。如果发生错误，该函数必须返回错误，并允许调用方决定如何处理它。

[cascading failures]: https://en.wikipedia.org/wiki/Cascading_failure

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
func run(args []string) {
  if len(args) == 0 {
    panic("an argument is required")
  }
  // ...
}

func main() {
  run(os.Args[1:])
}
```

</td><td>

```go
func run(args []string) error {
  if len(args) == 0 {
    return errors.New("an argument is required")
  }
  // ...
  return nil
}

func main() {
  if err := run(os.Args[1:]); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
```

</td></tr>
</tbody></table>

panic/recover 不是错误处理策略。仅当发生不可恢复的事情（例如：nil 引用）时，程序才必须 panic。程序初始化是一个例外：程序启动时应使程序中止的不良情况可能会引起 panic。

```go
var _statusTemplate = template.Must(template.New("name").Parse("_statusHTML"))
```

即使在测试代码中，也优先使用`t.Fatal`或者`t.FailNow`而不是 panic 来确保失败被标记。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// func TestFoo(t *testing.T)

f, err := ioutil.TempFile("", "test")
if err != nil {
  panic("failed to set up test")
}
```

</td><td>

```go
// func TestFoo(t *testing.T)

f, err := ioutil.TempFile("", "test")
if err != nil {
  t.Fatal("failed to set up test")
}
```

</td></tr>
</tbody></table>



#### 全局变量

##### 避免可变全局变量

使用选择依赖注入方式避免改变全局变量。 
既适用于函数指针又适用于其他值类型

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// sign.go
var _timeNow = time.Now
func sign(msg string) string {
  now := _timeNow()
  return signWithTime(msg, now)
}
```

</td><td>

```go
// sign.go
type signer struct {
  now func() time.Time
}
func newSigner() *signer {
  return &signer{
    now: time.Now,
  }
}
func (s *signer) Sign(msg string) string {
  now := s.now()
  return signWithTime(msg, now)
}
```

</td></tr>
<tr><td>

```go
// sign_test.go
func TestSign(t *testing.T) {
  oldTimeNow := _timeNow
  _timeNow = func() time.Time {
    return someFixedTime
  }
  defer func() { _timeNow = oldTimeNow }()
  assert.Equal(t, want, sign(give))
}
```

</td><td>

```go
// sign_test.go
func TestSigner(t *testing.T) {
  s := newSigner()
  s.now = func() time.Time {
    return someFixedTime
  }
  assert.Equal(t, want, s.Sign(give))
}
```

</td></tr>
</tbody></table>



### 编程陷阱

#### 切片

##### 向 []interface{} append 一个 []interface{} 需要明确是否需要解包

```go
var (
	data       = []interface{}{1, 2, 3}
	waitAppend = []interface{}{4, 5, 6}
)
fmt.Println(append(data, waitAppend))
fmt.Println(append(data, waitAppend...))
// [1 2 3 [4 5 6]]
// [1 2 3 4 5 6]
```



#### 指针与引用

引用类型：

- 切片
- 空接口
- map
- channel



##### 切片引用传递和指针传递的区别

引用传递

```go
func main() {
	data := []int{1, 2, 3}
	fn(data, []int{11, 12, 13})
	fmt.Println(data)
    // 引用传递能通过函数改变切片内值
    // 但无法新增外部切片的元素
    // [2 3 4]
}

func fn(data, waitAppend []int) {
	for i := range data {
		data[i]++
	}
	data = append(data, waitAppend...)
}
```

指针传递

```go
func main() {
	data := []int{1, 2, 3}
	fn(&data, []int{11, 12, 13})
	fmt.Println(data)
    // 指针传递不仅可以在函数内部修改切片值后影响到外部切片
    // 而且可以新增外部切片元素
    // [2 3 4 11 12 13]
}

func fn(data *[]int, waitAppend []int) {
	for i := range *data {
		(*data)[i]++
	}
	*data = append(*data, waitAppend...)
}
```



#### 通道 channel

##### 避免向已关闭的 channel 里写

```go
ch := make(chan int, 1)
close(ch)

ch <- 1
// panic: send on closed channel
```



##### 已关闭的 channel 不会阻塞读

```go
ch := make(chan int, 1)
close(ch)

i, ok := <-ch
fmt.Println(i, ok)
// 0 false
```



##### 无缓冲 channel 若无消费者，将阻塞程序

```go
ch := make(chan int)
ch <- 1
// fatal error: all goroutines are asleep - deadlock!
```



#### 接口 interface

##### Interface 合理性验证

在编译时验证接口的符合性。这包括：

- 将实现特定接口的导出类型作为接口API 的一部分进行检查
- 实现同一接口的(导出和非导出)类型属于实现类型的集合
- 任何违反接口合理性检查的场景,都会终止编译,并通知给用户

补充:上面3条是编译器对接口的检查机制,
大体意思是错误使用接口会在编译期报错.
所以可以利用这个机制让部分问题在编译期暴露.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// 如果Handler没有实现http.Handler,会在运行时报错
type Handler struct {
  // ...
}
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
```

</td><td>

```go
type Handler struct {
  // ...
}
// 用于触发编译期的接口的合理性检查机制
// 如果Handler没有实现http.Handler,会在编译期报错
var _ http.Handler = (*Handler)(nil)
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```

</td></tr>
</tbody></table>

如果 `*Handler` 与 `http.Handler` 的接口不匹配,
那么语句 `var _ http.Handler = (*Handler)(nil)` 将无法编译通过.

赋值的右边应该是断言类型的零值。
对于指针类型（如 `*Handler`）、切片和映射，这是 `nil`；
对于结构类型，这是空结构。

```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}
var _ http.Handler = LogHandler{}
func (h LogHandler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```



##### 处理类型断言失败

[类型断言] 的单个返回值形式针对不正确的类型将产生 panic。因此，请始终使用“comma ok”的惯用法。

[类型断言]: https://golang.org/ref/spec#Type_assertions

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
t := i.(string)
```

</td><td>

```go
t, ok := i.(string)
if !ok {
  // 优雅地处理错误
}
```

</td></tr>
</tbody></table>



### 测试

#### 测试命令

- 运行项目下所有的测试
  - go test -v ./...
- 显示测试过程
  - go test -v
- 测试覆盖率
  - go test -cover



#### 示例测试

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```



#### 表格驱动测试

当测试逻辑是重复的时候，通过  [subtests] 使用 table 驱动的方式编写 case 代码看上去会更简洁。

[subtests]: https://blog.golang.org/subtests

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>


```go
// func TestSplitHostPort(t *testing.T)

host, port, err := net.SplitHostPort("192.0.2.0:8000")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("192.0.2.0:http")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "http", port)

host, port, err = net.SplitHostPort(":8000")
require.NoError(t, err)
assert.Equal(t, "", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("1:8")
require.NoError(t, err)
assert.Equal(t, "1", host)
assert.Equal(t, "8", port)
```

</td><td>

```go
// func TestSplitHostPort(t *testing.T)

tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  {
    give:     "192.0.2.0:8000",
    wantHost: "192.0.2.0",
    wantPort: "8000",
  },
  {
    give:     "192.0.2.0:http",
    wantHost: "192.0.2.0",
    wantPort: "http",
  },
  {
    give:     ":8000",
    wantHost: "",
    wantPort: "8000",
  },
  {
    give:     "1:8",
    wantHost: "1",
    wantPort: "8",
  },
}

for _, tt := range tests {
  t.Run(tt.give, func(t *testing.T) {
    host, port, err := net.SplitHostPort(tt.give)
    require.NoError(t, err)
    assert.Equal(t, tt.wantHost, host)
    assert.Equal(t, tt.wantPort, port)
  })
}
```

</td></tr>
</tbody></table>

很明显，使用 test table 的方式在代码逻辑扩展的时候，比如新增 test case，都会显得更加的清晰。



### 设计模式

是指在软件开发中，经过验证的，用于解决在**特定环境**下**重复出现**的**特定问题**的**解决方案**。



#### 简单工厂

在同样的行为有不同的具体实现时，考虑使用简单工厂来选择实现。



```go
// Saver represent the object implemented this interface
// has ability to save data into somewhere
type Saver interface {
	Save()
}

// SaverKind ...
type SaverKind int

const (
	// SaverKindDB ...
	SaverKindDB SaverKind = iota + 1
	// SaverKindLog ...
	SaverKindLog
)

// NewSaver ...
func NewSaver(kind SaverKind) (Saver, error) {
	switch kind {
	case SaverKindDB:
		return &DB{}, nil
	case SaverKindLog:
		return &Log{}, nil
	default:
		return nil, fmt.Errorf("unknown SaverKind %d", kind)
	}
}

// DB ...
type DB struct{}

// Save save into database
func (db *DB) Save() {}

// Log ...
type Log struct{}

// Save save into log file
func (l *Log) Save() {}
```



## 参考资料

- uber-go/guide

  - [Chinese](https://github.com/xxjwxc/uber_go_guide_cn)

  - [English](https://github.com/uber-go/guide)

- [dgryski](https://github.com/dgryski)/[go-perfbook](https://github.com/dgryski/go-perfbook)

- [Go 高级编程](https://github.com/chai2010/advanced-go-programming-book)

- [Golang 面试题搜集](https://github.com/lifei6671/interview-go)

  - [https://github.com/xmge/gonote/tree/master/go%E9%9D%A2%E8%AF%95%E9%A2%98](https://github.com/xmge/gonote/tree/master/go面试题)

  - https://github.com/lifei6671/interview-go

- 《代码整洁之道》

- 设计模式[https://github.com/TheStarBoys/go-design-pattern]

- Learn Go with Tests

  - [Chinese](https://studygolang.gitbook.io/learn-go-with-tests/)

  - [English](https://github.com/quii/learn-go-with-tests)


