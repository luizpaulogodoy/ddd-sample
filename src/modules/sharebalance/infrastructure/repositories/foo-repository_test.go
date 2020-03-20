package repositories

import (
	"reflect"
	"testing"

	"github.com/luizpaulogodoy/ddd-sample/src/modules/sharebalance/domain/entities"
)

func TestFooRepository_FindAllFoos(t *testing.T) {
	tests := []struct {
		name       string
		a          *FooRepository
		wantAssets []entities.Foo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &FooRepository{}
			if gotAssets := a.FindAllFoos(); !reflect.DeepEqual(gotAssets, tt.wantAssets) {
				t.Errorf("FooRepository.FindAllFoos() = %v, want %v", gotAssets, tt.wantAssets)
			}
		})
	}
}
