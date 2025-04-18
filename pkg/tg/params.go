package tg

import "kinobot/pkg/tg/model"

type GetUpdatesParams struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SetWebhookParams struct {
	Url                string   `json:"url"`
	MaxConnections     int      `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool     `json:"drop_pending_updates,omitempty"`
	SecretToken        string   `json:"secret_token,omitempty"`
}

type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

type SendMessageParams struct {
	ChatId    int64           `json:"chat_id"`
	Text      string          `json:"text"`
	ParseMode model.ParseMode `json:"parse_mode,omitempty"`
}

type AnswerInlineQueryParams struct {
	Id         string                    `json:"inline_query_id"`
	Results    []model.InlineQueryResult `json:"results,omitzero"`
	CacheTime  int                       `json:"cache_time,omitempty"`
	IsPersonal bool                      `json:"is_personal,omitempty"`
}
