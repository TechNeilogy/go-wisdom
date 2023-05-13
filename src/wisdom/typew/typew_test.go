package typew

import "fmt"

func ExampleConvertThenAssert() {

	fmt.Println(ConvertThenAssert[int]("thirty-three"))
	fmt.Println(ConvertThenAssert[int](33))

	fmt.Println(ConvertThenAssert[string]("thirty-three"))
	fmt.Println(ConvertThenAssert[string](33))

	// output: 0 false
	// 33 true
	// thirty-three true
	//  false
}
