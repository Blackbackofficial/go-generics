package main

import (
	"fmt"
	index "generics/examples/easy/find_index/generics"
	repo "generics/examples/easy/repository/generics"
	"generics/examples/easy/repository/interface"
	"generics/examples/easy/repository/simple"
	slice "generics/examples/easy/slice_bytes/generics"
	uniq "generics/examples/easy/uniq_element/generics"
)

func main() {
	//		---- generics ----
	products := repo.NewRepository[repo.Product]()
	users := repo.NewRepository[repo.User]()

	p := repo.Product{ID: 1, Name: "iPhone", Price: 999.99}
	u := repo.User{ID: 1, Name: "Hasan", Age: 18}
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

	//     ---- uniq ----

	fmt.Println(uniq.UniqueElements([]int{4, 3, 2, 3, 1, 5, 1}))
	fmt.Println(uniq.UniqueElements([]string{"apple", "banana", "apple", "potato"}))

	//     ---- slice bytes ----

	buf := slice.AppendStringOrBytes([]byte{}, "foo")
	buf = slice.AppendStringOrBytes(buf, []byte("bar"))

	fmt.Println(string(buf)) // prints: foobar

	//     ---- find index ----

	numbers := []int{5, 3, 2, 8, 7}
	fmt.Println(index.FindIndex(numbers, 8)) // prints: 3

	usr := index.Users{1, 2, 3}
	fmt.Println(index.FindIndex(usr, 2)) // prints: 1
}
