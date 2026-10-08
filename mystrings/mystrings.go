package mystrings

import (
	"goprimer/myunicode"
)

func Index(s, substr string) int {
	if substr == "" {
		return 0
	}
	for i := 0; i < len(s)-(len(substr)-1); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}

	return -1
}

func HasPrefix(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func Cut(s, sep string) (before, after string, found bool) {
	for i := 0; i < len(s)-(len(sep)-1); i++ {
		if s[i:i+len(sep)] == sep {
			return s[:i], s[i+len(sep):], true
		}
	}

	return s, "", false
}

func Count(s, substr string) int {
	var counter int
	if substr == "" {
		return len([]rune(s)) + 1
	}
	for i := 0; i < len(s)-(len(substr)-1); i++ {
		if s[i:i+len(substr)] == substr {
			counter++
			i += len(substr) - 1
		}
	}

	return counter
}

func IndexRune(s string, r rune) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}

	return -1
}

func TrimSpace(s string) string {
	var start int
	var finish = len(s)
	for i := 0; i < len(s); i++ {
		if myunicode.IsSpace(rune(s[i])) {
			start++
		} else {
			break
		}
	}

	if start == len(s) {
		return ""
	}

	for i := len(s) - 1; i > start; i-- {
		if myunicode.IsSpace(rune(s[i])) {
			finish--
		} else {
			break
		}
	}
	return s[start:finish]
}

func ToLower(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		if myunicode.IsUpper(rune(bs[i])) {
			bs[i] += 32
		}
	}
	return string(bs)
}

func Repeat(s string, count int) (string, bool) {
	if count < 0 {
		return "", false
	}
	var res string
	for i := 0; i < count; i++ {
		res += s
	}
	return res, true
}
