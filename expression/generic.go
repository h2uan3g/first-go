package expression

type Ordered interface {
	int | float64 | string
}

type Stack[T Ordered] struct {
	vals []T
}

func GenericExp() {

}
