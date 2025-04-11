package tg

import (
	"context"
	"encoding/json"
	"kinobot/pkg/tg/model"
	"net/http"
)

const WebhookSecretTokenHeader = "X-Telegram-Bot-Api-Secret-Token"

type WebhookHandler struct {
	secretToken string
	updates     chan *model.Update
	done        chan struct{}
}

func (wh *WebhookHandler) forwardUpdates(ctx context.Context, dst chan<- *model.Update) {
	defer func() {
		close(wh.done)
		close(dst)
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case update := <-wh.updates:
			dst <- update
		}
	}
}

func (wh *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	select {
	case <-wh.done:
		w.WriteHeader(http.StatusServiceUnavailable)
		return

	default:
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if wh.secretToken != "" && wh.secretToken != r.Header.Get(WebhookSecretTokenHeader) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var update *model.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	select {
	case <-wh.done:
		w.WriteHeader(http.StatusServiceUnavailable)

	case wh.updates <- update:
		w.WriteHeader(http.StatusOK)
	}
}

func (b *Bot) WebhookHandler(ctx context.Context, secretToken string) http.Handler {
	wh := &WebhookHandler{
		secretToken: secretToken,
		updates:     make(chan *model.Update),
		done:        make(chan struct{}),
	}

	go wh.forwardUpdates(ctx, b.updates)

	return wh
}
