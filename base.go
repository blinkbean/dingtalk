package dingtalk

type MsgTypeType string

const (
	TEXT       MsgTypeType = "text"
	LINK       MsgTypeType = "link"
	MARKDOWN   MsgTypeType = "markdown"
	ACTION_CARD MsgTypeType = "actionCard"
	FEED_CARD   MsgTypeType = "feedCard"
)

type textModel struct {
	Content string `json:"content,omitempty"`
}

type atModel struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type linkModel struct {
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	MessageUrl string `json:"messageUrl,omitempty"`
}

type markDownModel struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type actionCardBtnOrientationType string

const (
	HORIZONTAL actionCardBtnOrientationType = "0" // 横向
	VERTICAL   actionCardBtnOrientationType = "1" // 竖向
)

type ActionCardModel struct {
	Title          string                       `json:"title,omitempty"`
	Text           string                       `json:"text,omitempty"`
	BtnOrientation actionCardBtnOrientationType `json:"btnOrientation,omitempty"`
	SingleTitle    string                       `json:"singleTitle,omitempty"`
	SingleURL      string                       `json:"singleURL,omitempty"`
	Btns           []ActionCardMultiBtnModel    `json:"btns,omitempty"`
}

type ActionCardMultiBtnModel struct {
	Title     string `json:"title,omitempty"`
	ActionURL string `json:"actionURL,omitempty"`
}

type feedCardModel struct {
	Links []FeedCardLinkModel `json:"links,omitempty"`
}

type FeedCardLinkModel struct {
	Title      string `json:"title,omitempty"`
	MessageURL string `json:"messageURL,omitempty"`
	PicURL     string `json:"picURL,omitempty"`
}
