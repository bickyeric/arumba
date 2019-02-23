package source

import (
	"github.com/bickyeric/arumba/updater"
)

type Mangatail struct{}

var _ updater.ISource = (*Mangatail)(nil)

func (Mangatail) Name() string { return "mangatail" }
func (Mangatail) GetID() int   { return 2 }
