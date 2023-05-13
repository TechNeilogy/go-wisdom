package algorithmw

import "fmt"

func ExampleScanSet() {

	count := 5
	for i := -2; i < count+2; i++ {

		data := make([]int, count)

		ScanSet(data, i, 1, 2)

		for _, aa := range data {
			fmt.Print(aa)
		}
		fmt.Println()
	}

	// output: 22222
	//22222
	//12222
	//11222
	//11122
	//11112
	//11111
	//11111
	//11111
}
