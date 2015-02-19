package contact

type StubProvider struct{}

func NewStubProvider() *StubProvider {
	return &StubProvider{}
}

func (sp *StubProvider) Get(id string) (Information, error) {
	return Information{Id: "d9356b78-6b54-4391-80d0-af2c4949d973", Email: "first@email.com", Title: "First", Content: "First Content"}, nil
}

func (sp *StubProvider) All() []Information {
	return []Information{Information{Id: "d9356b78-6b54-4391-80d0-af2c4949d973", Email: "first@email.com", Title: "First", Content: "First Content"}}
}

func (sp *StubProvider) Update(i Information) error {
	return nil
}

func (sp *StubProvider) Delete(id string) error {
	return nil
}

func (sp *StubProvider) Add(i *Information) error {
	return nil
}
