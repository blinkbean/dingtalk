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
### 创建钉钉群机器人
1. 选择添加`自定义`机器人。
2. 安全设置
    共有关键词、加签、IP白名单三种设置，需要根据情况进行选择。
    ![Xnip2020-07-05_15-55-24.jpg](https://i.loli.net/2020/07/05/4XqHG2dOwo8StEu.jpg)
3. 选择`自定义关键词`，这里设置的关键词在初始化机器人的时候会用到。
### 获取
-   ```go
    go get github.com/blinkbean/dingtalk
    ```
### 初始化
-   ```go
    // key 创建钉钉机器人需要设置的关键词，默认为`.`
    func InitDingTalk(tokens []string, key string) *dingTalk
    ```
-   ```go
    import "github.com/blinkbean/dingtalk"
    
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
- ![Xnip2020-07-05_10-46-59.jpg](https://i.loli.net/2020/07/05/LXErbH1KiRGstQ7.jpg)

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
- ![Xnip2020-07-05_10-25-33.jpg](https://i.loli.net/2020/07/05/wDG1sMPlU7XZQfr.jpg)

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
- Markdown进阶
    ```go
    // 按行写入数组，增强markdown的可读性
    msg := []string{
        "### Link test",
        "---",
        "- <font color=#00ff00 size=6>红色文字</font>",
        "- content2",
    }
    cli.SendMarkDownMessageBySlice("Markdown title", msg, WithAtAll())
  
    // 字体及颜色
    dm := DingMap()
    dm.Set("颜色测试", H2)
    dm.Set("失败：$$ 同行不同色 $$", RED)  // 双$分隔
    dm.Set("---", N)
    dm.Set("金色", GOLD)
    dm.Set("成功", GREEN)
    dm.Set("警告", BLUE)
    dm.Set("普通文字", N)
    cli.SendMarkDownMessageBySlice("color test", dm.Slice())
    ```
- ![Xnip2020-07-05_10-27-33.jpg](https://i.loli.net/2020/07/05/7LScefCZIGnDjBV.jpg)
- ![Xnip2020-07-26_17-14-40.jpg](https://i.loli.net/2020/07/26/PADJ5uqmfQht2cr.jpg)

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
- ![Xnip2020-07-05_10-28-57.jpg](https://i.loli.net/2020/07/05/kKELHAlomndiO9I.jpg)

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
- ![Xnip2020-07-05_10-29-21.jpg](https://i.loli.net/2020/07/05/t9GywHFSQUWCVDT.jpg)
- ![Xnip2020-07-26_17-14-56.jpg](https://i.loli.net/2020/07/26/pEg7hotXZnsaJPV.jpg)

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
- ![Xnip2020-07-05_10-30-02.jpg](https://i.loli.net/2020/07/05/F5WDLqyJ4Yzfj6A.jpg)

### OutGoing
- 消息格式
    ```json
    {
        "atUsers":[
            {
                "dingtalkId":"$:LWCP_v1:$1h0bmSzcLCHncx0lCt3Bb/UVz7xv/8vhs"
            }],
        "chatbotUserId":"$:LWCP_v1:$1h0bmSzcLCHncx0lCt3Bb/UVz7x/8vhd",
        "conversationId":"cidkkCwvtlh1L0RmFuhmashid==",
        "conversationTitle":"项目群",
        "conversationType":"2",
        "createAt":1595232438958,
        "isAdmin":false,
        "isInAtList":true,
        "msgId":"msgm/bJkKjTupFM7ZoRF/eKRA==",
        "msgtype":"text",
        "sceneGroupCode":"project",
        "senderId":"$:LWCP_v1:$x4wFOct/DGctv96o4IxxBB==",
        "senderNick":"blinkbean",
        "sessionWebhook":"https://oapi.dingtalk.com/robot/sendBySession?session=6d69b333f243db32d42c11sda9de6200",
        "sessionWebhookExpiredTime":1595237839032,
        "text":{
            "content":" outgoing"
        }
    }
    ```
  
- Usage
    ```go
    func OutGoing(ctx *gin.Context){
        cli := dingtalk.InitDingTalk([]string{"***"}, ".")
        msg, _ := cli.OutGoing(ctx.Request.Body)
        // 处理content
        res := doSomeThing(msg.Text.Content)
  
        textMsg := dingtalk.NewTextMsg(res)
        ctx.Set("respData", textMsg)
        ctx.JSON(ctx.Writer.Status(), data)
        return 
    }
    ```