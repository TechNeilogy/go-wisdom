package algorithmw

// Sets values in data, up to and including index i,
// to lo, all the remainder hi.
func ScanSet[T any](data []T, mid int, lo T, hi T) {

	if mid >= len(data) {
		mid = len(data) - 1
	} else {
		if mid < -1 {
			mid = -1
		}
	}

	i := 0
	for ; i <= mid; i++ {
		data[i] = lo
	}
	for ; i < len(data); i++ {
		data[i] = hi
	}
}
