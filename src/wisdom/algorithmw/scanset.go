package algorithmw

// Sets values in data, up to and including index i,
// to lo, all the remainder hi.
func ScanSet[T any](data []T, i int, lo T, hi T) {

	if i >= len(data) {
		i = len(data) - 1
	} else {
		if i < -1 {
			i = -1
		}
	}

	for j := 0; j <= i; j++ {
		data[j] = lo
	}
	for j := i + 1; j < len(data); j++ {
		data[j] = hi
	}
}
