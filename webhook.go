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

var keyWord = "."

func InitDingTalk(tokens []string, key string) *DingTalk {
	if len(tokens) == 0 {
		panic("no token")
	}
	if keyWord != "" {
		keyWord = key
	}
	return &DingTalk{
		robotToken: tokens,
	}
}

func (d *DingTalk) sendMessage(msg iDingMsg) error {
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

func (d *DingTalk) SendTextMessage(content string, opt ...atOption) error {
	content = content + keyWord
	return d.sendMessage(NewTextMsg(content, opt...))
}

func (d *DingTalk) SendMarkDownMessage(title, text string, opts ...atOption) error {
	title = title + keyWord
	return d.sendMessage(NewMarkDownMsg(title, text, opts...))
}

func (d *DingTalk) SendLinkMessage(title, text, picUrl, msgUrl string) error {
	title = title + keyWord
	return d.sendMessage(NewLinkMsg(title, text, picUrl, msgUrl))
}

func (d *DingTalk) SendActionCardMessage(title, text string, opts ...actionCardOption) error {
	title = title + keyWord
	return d.sendMessage(NewActionCardMsg(title, text, opts...))
}

func (d *DingTalk) SendFeedCardMessage(feedCard []FeedCardLinkModel) error {
	if len(feedCard) > 0 {
		feedCard[0].Title = feedCard[0].Title + keyWord
	}
	return d.sendMessage(NewFeedCardMsg(feedCard))
}
