package myunicode

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'

}

func IsLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func IsSpace(r rune) bool {
	return r == ' ' || (r >= '\t' && r <= '\r')
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsPunct(r rune) bool {
	//outside ASCII
	if r > 127 {
		return false
	}
	isAlpha := (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
	isNum := r >= '0' && r <= '9'
	isSpace := r == ' '
	return !isAlpha && !isNum && !isSpace
}

func IsControl(r rune) bool {
	return (r >= 0 && r <= 31) || r == 127
}

func ToUpper(r rune) rune {
	isLow := r >= 'a' && r <= 'z'
	if isLow {
		return r - 32
	}
	return r
}

func ToLower(r rune) rune {
	isUp := r >= 'A' && r <= 'Z'
	if isUp {
		return r + 32
	}
	return r
}
