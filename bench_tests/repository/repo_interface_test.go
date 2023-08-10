package repository

import (
	"generics/examples/repository/interface"
	"testing"
)

// go: go test -bench=.
func BenchmarkProductInterfaceSave(b *testing.B) {
	repo := _interface.NewRepository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.Save(i, &_interface.Product{ID: 1, Name: "Product", Price: 9.99})
	}
}

func BenchmarkUserInterfaceSave(b *testing.B) {
	repo := _interface.NewRepository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.Save(i, &_interface.User{ID: 1, Name: "User", Age: 30})
	}
}
