package model

type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle InlineQueryResultType = "article"
)

type (
	InlineQueryResult any

	InlineQueryResultArticle struct {
		Type                InlineQueryResultType `json:"type"`
		Id                  string                `json:"id"`
		Title               string                `json:"title"`
		InputMessageContent InputMessageContent   `json:"input_message_content"`
		Description         string                `json:"description,omitempty"`
		ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	}
)

type ParseMode string

const (
	ParseModeHtml       = "HTML"
	ParseModeMarkdownV2 = "MarkdownV2"
)

type (
	InputMessageContent any

	InputTextMessageContent struct {
		MessageText string    `json:"message_text"`
		ParseMode   ParseMode `json:"parse_mode,omitempty"`
	}
)
