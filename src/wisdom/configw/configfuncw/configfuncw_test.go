package configfuncw

import (
	"fmt"
)

func ExampleConfigBuilder_Build() {
	x1, y1 := Make(
		Name("1"),
		IsErr(true),
		Count(3),
	)

	x2, y2 := Make(
		Name("1"),
		Count(3),
	)

	fmt.Println(x1, y1)
	fmt.Println(x2, y2)

	// output: <nil> IsErr
	// &{1 false 3} <nil>
}
