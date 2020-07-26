package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dingToken = []string{"7bd675b66646ba890046c2198257576470099e1bda0770bad7dd6684fb1e0415"}

var dingTalkCli = InitDingTalk(dingToken, ".")

var testImg = "https://golang.google.cn/lib/godoc/images/footer-gopher.jpg"
var testUrl = "https://golang.google.cn/"
var testPhone = "1318282596*"

func init() {
	dingTalkCli = InitDingTalk(dingToken, ".")
}
func TestTextMsg(t *testing.T) {
	err := dingTalkCli.SendTextMessage("Text 测试.", WithAtMobiles([]string{testPhone}))
	assert.Equal(t, err, nil)
}

func TestLinkMsg(t *testing.T) {
	err := dingTalkCli.SendLinkMessage("Link title", "Link test.", testImg, testUrl)
	assert.Equal(t, err, nil)
}

func TestMarkDownMsg(t *testing.T) {
	err := dingTalkCli.SendMarkDownMessage("Markdown title", "### Link test\n --- \n- <font color=#00ff00 size=6>红色文字</font> \n - content2.", WithAtAll())
	assert.Equal(t, err, nil)
}

func TestSendMarkDownMessageByList(t *testing.T) {
	msg := []string{
		"### Link test",
		"---",
		"- <font color=#00ff00 size=6>红色文字</font>",
		"- content2",
	}
	err := dingTalkCli.SendMarkDownMessageBySlice("Markdown title", msg, WithAtAll())
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
	//err := dingTalkCli.SendActionCardMessage("ActionCard title", "ActionCard text.", WithCardSingleTitle("title"), WithCardSingleURL(testUrl))
	err := dingTalkCli.SendActionCardMessage("ActionCard title", "- ActionCard text.", WithCardBtns(Btns), WithCardBtnVertical())
	assert.Equal(t, err, nil)
}

func TestActionCardMultiMsgBySlice(t *testing.T) {
	Btns := []ActionCardMultiBtnModel{{
		Title:     "test1",
		ActionURL: testUrl,
	}, {
		Title:     "test2",
		ActionURL: testUrl,
	},
	}
	dm := DingMap()
	dm.Set("颜色测试", H2)
	dm.Set("失败：$$ 同行不同色 $$", RED)
	dm.Set("---", "")
	dm.Set("金色", GOLD)
	dm.Set("成功", GREEN)
	dm.Set("警告", BLUE)
	dm.Set("普通文字", N)
	//err := dingTalkCli.SendActionCardMessage("ActionCard title", "ActionCard text.", WithCardSingleTitle("title"), WithCardSingleURL(testUrl))
	err := dingTalkCli.SendActionCardMessageBySlice("ActionCard title", dm.Slice(), WithCardBtns(Btns), WithCardBtnVertical())
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
	err := dingTalkCli.SendFeedCardMessage(links)
	assert.Equal(t, err, nil)
}

func TestDingMap(t *testing.T) {
	dm := DingMap()
	dm.Set("颜色测试", H2)
	dm.Set("失败：$$ 同行不同色 $$", RED)
	dm.Set("---", "")
	dm.Set("金色", GOLD)
	dm.Set("成功", GREEN)
	dm.Set("警告", BLUE)
	dm.Set("普通文字", N)
	dingTalkCli.SendMarkDownMessageBySlice("color test", dm.Slice())
}
