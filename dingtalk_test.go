package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
var dingToken = []string{"b9230b8c762cb3a6f5dd977ad975c687e23fdefc6c762fe94a0f36bca73fb654"}
//var dingToken = []string{"0471c091acf1d9dfc553e57525f57139859858b905836b8f45b228e6d6e3a289"} // onlyOne

var dingTalkCli = InitDingTalk(dingToken, ".")
var dingTalkCliWithSecret = InitDingTalkWithSecret("47f656c278ff59b59e39486f98e5e7e1d1a3ef0bca22f9c01b37a90b81fe91aa", "SEC0fdc400a2b7e2f70f0b4d787a8f0623fe9742225bee747b4d5852162d24a63ed") // 加签

var testImg = "https://golang.google.cn/lib/godoc/images/footer-gopher.jpg"
var testUrl = "https://golang.google.cn/"
var testPhone = "1318282596*"

func init() {
	dingTalkCli = InitDingTalk(dingToken, ".")
}

func TestTextMsgWithSecret(t *testing.T) {
	err := dingTalkCliWithSecret.SendTextMessage("加签测试", WithAtMobiles([]string{testPhone}))
	assert.Equal(t, err, nil)
}

func TestTextMsg(t *testing.T) {
	err := dingTalkCli.SendTextMessage("Text 测试", WithAtMobiles([]string{testPhone}))
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

func TestSendDTMDMessage(t *testing.T) {
	// 有序map
	dtmdOrderMap := DingMap().
		Set("dtmdOrderMap1", "dtmdValue1").
		Set("dtmdOrderMap2", "dtmdValue2").
		Set("dtmdOrderMap3", "dtmdValue3")
	err := dingTalkCli.SendDTMDMessage("DTMD title", dtmdOrderMap)
	assert.Equal(t, err, nil)
}

func TestSendMarkDownMessageByList(t *testing.T) {
	msg := []string{
		"### Link test",
		"---",
		"- <font color=#00ff00 size=2>红色文字</font>",
		"- content2",
	}
	//err := dingTalkCli.SendMarkDownMessageBySlice("Markdown title", msg)
	err := dingTalkCli.SendMarkDownMessageBySlice("Markdown title", msg, WithAtMobiles([]string{testPhone}))
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
	err := dingTalkCli.SendMarkDownMessageBySlice("color test", dm.Slice())
	assert.Equal(t, err, nil)
}
