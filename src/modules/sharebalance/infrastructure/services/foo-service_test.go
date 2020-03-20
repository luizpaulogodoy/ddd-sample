package services

import (
	"reflect"
	"testing"

	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"
	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/repositories"
)

func TestFooService_Create(t *testing.T) {
	type fields struct {
		repository repositories.FooRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFoos entities.Foo
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FooService{
				repository: tt.fields.repository,
			}
			gotFoos, err := s.Create(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FooService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFoos, tt.wantFoos) {
				t.Errorf("FooService.Create() = %v, want %v", gotFoos, tt.wantFoos)
			}
		})
	}
}
