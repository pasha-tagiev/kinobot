package tmdb

import (
	"context"
	"iter"
	"kinobot/pkg/tmdb/model"
	"net/url"
	"strconv"
)

type MediaSeq = iter.Seq2[*model.Media, error]

type MediaStream struct {
	tmdbc  *TmdbClient
	path   Path
	values url.Values
}

func (ms *MediaStream) All() MediaSeq {
	return ms.AllContext(context.Background())
}

func (ms *MediaStream) AllContext(ctx context.Context) MediaSeq {
	return func(yield func(*model.Media, error) bool) {
		for i := 1; ; i++ {
			ms.values.Set("page", strconv.Itoa(i))

			var page model.Page
			if err := ms.tmdbc.doRequest(ctx, ms.path, ms.values, &page); err != nil {
				yield(nil, err)
				return
			}
			if len(page.Results) == 0 {
				return
			}

			for _, m := range page.Results {
				if err := ctx.Err(); err != nil {
					yield(nil, err)
					return
				}
				if !yield(m, nil) {
					return
				}
			}

			if i >= page.TotalPages {
				return
			}
		}
	}
}
