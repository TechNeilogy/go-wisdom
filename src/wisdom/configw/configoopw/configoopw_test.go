package configoopw

import (
	"fmt"
)

func ExampleConfigBuilder_Build() {
	x0, y0 := New().
		Name("1").
		IsErr(true).
		Count(3).
		Build()

	x1, y1 := New().
		Name("1").
		Count(3).
		Build()

	fmt.Println(x0, y0)
	fmt.Println(x1, y1)

	// output: <nil> IsErr
	// &{1 false 3} <nil>
}
