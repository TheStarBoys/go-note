# 日志收集系统架构

**项目背景**

- 每个系统都有日志，当系统出现问题时，需要通过日志解决问题
- 当系统机器比较少时，登陆到服务器上查看即可满足
- 当系统及其规模巨大，登陆到机器上查看几乎不现实

**解决方案**

- 把机器上的日志实时手机，同意的存储到中心系统
- 然后再对这些日志建立索引，通过搜索既可以找到对应日志
- 通过提供界面友好的web界面，通过web既可以完成日志搜索
