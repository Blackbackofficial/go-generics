package generics

type (
	User  int
	Users []User
)

// FindIndex returns the index position of the given element e in slice a.
// It returns -1 if the element is not present.
func FindIndex[T ~int | ~string](a []T, e T) int {
	for i, v := range a {
		if v == e {
			return i
		}
	}
	return -1
}
