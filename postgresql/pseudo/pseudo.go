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

//指定固定长度及其输入，生成不超过指定长度的伪装后的输出
//32位长度最大为10,64位下输出最大为19。所以这里len<=19
func PseudoEncrypt64(val uint64, len uint64) uint64 {
	n := GetMaxN(len)
	var l1 uint64
	var l2 uint64
	var r1 uint64
	var r2 uint64
	var bhalfofn uint64 = (1 << (n / 2)) - 1 //if there is 22 bits mask for a half-number => 44 bits total
	l1 = (val >> (n / 2))
	r1 = val & bhalfofn
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ (uint64((((1366*r1+150889)%714025)/714025.00)*32767*32767))&bhalfofn
		l1 = l2
		r1 = r2
	}

	return (l1<<(n/2) + r1)

}

//输入指定长度输出其最大幂次方，并且不能为奇数。
func GetMaxN(val uint64) (n uint64) {
	bigInt := exponent(10, val) - 1
	n = uint64(math.Floor(math.Log2(float64(bigInt))))
	if n%2 != 0 {
		n -= 1
	}
	return
}

func exponent(a, n uint64) uint64 {
	result := uint64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}
