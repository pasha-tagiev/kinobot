package model

type Updates []*Update

type Update struct {
	Id            int64          `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	InlineQuery   *InlineQuery   `json:"inline_query,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
	Id     int64  `json:"message_id"`
	From   *User  `json:"from,omitempty"`
	Date   int64  `json:"date"`
	Chat   Chat   `json:"chat"`
	ViaBot *User  `json:"via_bot,omitempty"`
	Text   string `json:"text,omitempty"`
}

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType ChatType `json:"chat_type,omitempty"`
}

type CallbackQuery struct {
	Id           string `json:"id"`
	From         User   `json:"from"`
	ChatInstance string `json:"chat_instance"`
	Data         string `json:"data,omitempty"`
}

type Chat struct {
	Id   int64    `json:"id"`
	Type ChatType `json:"type"`
}

type User struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}
