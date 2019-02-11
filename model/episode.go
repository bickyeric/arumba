package model

import "time"

type Episode struct {
	ID        int
	No        float32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ComicID   int
}
