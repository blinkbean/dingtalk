package dingtalk

import jsoniter "github.com/json-iterator/go"

type IDingMsg interface {
	Marshaler() []byte
}

type AtOption interface {
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

func WithAtAll() AtOption{
	return newFuncAtOption(func(o *atModel) {
		o.IsAtAll=true
	})
}

func WithAtMobiles(mobiles []string) AtOption{
	return newFuncAtOption(func(o *atModel) {
		o.AtMobiles=mobiles
	})
}

type textMsg struct {
	MsgType MsgTypeType `json:"msgtype,omitempty"`
	Text    textModel   `json:"text,omitempty"`
	At      atModel     `json:"at,omitempty"`
}

func (t textMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(t)
	return b
}

func NewTextMsg(content string, opts ...AtOption) *textMsg {
	msg := &textMsg{MsgType: TEXT, Text:textModel{Content:content}}
	for _,opt := range opts{
		opt.apply(&msg.At)
	}
	return msg
}

type linkMsg struct {
	MsgType MsgTypeType `json:"msgtype,omitempty"`
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
	MsgType  MsgTypeType   `json:"msgtype,omitempty"`
	Markdown markDownModel `json:"markdown,omitempty"`
	At       atModel       `json:"at,omitempty"`
}

func (m markDownMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(m)
	return b
}

func NewMarkDownMsg(title string, text interface{}, opts ...AtOption) *markDownMsg {

	msg := &markDownMsg{MsgType: MARKDOWN, Markdown:markDownModel{Title:title, Text:text.(string)}}
	for _,opt := range opts{
		opt.apply(&msg.At)
	}
	return msg
}

type ActionCardOption interface {
	apply(model *ActionCardModel)
}

type funcActionCardOption struct {
	f func(model *ActionCardModel)
}

func (fdo *funcActionCardOption) apply(do *ActionCardModel)  {
	fdo.f(do)
}

func newFuncActionCardOption(f func(model *ActionCardModel)) *funcActionCardOption{
	return &funcActionCardOption{f:f}
}


func WithCardBtnOrientation(orientation actionCardBtnOrientationType) ActionCardOption{
	return newFuncActionCardOption(func(o *ActionCardModel) {
		o.BtnOrientation=orientation
	})
}

func WithCardSingleTitle(title string) ActionCardOption{
	return newFuncActionCardOption(func(o *ActionCardModel) {
		o.SingleTitle=title
	})
}

func WithCardSingleURL(url string) ActionCardOption{
	return newFuncActionCardOption(func(o *ActionCardModel) {
		o.SingleURL=url
	})
}

func WithCardBtns(btns []ActionCardMultiBtnModel)ActionCardOption{
	return newFuncActionCardOption(func(o *ActionCardModel) {
		o.Btns=btns
	})
}

type actionCardMsg struct {
	MsgType    MsgTypeType     `json:"msgtype,omitempty"`
	ActionCard ActionCardModel `json:"actionCard,omitempty"`
}

func (a actionCardMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(a)
	return b
}

func NewActionCardMsg(title, text string, opts ...ActionCardOption) *actionCardMsg {
	card := &actionCardMsg{MsgType: ACTION_CARD, ActionCard:ActionCardModel{
		Title:          title,
		Text:           text,
		BtnOrientation: HORIZONTAL,
	}}
	for _, opt := range opts{
		opt.apply(&card.ActionCard)
	}
	return card
}


type FeedCardMsg struct {
	MsgType  MsgTypeType   `json:"msgtype,omitempty"`
	FeedCard feedCardModel `json:"feedCard,omitempty"`
}

func (f FeedCardMsg) Marshaler() []byte {
	b, _ := jsoniter.Marshal(f)
	return b
}

func NewFeedCardMsg(feedCard []FeedCardLinkModel) *FeedCardMsg {
	return &FeedCardMsg{MsgType: FEED_CARD, FeedCard:feedCardModel{Links:feedCard}}
}
