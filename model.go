package dingtalk

import jsoniter "github.com/json-iterator/go"

type iDingMsg interface {
	Marshaler() []byte
}

type atOption interface {
	apply(model *atModel)
}

type funcAtOption struct {
	f func(model *atModel)
}

func (fdo *funcAtOption) apply(do *atModel)  {
	fdo.f(do)
}

func newFuncAtOption(f func(model *atModel)) *funcAtOption{
	return &funcAtOption{f:f}
}

func WithAtAll() atOption{
	return newFuncAtOption(func(o *atModel) {
		o.IsAtAll=true
	})
}

func WithAtMobiles(mobiles []string) atOption{
	return newFuncAtOption(func(o *atModel) {
		o.AtMobiles=mobiles
	})
}

type textMsg struct {
	MsgType msgTypeType `json:"msgtype,omitempty"`
	Text    textModel   `json:"text,omitempty"`
	At      atModel     `json:"at,omitempty"`
}

func (t textMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(t)
	return b
}

func NewTextMsg(content string, opts ...atOption) *textMsg {
	msg := &textMsg{MsgType: TEXT, Text:textModel{Content:content}}
	for _,opt := range opts{
		opt.apply(&msg.At)
	}
	return msg
}

type linkMsg struct {
	MsgType msgTypeType `json:"msgtype,omitempty"`
	Link    linkModel   `json:"link,omitempty"`
}

func (l linkMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(l)
	return b
}

func NewLinkMsg(title, text, picUrl, msgUrl string) *linkMsg {
	return &linkMsg{MsgType: LINK, Link:linkModel{
		Text:       text,
		Title:      title,
		PicUrl:     picUrl,
		MessageUrl: msgUrl,
	}}
}

type markDownMsg struct {
	MsgType  msgTypeType   `json:"msgtype,omitempty"`
	Markdown markDownModel `json:"markdown,omitempty"`
	At       atModel       `json:"at,omitempty"`
}

func (m markDownMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(m)
	return b
}

func NewMarkDownMsg(title string, text interface{}, opts ...atOption) *markDownMsg {

	msg := &markDownMsg{MsgType: MARKDOWN, Markdown:markDownModel{Title:title, Text:text.(string)}}
	for _,opt := range opts{
		opt.apply(&msg.At)
	}
	return msg
}

type actionCardOption interface {
	apply(model *actionCardModel)
}

type funcActionCardOption struct {
	f func(model *actionCardModel)
}

func (fdo *funcActionCardOption) apply(do *actionCardModel)  {
	fdo.f(do)
}

func newFuncActionCardOption(f func(model *actionCardModel)) *funcActionCardOption{
	return &funcActionCardOption{f:f}
}


func WithCardBtnVertical() actionCardOption{
	return newFuncActionCardOption(func(o *actionCardModel) {
		o.BtnOrientation=vertical
	})
}

func WithCardSingleTitle(title string) actionCardOption{
	return newFuncActionCardOption(func(o *actionCardModel) {
		o.SingleTitle=title
	})
}

func WithCardSingleURL(url string) actionCardOption{
	return newFuncActionCardOption(func(o *actionCardModel) {
		o.SingleURL=url
	})
}

func WithCardBtns(btns []actionCardMultiBtnModel)actionCardOption{
	return newFuncActionCardOption(func(o *actionCardModel) {
		o.Btns=btns
	})
}

type actionCardMsg struct {
	MsgType    msgTypeType     `json:"msgtype,omitempty"`
	ActionCard actionCardModel `json:"actionCard,omitempty"`
}

func (a actionCardMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(a)
	return b
}

func NewActionCardMsg(title, text string, opts ...actionCardOption) *actionCardMsg {
	card := &actionCardMsg{MsgType: ACTION_CARD, ActionCard:actionCardModel{
		Title:          title,
		Text:           text,
		BtnOrientation: horizontal,
	}}
	for _, opt := range opts{
		opt.apply(&card.ActionCard)
	}
	return card
}


type feedCardMsg struct {
	MsgType  msgTypeType   `json:"msgtype,omitempty"`
	FeedCard feedCardModel `json:"feedCard,omitempty"`
}

func (f feedCardMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(f)
	return b
}

func NewFeedCardMsg(feedCard []feedCardLinkModel) *feedCardMsg {
	return &feedCardMsg{MsgType: FEED_CARD, FeedCard:feedCardModel{Links:feedCard}}
}
