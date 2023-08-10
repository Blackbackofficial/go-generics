package main

import (
	"fmt"
	"generics/examples/repository/generics"
	_interface "generics/examples/repository/interface"
	"generics/examples/repository/simple"
)

func main() {
	//		---- generics ----
	products := generics.NewRepository[generics.Product]()
	users := generics.NewRepository[generics.User]()

	p := generics.Product{ID: 1, Name: "iPhone", Price: 999.99}
	u := generics.User{ID: 1, Name: "Hasan", Age: 18}
	fmt.Print(p, u)

	products.Save(p.ID, p)
	users.Save(u.ID, u)

	//		---- simple ----

	productRepo := simple.NewProductRepository()
	userRepo := simple.NewUserRepository()

	product := simple.Product{ID: 1, Name: "iPhone", Price: 999.99}
	productRepo.Save(product)

	user := simple.User{ID: 1, Name: "Hasan", Age: 18}
	userRepo.Save(user)

	//		---- interface ----

	repo := _interface.NewRepository()

	productI := &_interface.Product{ID: 1, Name: "Coffee", Price: 10.99}
	repo.Save(product.ID, productI)

	userI := &_interface.User{ID: 1, Name: "John", Age: 30}
	repo.Save(user.ID, userI)

	if savedProductI, ok := repo.Entities[product.ID].(*_interface.Product); ok {
		fmt.Printf("Saved product: %+v\n", savedProductI)
	}
	if savedUser, ok := repo.Entities[user.ID].(*_interface.User); ok {
		fmt.Printf("Saved user: %+v\n", savedUser)
	}
}
