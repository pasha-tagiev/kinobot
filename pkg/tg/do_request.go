package tg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrUnexpectedEntity = errors.New("unexpected entity")

type ResponseError struct {
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("telegram bot api [%d]: %s", re.ErrorCode, re.Description)
}

type Wrapper struct {
	*ResponseError
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

func (bc *BotClient) doRequest(ctx context.Context, method Method, params any, data any) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	endpoint := bc.base + string(method)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := bc.httpc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var wrapper Wrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return err
	}

	if !wrapper.Ok {
		if wrapper.ResponseError != nil {
			return wrapper.ResponseError
		}
		return ErrUnexpectedEntity
	}

	if err := json.Unmarshal(wrapper.Result, data); err != nil {
		return err
	}

	return nil
}
