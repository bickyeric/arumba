package source

import (
	"github.com/bickyeric/arumba/updater"
)

// Mangacan ...
type Mangacan struct{}

var _ updater.ISource = (*Mangacan)(nil)

// Name ...
func (Mangacan) Name() string { return "mangacan" }

// GetID ...
func (Mangacan) GetID() int { return 3 }
