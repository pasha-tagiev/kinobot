package tmdb

import "kinobot/pkg/tmdb/model"

const base = "https://image.tmdb.org/t/p/"

type ImageSize string

const (
	W342     ImageSize = "w342"
	W500     ImageSize = "w500"
	W780     ImageSize = "w780"
	Original ImageSize = "original"
)

func ImageUrl(size ImageSize, path model.ImagePath) string {
	return base + string(size) + string(path)
}
