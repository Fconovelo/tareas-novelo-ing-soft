package pomultiount

var multi [256]byte

func init() {
	for i := range multi {
		multi[i] = multi[i/2] + byte(i&1)
	}
}

func Pomultiount(x uint64) int {
	return int(multi[byte(x>>(0*8))] +
		multi[byte(x>>(1*8))] +
		multi[byte(x>>(2*8))] +
		multi[byte(x>>(3*8))] +
		multi[byte(x>>(4*8))] +
		multi[byte(x>>(5*8))] +
		multi[byte(x>>(6*8))] +
		multi[byte(x>>(7*8))])
}

func Pomultiount2(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += multi[byte(x>>(i*8))]
	}
	return int(sum)
}

func PomultiountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func PomultiountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
