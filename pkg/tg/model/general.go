package model

type Update struct {
	Id      int64    `json:"update_id"`
	Message *Message `json:"message,omitempty"`
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
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSuperGroup = "supergroup"
	ChatTypeChannel    = "channel"
)

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
