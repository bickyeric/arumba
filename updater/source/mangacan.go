package source

import (
	"github.com/bickyeric/arumba/updater"
)

type Mangacan struct{}

var _ updater.ISource = (*Mangacan)(nil)

func (Mangacan) Name() string { return "mangacan" }
func (Mangacan) GetID() int   { return 3 }
