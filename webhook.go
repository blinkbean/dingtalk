package dingtalk

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type DingTalk struct {
	robotToken []string
}

var KEYWORD = "."
var dingTalkClient = DingTalk{}

func InitDingTalk(tokens []string, keyWord string) *DingTalk {
	if len(tokens) == 0 {
		panic("no token")
	}
	if keyWord != "" {
		KEYWORD = keyWord
	}
	dingTalkClient = DingTalk{
		robotToken: tokens,
	}
	return &dingTalkClient
}

func (d *DingTalk) sendMessage(msg IDingMsg) error {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		url    string
		resp   *http.Response
		err    error
	)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	url = dingTalkURL + d.robotToken[rand.Intn(len(d.robotToken))]
	header := map[string]string{
		"Content-type": "application/json",
	}
	resp, err = DoRequest(ctx, "POST", url, header, msg.Marshaler())

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("send msg err: %s, token: %s, msg: %s", string(body), d.robotToken, msg.Marshaler())
	}
	return nil
}

func (d *DingTalk) SendTextMessage(content string, opt ...AtOption) error {
	content = content + KEYWORD
	return dingTalkClient.sendMessage(NewTextMsg(content, opt...))
}

func (d *DingTalk) SendMarkDownMessage(title, text string, opts ...AtOption) error {
	title = title + KEYWORD
	return dingTalkClient.sendMessage(NewMarkDownMsg(title, text, opts...))
}

func (d *DingTalk) SendLinkMessage(title, text, picUrl, msgUrl string) error {
	title = title + KEYWORD
	return dingTalkClient.sendMessage(NewLinkMsg(title, text, picUrl, msgUrl))
}

func (d *DingTalk) SendActionSingleMessage(title, text string, opts ...ActionCardOption) error {
	title = title + KEYWORD
	return dingTalkClient.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendFeedCardMessage(feedCard []FeedCardLinkModel) error {
	if len(feedCard) > 0 {
		feedCard[0].Title = feedCard[0].Title + KEYWORD
	}
	return dingTalkClient.sendMessage(NewFeedCardMsg(feedCard))
}
