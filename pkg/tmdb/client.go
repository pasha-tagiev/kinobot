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

type TmdbClient struct {
	apiKey string
	httpc  HttpClient
}

func NewTmdbClient(apiKey string, httpc HttpClient) *TmdbClient {
	return &TmdbClient{
		apiKey: apiKey,
		httpc:  httpc,
	}
}

func (tmdbc *TmdbClient) makeValues(params any) url.Values {
	values, _ := query.Values(params)
	values.Set("api_key", tmdbc.apiKey)
	return values
}

func (tmdbc *TmdbClient) SearchMulti(params SearchMultiParams) *MediaStream {
	return &MediaStream{
		tmdbc:  tmdbc,
		path:   SearchMulti,
		values: tmdbc.makeValues(params),
	}
}

func (tmdbc *TmdbClient) TvTopRated(params TvTopRatedParams) *MediaStream {
	return &MediaStream{
		tmdbc:  tmdbc,
		path:   TvTopRated,
		values: tmdbc.makeValues(params),
	}
}

func (tmdbc *TmdbClient) MovieTopRated(params MovieTopRatedParams) *MediaStream {
	return &MediaStream{
		tmdbc:  tmdbc,
		path:   MovieTopRated,
		values: tmdbc.makeValues(params),
	}
}
