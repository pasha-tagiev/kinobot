package tg

import (
	"context"
	"errors"
	"kinobot/pkg/tg/model"

	"github.com/cenkalti/backoff/v5"
)

func (bc *BotClient) Polling(ctx context.Context, maxTries uint, params GetUpdatesParams) error {
	defer close(bc.updates)

	operation := func() (model.Updates, error) {
		updates, err := bc.GetUpdatesContext(ctx, params)
		if err == nil {
			return updates, nil
		}

		var rerr *ResponseError
		if errors.As(err, &rerr) || errors.Is(err, ErrUnexpectedEntity) {
			return nil, backoff.Permanent(err)
		}

		return nil, err
	}

	for {
		updates, err := backoff.Retry(ctx, operation,
			backoff.WithMaxTries(maxTries),
			backoff.WithBackOff(backoff.NewExponentialBackOff()),
		)
		if err != nil {
			return err
		}
		if len(updates) == 0 {
			continue
		}

		params.Offset = updates[len(updates)-1].Id + 1

		for _, update := range updates {
			select {
			case <-ctx.Done():
				return context.Cause(ctx)

			case bc.updates <- update:
				continue
			}
		}
	}
}
