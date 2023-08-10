package repository

import (
	"generics/examples/repository/simple"
	"testing"
)

// go: go test -bench=.
func BenchmarkProductSave(b *testing.B) {
	productRepo := simple.NewProductRepository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productRepo.Save(simple.Product{ID: 1, Name: "iPhone", Price: 999.99})
	}
}

func BenchmarkUserSave(b *testing.B) {
	userRepo := simple.NewUserRepository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userRepo.Save(simple.User{ID: 1, Name: "Hasan", Age: 18})
	}
}
