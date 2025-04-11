package tg

type SetWebhookParams struct {
	Url                string   `json:"url"`
	MaxConnections     int      `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool     `json:"drop_pending_updates,omitempty"`
	SecretToken        string   `json:"secret_token,omitempty"`
}

type ParseMode string

const (
	ParseModeHtml       = "HTML"
	ParseModeMarkdownV2 = "MarkdownV2"
)

type SendMessageParams struct {
	ChatId    int64     `json:"chat_id"`
	Text      string    `json:"text"`
	ParseMode ParseMode `json:"parse_mode,omitempty"`
}
