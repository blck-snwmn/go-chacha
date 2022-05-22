package gochacha

func rotationN(n uint32, shift uint) uint32 {
	return n>>(32-shift) | n<<shift
}

func quarterRound(a, b, c, d uint32) (uint32, uint32, uint32, uint32) {
	a += b
	d ^= a
	d = rotationN(d, 16)

	c += d
	b ^= c
	b = rotationN(b, 12)

	a += b
	d ^= a
	d = rotationN(d, 8)

	c += d
	b ^= c
	b = rotationN(b, 7)
	return a, b, c, d
}
