package model

import "time"

type Episode struct {
	ID        int
	No        float64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ComicID   int
}
