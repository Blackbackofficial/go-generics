package generics

import "cmp"

// UniqueElements returns a new slice containing the unique elements of a
func UniqueElements[T comparable](a []T) []T {
	uniq := make([]T, 0)
	mp := make(map[T]bool)
	for _, v := range a {
		if _, ok := mp[v]; !ok {
			mp[v] = true
			uniq = append(uniq, v)
		}
	}
	return uniq
}

func Equal[T cmp.Ordered](a, b T) bool {
	return a == b
}
