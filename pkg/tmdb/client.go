package tmdb

import (
	"context"
	"kinobot/pkg/tmdb/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const api = "https://api.themoviedb.org/3/"

type Path string

const (
	SearchMulti   Path = "search/multi"
	TvTopRated    Path = "tv/top_rated"
	MovieTopRated Path = "movie/top_rated"
	TvDetails     Path = "tv/{series_id}"
	MovieDetails  Path = "movie/{movie_id}"
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

func applyPathParams(path Path, params ...string) Path {
	builder := strings.Builder{}
	builder.Grow(len(path))

	p := string(path)
	for _, param := range params {
		i := strings.IndexRune(p, '{')
		j := strings.IndexRune(p, '}')

		builder.WriteString(p[:i])
		builder.WriteString(param)

		p = p[j+1:]
	}
	builder.WriteString(p)

	return Path(builder.String())
}

func (tmdbc *TmdbClient) TvDetails(id int, params TvDetailsParams) (*model.Media, error) {
	return tmdbc.TvDetailsContext(context.Background(), id, params)
}

func (tmdbc *TmdbClient) TvDetailsContext(ctx context.Context, id int, params TvDetailsParams) (*model.Media, error) {
	path := applyPathParams(TvDetails, strconv.Itoa(id))

	var media *model.Media
	if err := tmdbc.doRequest(ctx, path, tmdbc.makeValues(params), &media); err != nil {
		return nil, err
	}

	return media, nil
}

func (tmdbc *TmdbClient) MovieDetails(id int, params MovieDetailsParams) (*model.Media, error) {
	return tmdbc.MovieDetailsContext(context.Background(), id, params)
}

func (tmdbc *TmdbClient) MovieDetailsContext(ctx context.Context, id int, params MovieDetailsParams) (*model.Media, error) {
	path := applyPathParams(MovieDetails, strconv.Itoa(id))

	var media *model.Media
	if err := tmdbc.doRequest(ctx, path, tmdbc.makeValues(params), &media); err != nil {
		return nil, err
	}

	return media, nil
}
