package source

import (
	"github.com/bickyeric/arumba/updater"
)

// Mangatail ...
type Mangatail struct{}

var _ updater.ISource = (*Mangatail)(nil)

// Name ...
func (Mangatail) Name() string { return "mangatail" }

// GetID ...
func (Mangatail) GetID() int { return 2 }
