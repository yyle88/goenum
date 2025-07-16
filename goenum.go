package goenum

import "github.com/yyle88/goenum/internal/constraint"

func Valid[E constraint.EnumType[E]](value E) bool {
	for _, elem := range value.Enums() {
		if elem == value {
			return true
		}
	}
	return false
}

func Check[E constraint.EnumItem[E]](value E) bool {
	return value == value.Basic() || Valid(value)
}
