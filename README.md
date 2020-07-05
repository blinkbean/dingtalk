# dingtalk
钉钉机器人消息封装——Golang

目前自定义机器人支持
- 文本（text）
- 链接（link）
- markdown
- ActionCard
    - 整体跳转
    - 独立跳转
- FeedCard

[机器人官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## 使用
### 获取
    ```shell script
    go get github.com/liyuxinger/dingtalk
    ```
### 初始化
    ```go
    import "github.com/liyuxinger/dingtalk"
    
    func main() {
        // 单个机器人有单位时间内消息条数的限制，如果有需要可以初始化多个token，发消息时随机发给其中一个机器人。
        var dingToken = []string{"7bd675b66646ba890046c2198257576470099e1bda0770bad7dd6684fb1e0415"}
        cli := dingtalk.InitDingTalk(dingToken, ".")
        cli.SendTextMessage("content")
    }
    ```
### text类型
- 方法及可选参数
    ```go
    // 方法定义
    SendTextMessage(content string, opt ...atOption) error

    // 可选参数
    // @所有人
    WithAtAll()
    
    // @指定群成员
    WithAtMobiles(mobiles []string)
    ```
- 使用
    ```go
    // at所有人
    cli.SendTextMessage("content", WithAtAll())

    // at指定群成员
    mobiles := []string{"131********"}
    cli.SendTextMessage("content", WithAtMobiles(mobiles))
    ```
- ![text](./img/Xnip2020-07-05_10-46-59.jpg)

### link类型
- 方法
    ```go
    // 方法定义
    SendLinkMessage(title, text, picUrl, msgUrl string) error
    ```
- 使用
    ```go
    cli.SendLinkMessage(title, text, picUrl, msgUrl)
    ```
- ![text](./img/Xnip2020-07-05_10-25-33.jpg)

### markdown类型
- 方法及可选参数
    ```go
    // 方法定义
    // text：markdown格式字符串
    SendMarkDownMessage(title, text string, opts ...atOption) error
    
    // 可选参数 目前钉钉markdown格式消息不支持@（可能是钉钉的bug），所以以下可选参数暂时不生效。
    // @所有人
    WithAtAll()
    
    // @指定群成员
    WithAtMobiles(mobiles []string)
    ```
- 使用
    ```go
    cli.SendMarkDownMessage(title, text)
    ```
- ![text](./img/Xnip2020-07-05_10-27-33.jpg)

### 整体跳转ActionCard类型
- 方法及可选参数
    ```go
    // 方法定义
    SendActionCardMessage(title, text string, opts ...actionCardOption) error
    
    // 可选参数
    // 标题
    WithCardSingleTitle(title string)
    
    // 跳转地址
    WithCardSingleURL(url string)
    ```
- 使用
    ```go
    cli.SendActionSingleMessage(title, text, WithCardSingleTitle(sTitle), WithCardSingleURL(url))
    ```
- ![text](./img/Xnip2020-07-05_10-28-57.jpg)

### 独立跳转ActionCard类型
- 方法及可选参数
    ```go
    // 方法定义
    SendActionCardMessage(title, text string, opts ...actionCardOption) error
    
    // 可选参数
    // 按钮排列方向，默认水平
    WithCardBtnVertical()
  
    // 跳转按钮
    WithCardBtns(btns []ActionCardMultiBtnModel)

    // ActionCardMultiBtnModel
    type ActionCardMultiBtnModel struct {
    	Title     string `json:"title,omitempty"`
    	ActionURL string `json:"actionURL,omitempty"`
    }
    ```
- 使用
    ```go
    btns := []ActionCardMultiBtnModel{{
        Title:     "test1",
        ActionURL: testUrl,
        },{
        Title:     "test2",
        ActionURL: testUrl,
        },
    }
    cli.SendActionSingleMessage(title, text, WithCardBtns(btns))
    ```
- ![text](./img/Xnip2020-07-05_10-29-21.jpg)

### FeedCard类型
- 方法
    ```go
    // 方法定义
    SendFeedCardMessage(feedCard []FeedCardLinkModel) error
    
    // FeedCardLinkModel
    type FeedCardLinkModel struct {
    	Title      string `json:"title,omitempty"`
    	MessageURL string `json:"messageURL,omitempty"`
    	PicURL     string `json:"picURL,omitempty"`
    }
    ```
- 使用
    ```go
    links := []FeedCardLinkModel{
        {
            Title:      "FeedCard1.",
            MessageURL: testUrl,
            PicURL:     testImg,
        },
        {
            Title:      "FeedCard2",
            MessageURL: testUrl,
            PicURL:     testImg,
        },
        {
            Title:      "FeedCard3",
            MessageURL: testUrl,
            PicURL:     testImg,
        },
    }
    cli.SendFeedCardMessage(links)
    ```
- ![text](./img/Xnip2020-07-05_10-30-02.jpg)