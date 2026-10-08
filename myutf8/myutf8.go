package myutf8

func ValidRune(r rune) bool {
	withinRuneRange := r >= 0 && r <= 1114111       //from MIN rune to MAX rune: 0x0000 - 0x10FFFF
	notWithinScalarValues := r < 55296 || r > 57343 //values between 0xD800 - 0xDFFF are reserved for 16bit PAIRS (for emojis etc)
	return withinRuneRange && notWithinScalarValues
}

func RuneLen(r rune) int {
	if !ValidRune(r) {
		return -1
	}
	ASCII := r >= 0 && r <= 127
	extendedAlpha := r >= 128 && r <= 2047
	basicMultilingualPlate := r >= 2048 && r <= 65535
	astralPlanes := r >= 65536 && r <= 1114111

	switch {
	case ASCII:
		return 1
	case extendedAlpha:
		return 2
	case basicMultilingualPlate:
		return 3
	case astralPlanes:
		return 4
	}
	return -1
}

func EncodeRune(p []byte, r rune) (int, bool) {
	if !ValidRune(r) {
		return 0, false
	}
	l := RuneLen(r)
	if l > len(p) {
		return 0, false
	}

	switch l {
	case 1:
		p[0] = byte(r)
	case 2:
		b0 := (r >> 6) | 0xC0
		b1 := (r & 0x3F) | 0x80
		p[0], p[1] = byte(b0), byte(b1)
	case 3:
		b0 := (r >> 12) | 0xE0
		b1 := ((r >> 6) & 0x3F) | 0x80
		b2 := (r & 0x3F) | 0x80
		p[0], p[1], p[2] = byte(b0), byte(b1), byte(b2)
	case 4:
		b0 := (r >> 18) | 0xF0
		b1 := ((r >> 12) & 0x3F) | 0x80
		b2 := ((r >> 6) & 0x3F) | 0x80
		b3 := (r & 0x3F) | 0x80
		p[0], p[1], p[2], p[3] = byte(b0), byte(b1), byte(b2), byte(b3)
	}

	return l, true
}

func AppendRune(p []byte, r rune) []byte {
	var buf [4]byte
	n, _ := EncodeRune(buf[:], r)
	return append(p, buf[:n]...)
}
