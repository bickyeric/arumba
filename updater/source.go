package updater

import "github.com/bickyeric/arumba/updater/source"

var Sources = []source.ISource{
	source.Komikcast{},
	source.Komikindo{},
	source.Mangacan{},
	source.Mangaku{},
	source.Mangatail{},
}
