package dispatcher

import (
	"errors"
	"kinobot/pkg/tg/model"
	"sync"
)

var ErrNotMatch = errors.New("not match")

func NotMatch() error {
	return ErrNotMatch
}

type (
	ErrorHandler   func(err error)
	Handler[T any] func(update *T) error
)

type (
	MessageHandler       = Handler[model.Message]
	InlineQueryHandler   = Handler[model.InlineQuery]
	CallbackQueryHandler = Handler[model.CallbackQuery]
)

type Dispatcher struct {
	ErrorHandler          ErrorHandler
	messageHandlers       []MessageHandler
	inlineQueryHandlers   []InlineQueryHandler
	callbackQueryHandlers []CallbackQueryHandler
}

func (d *Dispatcher) AddMessageHandler(handler MessageHandler) {
	d.messageHandlers = append(d.messageHandlers, handler)
}

func (d *Dispatcher) AddInlineQueryHandler(handler InlineQueryHandler) {
	d.inlineQueryHandlers = append(d.inlineQueryHandlers, handler)
}

func (d *Dispatcher) AddCallbackQueryHandler(handler CallbackQueryHandler) {
	d.callbackQueryHandlers = append(d.callbackQueryHandlers, handler)
}

func findMatch[T any](update *T, handlers []Handler[T], errorHandler ErrorHandler) {
	for _, handler := range handlers {
		err := handler(update)
		if err == nil {
			return
		}
		if !errors.Is(err, ErrNotMatch) {
			if errorHandler != nil {
				errorHandler(err)
			}
			return
		}
	}
}

func UserId(update *model.Update) int64 {
	switch {
	case update.Message != nil && update.Message.From != nil:
		return update.Message.From.Id
	case update.InlineQuery != nil:
		return update.InlineQuery.From.Id
	case update.CallbackQuery != nil:
		return update.CallbackQuery.From.Id
	}
	return -1
}

func (d *Dispatcher) dispatch(update *model.Update) {
	switch {
	case update.Message != nil:
		findMatch(update.Message, d.messageHandlers, d.ErrorHandler)
	case update.InlineQuery != nil:
		findMatch(update.InlineQuery, d.inlineQueryHandlers, d.ErrorHandler)
	case update.CallbackQuery != nil:
		findMatch(update.CallbackQuery, d.callbackQueryHandlers, d.ErrorHandler)
	}
}

func (d *Dispatcher) Start(workerNum, perWorkerBuf int, updates <-chan *model.Update) {
	wg := sync.WaitGroup{}
	wg.Add(workerNum)

	channels := make([]chan *model.Update, workerNum)

	for i := range workerNum {
		updates := make(chan *model.Update, perWorkerBuf)

		go func() {
			defer wg.Done()
			for update := range updates {
				d.dispatch(update)
			}
		}()

		channels[i] = updates
	}

	for update := range updates {
		userId := UserId(update)
		if userId < 0 {
			continue
		}

		workerId := userId % int64(workerNum)
		channels[workerId] <- update
	}

	for _, channel := range channels {
		close(channel)
	}

	wg.Wait()
}
