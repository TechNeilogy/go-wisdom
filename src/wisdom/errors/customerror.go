package customerror

import "fmt"

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
	a := 10
	b := 3
	res, err := div(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v / %v = %v\n", a, b, res)
	}

	a = 10
	b = 3
	res, err = div(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v / %v = %v\n", a, b, res)
	}
}
