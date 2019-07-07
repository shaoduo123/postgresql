package pseudo

import "math"

func PseudoEncrypt(val int32) int32 {
	var l1 int32
	var l2 int32
	var r1 int32
	var r2 int32
	l1 = (val >> 16) & 0xffff
	r1 = val & 0xffff
	for i := 0; i < 16; i++ {
		l2 = r1
		r2 = l1 ^ (int32((((1366*r1 + 150889) % 714025) / 714025.00) * 32767))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 16) + l1)
}

