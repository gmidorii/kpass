package gate

type Gater interface {
	Examine(c Condition, e string, r chan Result)
}

type Result struct {
	OK      bool
	Message string
}

type Condition struct {
	Elements []string
	Num      int
	Nums     []int
	FNum     float64
	FNums    []float64
	Str      string
	Strs     []string
}
