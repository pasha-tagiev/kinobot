package tg

import (
	"context"
	"errors"
	"fmt"
	"kinobot/pkg/tg/model"
	"net/http"
	"strconv"
	"strings"
)

var ErrBadToken = errors.New("bad token")

const api = "https://api.telegram.org/"

type Method string

const (
	GetMe             = "GetMe"
	SetWebhook        = "SetWebhook"
	SendMessage       = "SendMessage"
	AnswerInlineQuery = "AnswerInlineQuery"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Bot struct {
	id      int64
	base    string
	httpc   HttpClient
	updates chan *model.Update
}

func parseId(token string) (int64, error) {
	strId, _, found := strings.Cut(token, ":")
	if strId == "" || !found {
		return 0, fmt.Errorf("%w: %q", ErrBadToken, token)
	}

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrBadToken, err)
	}

	return id, nil
}

func makeBase(token string) string {
	return api + "bot" + token + "/"
}

func New(token string) (*Bot, error) {
	id, err := parseId(token)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		id:      id,
		base:    makeBase(token),
		httpc:   new(http.Client),
		updates: make(chan *model.Update),
	}

	return b, nil
}

func (b *Bot) Id() int64 {
	return b.id
}

func (b *Bot) Updates() <-chan *model.Update {
	return b.updates
}

func (b *Bot) GetMe() (*model.User, error) {
	return b.GetMeContext(context.Background())
}

func (b *Bot) GetMeContext(ctx context.Context) (*model.User, error) {
	var user *model.User
	if err := b.doRequest(ctx, GetMe, nil, &user); err != nil {
		return nil, fmt.Errorf("%s: %w", GetMe, err)
	}
	return user, nil
}

func (b *Bot) SetWebhook(params SetWebhookParams) (bool, error) {
	return b.SetWebhookContext(context.Background(), params)
}

func (b *Bot) SetWebhookContext(ctx context.Context, params SetWebhookParams) (bool, error) {
	var ok bool
	if err := b.doRequest(ctx, SetWebhook, params, &ok); err != nil {
		return false, fmt.Errorf("%s: %w", SetWebhook, err)
	}
	return ok, nil
}

func (b *Bot) SendMessage(params SendMessageParams) (*model.Message, error) {
	return b.SendMessageContext(context.Background(), params)
}

func (b *Bot) SendMessageContext(ctx context.Context, params SendMessageParams) (*model.Message, error) {
	var message *model.Message
	if err := b.doRequest(ctx, SendMessage, params, &message); err != nil {
		return nil, fmt.Errorf("%s: %w", SendMessage, err)
	}
	return message, nil
}

func (b *Bot) AnswerInlineQuery(params AnswerInlineQueryParams) (bool, error) {
	return b.AnswerInlineQueryContext(context.Background(), params)
}

func (b *Bot) AnswerInlineQueryContext(ctx context.Context, params AnswerInlineQueryParams) (bool, error) {
	var ok bool
	if err := b.doRequest(ctx, AnswerInlineQuery, params, &ok); err != nil {
		return false, fmt.Errorf("%s: %w", AnswerInlineQuery, err)
	}
	return ok, nil
}
