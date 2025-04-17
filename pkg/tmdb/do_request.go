package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ResponseError struct {
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("tmdb api [%d]: %s", re.StatusCode, re.StatusMessage)
}

func (tmdbc *TmdbClient) doRequest(ctx context.Context, path Path, params url.Values, data any) error {
	endpoint := api + string(path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Accept", "application/json")

	resp, err := tmdbc.httpc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr *ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return err
		}
		return respErr
	}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}
