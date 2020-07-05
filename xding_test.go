package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dingToken = []string{"7bd675b66646ba890046c2198257576470099e1bda0770bad7dd6684fb1e0415"}

var dingTalkCli = InitDingTalk(dingToken, ".")

var testImg = "https://golang.google.cn/lib/godoc/images/footer-gopher.jpg"
var testUrl = "https://golang.google.cn/"
var testPhone = "13182825966"

func TestTextMsg(t *testing.T) {
	err := dingTalkCli.sendMessage(NewTextMsg("Text 测试.", WithAtMobiles([]string{testPhone})))
	assert.Equal(t, err, nil)
}

func TestLinkMsg(t *testing.T) {
	err := dingTalkCli.sendMessage(NewLinkMsg("Link title", "Link test.", testImg, testUrl))
	assert.Equal(t, err, nil)
}

func TestMarkDownMsg(t *testing.T) {
	err := dingTalkCli.sendMessage(NewMarkDownMsg("Markdown title", "### Link test\n - content1 \n - content2.", WithAtAll()))
	assert.Equal(t, err, nil)
}

func TestActionCardMultiMsg(t *testing.T) {
	Btns := []ActionCardMultiBtnModel{{
		Title:     "test1",
		ActionURL: testUrl,
	}, {
		Title:     "test2",
		ActionURL: testUrl,
	},
	}
	//err := dingTalkCli.sendMessage(NewActionCardMsg("ActionCard title", "ActionCard text.", WithCardSingleTitle("title"), WithCardSingleURL(testUrl)))
	err := dingTalkCli.sendMessage(NewActionCardMsg("ActionCard title", "ActionCard text.", WithCardBtns(Btns), WithCardBtnVertical()))
	assert.Equal(t, err, nil)
}

func TestFeedCardMsg(t *testing.T) {
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
	err := dingTalkCli.sendMessage(NewFeedCardMsg(links))
	assert.Equal(t, err, nil)
}
