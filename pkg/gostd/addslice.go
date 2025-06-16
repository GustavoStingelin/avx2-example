package gostd

func AddSlice(dst, a, b []uint64) {
	for i := range a {
		dst[i] = a[i] + b[i]
	}
}
