package repositories

import "github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"

// FooRepository -
type FooRepository interface {
	FindAllFoos() (foos []entities.Foo)
}
