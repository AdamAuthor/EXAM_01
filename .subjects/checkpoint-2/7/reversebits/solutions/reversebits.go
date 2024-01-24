package solutions

func ReverseBits(oct byte) byte {
	r := 0
	for i := 0; i < 7; i++ {
		if oct%2 > 0 {
			r = r + 1
		}
		r *= 2
		oct /= 2
	}
	return byte(r)
}
