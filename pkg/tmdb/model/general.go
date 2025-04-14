package model

type Page struct {
	Page         int      `json:"page"`
	Results      []*Media `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type MediaType string

const (
	MediaTypeTv     = "tv"
	MediaTypeMovie  = "movie"
	MediaTypePerson = "person"
)

type ImagePath string

type Media struct {
	Id            int       `json:"id"`
	MediaType     MediaType `json:"media_type"`
	Name          string    `json:"name"`
	OriginalName  string    `json:"original_name"`
	Title         string    `json:"title"`
	OriginalTitle string    `json:"original_title"`
	ReleaseDate   string    `json:"release_date"`
	FirstAirDate  string    `json:"first_air_date"`
	Overview      string    `json:"overview"`
	Popularity    float64   `json:"popularity"`
	PosterPath    ImagePath `json:"poster_path"`
}
