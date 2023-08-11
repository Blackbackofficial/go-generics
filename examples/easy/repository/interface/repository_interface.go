package _interface

type Entity interface{}

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

type Repository struct {
	Entities map[int]Entity
}

func NewRepository() *Repository {
	return &Repository{
		Entities: make(map[int]Entity),
	}
}

func (r *Repository) Save(id int, entity Entity) {
	r.Entities[id] = entity
}
