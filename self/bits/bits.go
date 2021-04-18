package bits

// set with or
// b | m = x
// 1 | 0 = 1
// 1 | 1 = 1
// 0 | 1 = 1
// 0 | 0 = 0
func SetByBit(x, mask uint32) uint32 {
	return x | mask
}

// get with and
// b & m = x
// 1 & 0 = 0
// 1 & 1 = 1
// 0 & 1 = 0
// 0 & 0 = 0
func GetByBit(x, mask uint32) uint32 {

	return x & mask
}

// not with
// ~x
// 11101001
// 00010110
func Not(x uint32) uint32 {
	return x ^ 0xffffffff
}

// if x == y will return 0

// use this to dedup

func DeDup(x uint32, y uint32) uint32 {
	return x ^ y
}

// 获取最后一位1
// 11110110
// get
// 00000010
// n & -n
// -n equal ~n + 1
func GetLastOne(x uint32) uint32 {
	//return x & (- x)
	return x & (Not(x) + 1)
}

func DelLastOne1(x uint32) uint32 {
	return x & (x - 1)
}
