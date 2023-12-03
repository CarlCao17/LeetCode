package utils

type Comparable interface {
	Boolean | Integer | Float | Complex | ~string
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64
}

type UnsignedInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

type Number interface {
	Integer | Float | Complex
}

type Complex interface {
	~complex64 | ~complex128
}

type Float interface {
	~float32 | ~float64
}

type Boolean interface {
	~bool
}
