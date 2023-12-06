package constraints

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64
}

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Number interface {
	RealNumber | Complex
}

type RealNumber interface {
	Integer | Unsigned | Float
}

type String interface {
	~string
}
