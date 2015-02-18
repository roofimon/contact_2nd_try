package contact

type Provider interface {
	Get(id string) (Information, error)
	All() []Information
	Update(i Information) error
	Delete(id string) error
	Add(i *Information) error
}
