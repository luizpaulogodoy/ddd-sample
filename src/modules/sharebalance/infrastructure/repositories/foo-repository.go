package repositories

import "github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"

// FooRepository is an implementation of the FooRepository interface using Postgres db as database
type FooRepository struct {
}

// NewFooRepository - Returns a new foo repository using the client factory to create a new client
func NewFooRepository() (*FooRepository, error) {
	// TODO Create db connection
	return &FooRepository{}, nil
}

// FindAllFoos -
func (a *FooRepository) FindAllFoos() (assets []entities.Foo) {
	return
}
