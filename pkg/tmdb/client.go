package tmdb

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const api = "https://api.themoviedb.org/3/"

type Path string

const (
	SearchMulti   Path = "search/multi"
	TvTopRated    Path = "tv/top_rated"
	MovieTopRated Path = "movie/top_rated"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Tmdb struct {
	apiKey string
	httpc  HttpClient
}

func New(apiKey string, httpc HttpClient) *Tmdb {
	return &Tmdb{
		apiKey: apiKey,
		httpc:  httpc,
	}
}

func (tmdb *Tmdb) makeValues(params any) url.Values {
	values, _ := query.Values(params)
	values.Set("api_key", tmdb.apiKey)
	return values
}

func (tmdb *Tmdb) SearchMulti(params SearchMultiParams) *MediaStream {
	return &MediaStream{
		tmdb:   tmdb,
		path:   SearchMulti,
		values: tmdb.makeValues(params),
	}
}

func (tmdb *Tmdb) TvTopRated(params TvTopRatedParams) *MediaStream {
	return &MediaStream{
		tmdb:   tmdb,
		path:   TvTopRated,
		values: tmdb.makeValues(params),
	}
}

func (tmdb *Tmdb) MovieTopRated(params MovieTopRatedParams) *MediaStream {
	return &MediaStream{
		tmdb:   tmdb,
		path:   MovieTopRated,
		values: tmdb.makeValues(params),
	}
}
