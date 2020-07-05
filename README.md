# dingtalk
钉钉机器人消息封装——Golang

目前自定义机器人支持文本（text）、链接（link）、Markdown三种消息格式，五种消息类型，[机器人官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)。

## 使用
1. 获取
```shell script
go get github.com/hugozhu/godingtalk
```
```go
import (
    "github.com/hugozhu/godingtalk"
)
```
### text类型

![text](./img/Xnip2020-07-05_10-46-59.jpg)

### link类型

![text](./img/Xnip2020-07-05_10-25-33.jpg)

### markdown类型

![text](./img/Xnip2020-07-05_10-27-33.jpg)

### 整体跳转ActionCard类型

![text](./img/Xnip2020-07-05_10-28-57.jpg)

### 独立跳转ActionCard类型

![text](./img/Xnip2020-07-05_10-29-21.jpg)

#### FeedCard类型

![text](./img/Xnip2020-07-05_10-30-02.jpg)