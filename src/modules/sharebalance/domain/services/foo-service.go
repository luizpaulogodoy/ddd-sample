package services

import "github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"

// FooService -
type FooService interface {
	Create(name string) (foos entities.Foo, err error)
}
