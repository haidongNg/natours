package models

import "time"

type Tour struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Duration        int       `json:"duration"`
	MaxGroupSize    int       `json:"maxGroupSize"`
	Difficulty      string    `json:"difficulty"`
	RatingsAverage  int       `json:"ratingsAverage"`
	RatingsQuantity int       `json:"ratingsQuantity"`
	Price           int       `json:"price"`
	PriceDiscount   int       `json:"priceDiscount"`
	Summary         string    `json:"summary"`
	Description     string    `json:"description"`
	ImageCover      string    `json:"imageCover"`
	Images          []string  `json:"images"`
	StartDates      time.Time `json:"startDates"`
	SecretTour      string    `json:"secretTour"`
}
