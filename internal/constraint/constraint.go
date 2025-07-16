package constraint

type EnumType[E comparable] interface {
	comparable
	Enums() []E
}

type EnumItem[E comparable] interface {
	comparable
	Basic() E
	Enums() []E
}
