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
	GetUpdates        = "GetUpdates"
	SetWebhook        = "SetWebhook"
	DeleteWebhook     = "DeleteWebhook"
	SendMessage       = "SendMessage"
	AnswerInlineQuery = "AnswerInlineQuery"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type BotClient struct {
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

func NewBotClient(token string) (*BotClient, error) {
	id, err := parseId(token)
	if err != nil {
		return nil, err
	}

	bc := &BotClient{
		id:      id,
		base:    makeBase(token),
		httpc:   new(http.Client),
		updates: make(chan *model.Update),
	}

	return bc, nil
}

func (bc *BotClient) Id() int64 {
	return bc.id
}

func (bc *BotClient) Updates() <-chan *model.Update {
	return bc.updates
}

func (bc *BotClient) GetMe() (*model.User, error) {
	return bc.GetMeContext(context.Background())
}

func (bc *BotClient) GetMeContext(ctx context.Context) (*model.User, error) {
	var user *model.User
	if err := bc.doRequest(ctx, GetMe, nil, &user); err != nil {
		return nil, fmt.Errorf("%s: %w", GetMe, err)
	}
	return user, nil
}

func (bc *BotClient) GetUpdates(params GetUpdatesParams) (model.Updates, error) {
	return bc.GetUpdatesContext(context.Background(), params)
}

func (bc *BotClient) GetUpdatesContext(ctx context.Context, params GetUpdatesParams) (model.Updates, error) {
	var updates model.Updates
	if err := bc.doRequest(ctx, GetUpdates, params, &updates); err != nil {
		return nil, fmt.Errorf("%s: %w", GetUpdates, err)
	}
	return updates, nil
}

func (bc *BotClient) SetWebhook(params SetWebhookParams) (bool, error) {
	return bc.SetWebhookContext(context.Background(), params)
}

func (bc *BotClient) SetWebhookContext(ctx context.Context, params SetWebhookParams) (bool, error) {
	var ok bool
	if err := bc.doRequest(ctx, SetWebhook, params, &ok); err != nil {
		return false, fmt.Errorf("%s: %w", SetWebhook, err)
	}
	return ok, nil
}

func (bc *BotClient) DeleteWebhook() (bool, error) {
	return bc.DeleteWebhookContext(context.Background())
}

func (bc *BotClient) DeleteWebhookContext(ctx context.Context) (bool, error) {
	var ok bool
	if err := bc.doRequest(ctx, DeleteWebhook, nil, &ok); err != nil {
		return false, fmt.Errorf("%s: %w", DeleteWebhook, err)
	}
	return ok, nil
}

func (bc *BotClient) SendMessage(params SendMessageParams) (*model.Message, error) {
	return bc.SendMessageContext(context.Background(), params)
}

func (bc *BotClient) SendMessageContext(ctx context.Context, params SendMessageParams) (*model.Message, error) {
	var message *model.Message
	if err := bc.doRequest(ctx, SendMessage, params, &message); err != nil {
		return nil, fmt.Errorf("%s: %w", SendMessage, err)
	}
	return message, nil
}

func (bc *BotClient) AnswerInlineQuery(params AnswerInlineQueryParams) (bool, error) {
	return bc.AnswerInlineQueryContext(context.Background(), params)
}

func (bc *BotClient) AnswerInlineQueryContext(ctx context.Context, params AnswerInlineQueryParams) (bool, error) {
	var ok bool
	if err := bc.doRequest(ctx, AnswerInlineQuery, params, &ok); err != nil {
		return false, fmt.Errorf("%s: %w", AnswerInlineQuery, err)
	}
	return ok, nil
}
