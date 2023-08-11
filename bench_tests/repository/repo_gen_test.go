package repository

import (
	"generics/examples/easy/repository/generics"
	"testing"
)

// go: go test -bench=.
func BenchmarkProductGenericsSave(b *testing.B) {
	repo := generics.NewRepository[generics.Product]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.Save(i, generics.Product{ID: i, Name: "Product", Price: 9.99})
	}
}

func BenchmarkUserGenericsSave(b *testing.B) {
	repo := generics.NewRepository[generics.User]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		repo.Save(i, generics.User{ID: i, Name: "User", Age: 30})
	}
}
