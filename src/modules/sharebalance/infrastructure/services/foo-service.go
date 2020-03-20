package services

import (
	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"
	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/repositories"
)

// FooService -
type FooService struct {
	// Repositories...
	repository repositories.FooRepository
}

// NewFooService - Receive repositories and returns a NewFooService
func NewFooService(fooRepository repositories.FooRepository) *FooService {
	return &FooService{
		fooRepository,
	}
}

// Create -
func (s *FooService) Create(name string) (foos entities.Foo, err error) {
	// TODO
	return
}
