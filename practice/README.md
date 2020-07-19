# 编程实战





## 适合上手的项目

> 引用自：[HelloGitHub](https://github.com/521xueweihan/HelloGitHub)

1. [7days-golang](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/geektutu/7days-golang)：用 Go 在 7 天时间内实现 Web 框架、分布式缓存等应用的实战教程

2. [statping](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/statping/statping)：一个 Go 编写的服务状态展示页项目。通过该项目可以快速搭建起一个展示服务可用状态、服务质量的页面

    ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/46/img/statping.gif)

3. [gormt](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/xxjwxc/gormt)：一款 MySQL 数据库转 Go struct 的工具。支持：

   - 命令行、界面方式生成
   - YML 文件灵活配置
   - 自动生成快捷操作函数
   - 支持索引、外键等

    ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/46/img/gormt.gif) 

4. [gojsonq](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/thedevsaddam/gojsonq)：一款支持解析、查询 JSON/YAML/XML/CSV 数据的 Go 三方开源库。示例代码：

   ```go
   package main
   
   import "github.com/thedevsaddam/gojsonq"
   
   func main() {
   	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
   	name := gojsonq.New().FromString(json).Find("name.first")
   	println(name.(string)) // Tom
   }
   ```

   

5. [gods](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/emirpasic/gods)：简单易用的 Go 语言各种数据结构和算法，并封装成了一个库，开箱即食。示例代码：

   ```go
   package main
   
   import (
   	"github.com/emirpasic/gods/lists/arraylist"
   	"github.com/emirpasic/gods/utils"
   )
   
   func main() {
   	list := arraylist.New()
   	list.Add("a")                         // ["a"]
   	list.Add("c", "b")                    // ["a","c","b"]
   	list.Sort(utils.StringComparator)     // ["a","b","c"]
   	_, _ = list.Get(0)                    // "a",true
   	_, _ = list.Get(100)                  // nil,false
   	_ = list.Contains("a", "b", "c")      // true
   	_ = list.Contains("a", "b", "c", "d") // false
   	list.Swap(0, 1)                       // ["b","a",c"]
   	list.Remove(2)                        // ["b","a"]
   	list.Remove(1)                        // ["b"]
   	list.Remove(0)                        // []
   	list.Remove(0)                        // [] (ignored)
   	_ = list.Empty()                      // true
   	_ = list.Size()                       // 0
   	list.Add("a")                         // ["a"]
   	list.Clear()                          // []
   	list.Insert(0, "b")                   // ["b"]
   	list.Insert(0, "a")                   // ["a","b"]
   }
   ```

   

6. [gowp](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/xxjwxc/gowp)：Go 高性能异步并发线程池。接口调用简单、支持错误返回、无论排队多少任务，都不会阻止提交任务。可用于控制并发访问、并发执行。示例代码：

   ```go
   package main
   
   import (
   	"fmt"
   	"time"
   
   	"github.com/xxjwxc/gowp/workpool"
   )
   
   func main() {
   	wp := workpool.New(10)             //设置最大线程数
   	for i := 0; i < 20; i++ { //开启20个请求
   		ii := i
   		wp.Do(func() error {
   			for j := 0; j < 10; j++ { //每次打印0-10的值
   				time.Sleep(1 * time.Second)
   			}
   			return nil
   		})
   	}
   
   	wp.Wait()
   	fmt.Println("down")
   }
   ```

   

7. [evans](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/ktr0731/evans)：基于 Go 语言实现的支持交互模式的 gRPC 客户端，让调试、测试 gRPC API 更加容易

    ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/44/img/evans.png) 

8. [gochat](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/LockGit/gochat)：纯 Go 实现的轻量级即时通讯系统。技术上各层之间通过 rpc 通讯，使用 redis 作为消息存储与投递的载体，相对 kafka 操作起来更加方便快捷。各层之间基于 etcd 服务发现，在扩容部署时将会方便很多。架构、目录结构清晰，文档详细。而且还提供了 docker 一件构建，安装运行十分方便，推荐作为学习项目

    ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/44/img/gochat.gif) 

9. [go-admin](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/GoAdminGroup/go-admin)：基于 Golang 语言的数据可视化与管理平台。特性如下：

   - 🚀高生产效率：10 分钟内做一个好看的管理后台
   - 🎨主题：默认为 adminlte，更多好看的主题正在制作中，欢迎给我们留言
   - 🔢插件化：提供插件使用，真正实现一个插件解决不了问题，那就两个
   - ✅认证：开箱即用的 rbac 认证系统
   - ⚙️框架支持：支持大部分框架接入，让你更容易去上手和扩展

    ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/43/img/go-admin.png) 

10. [zerolog](https://hellogithub.com/periodical/statistics/click/?target=https://github.com/rs/zerolog)：一个速度快、专门用于输出 JSON 格式日志的库。还在为解析不规则的日志而烦恼吗？有了 zerolog 你可以跳起来了！当然它还有低效但可在控制台输出漂亮日志的模式，快去试试吧。示例代码

    ```go
    package main
    
    import (
        "github.com/rs/zerolog"
        "github.com/rs/zerolog/log"
    )
    
    func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    
        log.Info().Msg("hello world")
    }
    
    // Output: {"time":1516134303,"level":"info","message":"hello world"}
    ```

     ![img](https://raw.githubusercontent.com/521xueweihan/img/master/hellogithub/43/img/zerolog.png) 

11. 

 