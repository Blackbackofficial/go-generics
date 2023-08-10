package generics

type Product struct {
	ID    int
	Name  string
	Price float64
	//...
}

type User struct {
	ID   int
	Name string
	Age  int
	//...
}

type Repository[T any] struct {
	entities map[int]T
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		entities: make(map[int]T),
	}
}

func (r *Repository[T]) Save(id int, entity T) {
	r.entities[id] = entity
}
