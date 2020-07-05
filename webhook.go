package dingtalk

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type dingTalk struct {
	robotToken []string
}

var keyWord = "."
var dingTalkClient = dingTalk{}

func InitDingTalk(tokens []string, key string) *dingTalk {
	if len(tokens) == 0 {
		panic("no token")
	}
	if keyWord != "" {
		keyWord = key
	}
	dingTalkClient = dingTalk{
		robotToken: tokens,
	}
	return &dingTalkClient
}

func (d *dingTalk) sendMessage(msg iDingMsg) error {
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
	resp, err = doRequest(ctx, "POST", url, header, msg.Marshaler())

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("send msg err: %s, token: %s, msg: %s", string(body), d.robotToken, msg.Marshaler())
	}
	return nil
}

func (d *dingTalk) SendTextMessage(content string, opt ...atOption) error {
	content = content + keyWord
	return dingTalkClient.sendMessage(NewTextMsg(content, opt...))
}

func (d *dingTalk) SendMarkDownMessage(title, text string, opts ...atOption) error {
	title = title + keyWord
	return dingTalkClient.sendMessage(NewMarkDownMsg(title, text, opts...))
}

func (d *dingTalk) SendLinkMessage(title, text, picUrl, msgUrl string) error {
	title = title + keyWord
	return dingTalkClient.sendMessage(NewLinkMsg(title, text, picUrl, msgUrl))
}

func (d *dingTalk) SendActionSingleMessage(title, text string, opts ...actionCardOption) error {
	title = title + keyWord
	return dingTalkClient.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *dingTalk) SendFeedCardMessage(feedCard []FeedCardLinkModel) error {
	if len(feedCard) > 0 {
		feedCard[0].Title = feedCard[0].Title + keyWord
	}
	return dingTalkClient.sendMessage(NewFeedCardMsg(feedCard))
}
