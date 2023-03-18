package nongow

// Sets values in data, up to and including i,
// to one value, all the remainder to another.
func scanset(data []int, i int) {

	if i >= len(data) {
		i = len(data) - 1
	} else {
		if i < -1 {
			i = -1
		}
	}

	for j := 0; j <= i; j++ {
		data[j] = 1
	}
	for j := i + 1; j < len(data); j++ {
		data[j] = 2
	}
}

func RunScanset(run bool) {
	if !run {
		return
	}

	count := 5
	for i := -2; i < count+2; i++ {

		data := make([]int, count)

		scanset(data, i)

		print(i, " ")
		for _, aa := range data {
			print(aa)
		}
		println()
	}
}
