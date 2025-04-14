package tmdb

type SearchMultiParams struct {
	Query        string `url:"query"`
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Language     string `url:"language,omitempty"`
}

type TvTopRatedParams struct {
	Language string `url:"language"`
}

type MovieTopRatedParams struct {
	Language string `url:"language"`
	Region   string `url:"region"`
}
