package factory

import (
	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/repositories"
	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/services"
	implRepositories "github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/infrastructure/repositories"
	impServices "github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/infrastructure/services"
)

// ShareBalanceFactory -
type ShareBalanceFactory interface {
	NewFooService() services.FooService
	NewFooRepository() repositories.FooRepository
}

// ShareBalanceFactoryImpl -
type ShareBalanceFactoryImpl struct{}

// NewFooRepository -
func (s ShareBalanceFactoryImpl) NewFooRepository() repositories.FooRepository {
	repo, _ := implRepositories.NewFooRepository()

	return repo
}

// NewFooService -
func (s ShareBalanceFactoryImpl) NewFooService() services.FooService {
	return impServices.NewFooService(s.NewFooRepository())
}
