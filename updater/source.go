package updater

// ISource ...
type ISource interface {
	Name() string
	GetID() int
}
