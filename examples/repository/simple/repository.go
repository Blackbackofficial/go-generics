package simple

// -- Product --

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Products struct {
	products map[int]Product
}

func (pr *Products) Save(product Product) {
	pr.products[product.ID] = product
}

func NewProductRepository() *Products {
	return &Products{
		products: make(map[int]Product),
	}
}

// -- User --

type User struct {
	ID   int
	Name string
	Age  int
}

type Users struct {
	users map[int]User
}

func (ur *Users) Save(user User) {
	ur.users[user.ID] = user
}

func NewUserRepository() *Users {
	return &Users{
		users: make(map[int]User),
	}
}
