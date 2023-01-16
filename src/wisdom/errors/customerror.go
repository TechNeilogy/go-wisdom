package customerror

import (
	"fmt"
	"github.com/TechNeilogy/go-wisdom/src/util"
)

type DivZero struct {
	v0 int
	v1 int
}

func (e DivZero) Error() string {
	return fmt.Sprintf("division by zero: %v / %v", e.v0, e.v1)
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, DivZero{a, b}
	}
	return a / b, nil
}

func Run() {

	util.PrintHeader("Custom Error Handling")

	a := 10
	for _, b := range []int{3, 0} {
		res, err := div(10, b)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("No Error: %v / %v = %v\n", a, b, res)
		}
	}
}
