package gate

import (
	"fmt"
	"strconv"
)

type Minimumer struct{}

func (m Minimumer) Examine(c Condition, e string, r chan Result) {
	defer func() { r <- Result{OK: true} }()

	n, err := strconv.Atoi(e)
	if err != nil {
		r <- Result{
			OK:      false,
			Message: fmt.Sprintf("not int number: %v: line(%v)", e, c.Elements),
		}
		return
	}
	if n < c.Num {
		r <- Result{
			OK:      false,
			Message: fmt.Sprintf("min number n:%v min:%v => line(%v)", n, c.Num, c.Elements),
		}
		return
	}
}
