# Go 使用过程中的一些问题解决方案

## 包管理

### Module

**代理**

https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md

```bash
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.cn
```

**Goland 使用 Go Module**

编辑 编译运行时的 Configuration，添加一个Environment：

GO111MODULE=on

