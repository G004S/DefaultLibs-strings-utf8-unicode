package mystrings

import "testing"

func TestIndex(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		substr string
		want   int
	}{ // empty strings
		{name: "empty substr", s: "abc", substr: "", want: 0},
		{name: "both empty", s: "", substr: "", want: 0},
		{name: "empty s", s: "", substr: "a", want: -1},
		// different length
		{name: "substr longer than s", s: "ab", substr: "abc", want: -1},
		{name: "substr one byte longer than s", s: "abc", substr: "abcd", want: -1},
		{name: "equal strings", s: "abc", substr: "abc", want: 0},
		{name: "single char s and substr", s: "a", substr: "a", want: 0},
		// dif positions
		{name: "at start", s: "chicken", substr: "chi", want: 0},
		{name: "in middle", s: "chicken", substr: "ick", want: 2},
		{name: "at end", s: "chicken", substr: "ken", want: 4},
		{name: "last byte only", s: "chicken", substr: "n", want: 6},
		{name: "not found", s: "chicken", substr: "dmr", want: -1},
		// dif cases
		{name: "case sensitive", s: "Go", substr: "go", want: -1},
		// unicode
		{name: "cyrillic, byte index", s: "жук", substr: "к", want: 4},
		{name: "cyrillic substr at start", s: "жук", substr: "ж", want: 0},
		{name: "emoji, byte index", s: "a😀b", substr: "b", want: 5},
		{name: "multi-byte substr", s: "привет мир", substr: "мир", want: 13},
		// invalid UTF8
		{name: "invalid byte found", s: "a\xffb", substr: "\xff", want: 1},
		{name: "truncated rune as substr", s: "ж", substr: "\xd0", want: 0}}

	for _, tt := range testCases {
		got := Index(tt.s, tt.substr)
		if got != tt.want {
			t.Errorf("%v   Index()   s:%v substr:%v   got:%v   want:%v", tt.name, tt.s, tt.substr, got, tt.want)
		}
	}
}

func TestHasPrefix(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		prefix string
		want   bool
	}{ // empty strings
		{name: "empty prefix", s: "abc", prefix: "", want: true},
		{name: "both empty", s: "", prefix: "", want: true},
		{name: "empty s", s: "", prefix: "a", want: false},
		// lengths
		{name: "prefix longer than s", s: "go", prefix: "golang", want: false},
		{name: "prefix one byte longer than s", s: "go", prefix: "gol", want: false},
		{name: "equal strings", s: "go", prefix: "go", want: true},
		{name: "single char equal", s: "a", prefix: "a", want: true},
		// default
		{name: "real prefix", s: "golang", prefix: "go", want: true},
		{name: "one byte prefix", s: "golang", prefix: "g", want: true},
		{name: "one byte shorter than s", s: "golang", prefix: "golan", want: true},
		{name: "not a prefix", s: "golang", prefix: "lang", want: false},
		{name: "suffix is not prefix", s: "golang", prefix: "ng", want: false},
		{name: "differs in first byte", s: "golang", prefix: "Xo", want: false},
		{name: "differs in last byte", s: "golang", prefix: "gola!", want: false},
		// case
		{name: "case sensitive", s: "Golang", prefix: "go", want: false},
		// unicode
		{name: "cyrillic prefix", s: "жук", prefix: "ж", want: true},
		{name: "cyrillic not a prefix", s: "жук", prefix: "у", want: false},
		{name: "emoji prefix", s: "😀a", prefix: "😀", want: true},
		// invalid utf8
		{name: "first byte of rune is a prefix", s: "жук", prefix: "\xd0", want: true},
		{name: "second byte of rune is not a prefix", s: "жук", prefix: "\xb6", want: false},
		{name: "invalid byte prefix", s: "\xffabc", prefix: "\xff", want: true},
	}

	for _, tt := range testCases {
		got := HasPrefix(tt.s, tt.prefix)
		if got != tt.want {
			t.Errorf("%v   Index()   s:%v prefix:%v   got:%v   want:%v", tt.name, tt.s, tt.prefix, got, tt.want)
		}
	}
}

func TestCut(t *testing.T) {
	testCases := []struct {
		name       string
		s          string
		sep        string
		wantBefore string
		wantAfter  string
		wantBool   bool
	}{ // empty strings
		{name: "empty sep", s: "abc", sep: "", wantBefore: "", wantAfter: "abc", wantBool: true},
		{name: "both empty", s: "", sep: "", wantBefore: "", wantAfter: "", wantBool: true},
		{name: "empty s", s: "", sep: "=", wantBefore: "", wantAfter: "", wantBool: false},
		// lengths
		{name: "sep longer than s", s: "a", sep: "ab", wantBefore: "a", wantAfter: "", wantBool: false},
		{name: "sep one byte longer than s", s: "ab", sep: "abc", wantBefore: "ab", wantAfter: "", wantBool: false},
		{name: "equal strings", s: "=", sep: "=", wantBefore: "", wantAfter: "", wantBool: true},
		{name: "s equals multi-byte sep", s: "::", sep: "::", wantBefore: "", wantAfter: "", wantBool: true},
		// sep position
		{name: "key=value", s: "key=value", sep: "=", wantBefore: "key", wantAfter: "value", wantBool: true},
		{name: "sep at start", s: "=value", sep: "=", wantBefore: "", wantAfter: "value", wantBool: true},
		{name: "sep at end", s: "key=", sep: "=", wantBefore: "key", wantAfter: "", wantBool: true},
		{name: "not found", s: "novalue", sep: "=", wantBefore: "novalue", wantAfter: "", wantBool: false},
		// multiple presense
		{name: "first of several", s: "a=b=c", sep: "=", wantBefore: "a", wantAfter: "b=c", wantBool: true},
		{name: "adjacent seps", s: "a==b", sep: "=", wantBefore: "a", wantAfter: "=b", wantBool: true},
		// multibyte separator
		{name: "multi-char sep", s: "a::b", sep: "::", wantBefore: "a", wantAfter: "b", wantBool: true},
		{name: "partial match then real", s: "aaab", sep: "aab", wantBefore: "a", wantAfter: "", wantBool: true},
		{name: "almost match, not found", s: "abcabd", sep: "abe", wantBefore: "abcabd", wantAfter: "", wantBool: false},
		// case
		{name: "case sensitive", s: "Key=V", sep: "key", wantBefore: "Key=V", wantAfter: "", wantBool: false},
		// unicode
		{name: "cyrillic", s: "ключ=значение", sep: "=", wantBefore: "ключ", wantAfter: "значение", wantBool: true},
		{name: "cyrillic sep", s: "ключ→значение", sep: "→", wantBefore: "ключ", wantAfter: "значение", wantBool: true},
		{name: "emoji sep", s: "a😀b", sep: "😀", wantBefore: "a", wantAfter: "b", wantBool: true},
		//  invalid utf8
		{name: "invalid byte as sep", s: "a\xffb", sep: "\xff", wantBefore: "a", wantAfter: "b", wantBool: true},
		{name: "first byte of rune as sep", s: "ж", sep: "\xd0", wantBefore: "", wantAfter: "\xb6", wantBool: true}}

	for _, tt := range testCases {
		gotBefore, gotAfter, gotBool := Cut(tt.s, tt.sep)
		if gotBefore != tt.wantBefore || gotAfter != tt.wantAfter || gotBool != tt.wantBool {
			t.Errorf("%v   Index()   s:%v sep:%v   got:%v; %v; %v   want:%v; %v; %v", tt.name, tt.s, tt.sep, gotBefore, gotAfter, gotBool, tt.wantBefore, tt.wantAfter, tt.wantBool)
		}
	}
}

func TestCount(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		substr string
		want   int
	}{ // empty substr
		{name: "empty substr", s: "five", substr: "", want: 5},
		{name: "both empty", s: "", substr: "", want: 1},
		{name: "empty substr, cyrillic", s: "жук", substr: "", want: 4},
		{name: "empty substr, emoji", s: "a😀", substr: "", want: 3},
		{name: "empty substr, invalid byte", s: "a\xffb", substr: "", want: 4},
		// lengths
		{name: "empty s", s: "", substr: "a", want: 0},
		{name: "substr longer than s", s: "a", substr: "ab", want: 0},
		{name: "substr one byte longer than s", s: "ab", substr: "abc", want: 0},
		{name: "equal strings", s: "abc", substr: "abc", want: 1},
		{name: "single char equal", s: "a", substr: "a", want: 1},
		// default
		{name: "single char repeated", s: "cheese", substr: "e", want: 3},
		{name: "not found", s: "abc", substr: "x", want: 0},
		{name: "at start", s: "chicken", substr: "chi", want: 1},
		{name: "at end", s: "chicken", substr: "ken", want: 1},
		{name: "last byte only", s: "chicken", substr: "n", want: 1},
		{name: "at start and end", s: "abxab", substr: "ab", want: 2},
		{name: "adjacent", s: "ababab", substr: "ab", want: 3},
		// overlaps
		{name: "overlapping, even", s: "aaaa", substr: "aa", want: 2},
		{name: "overlapping, odd", s: "aaa", substr: "aa", want: 1},
		{name: "overlapping, five", s: "aaaaa", substr: "aa", want: 2},
		{name: "overlapping, aba", s: "ababa", substr: "aba", want: 1},
		// partial match
		{name: "partial match then real", s: "aaab", substr: "aab", want: 1},
		{name: "almost match, not found", s: "abcabd", substr: "abe", want: 0},
		// dif cases
		{name: "case sensitive", s: "Go go", substr: "go", want: 1},
		// unicode
		{name: "cyrillic char", s: "жужу", substr: "ж", want: 2},
		{name: "cyrillic pair", s: "жужу", substr: "жу", want: 2},
		{name: "emoji", s: "😀😀", substr: "😀", want: 2},

		// non-Utf8
		{name: "invalid byte repeated", s: "\xff\xff", substr: "\xff", want: 2},
		{name: "first byte of rune", s: "жж", substr: "\xd0", want: 2}}

	for _, tt := range testCases {
		got := Count(tt.s, tt.substr)
		if got != tt.want {
			t.Errorf("%v   Count()   s:%v   substr:%v   got:%v   want:%v", tt.name, tt.s, tt.substr, got, tt.want)
		}
	}
}

func TestIndexRune(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		r    rune
		want int
	}{
		{name: "rune at the beginning", s: "hello", r: 'h', want: 0},
		{name: "rune in the middle", s: "hello", r: 'l', want: 2},
		{name: "rune at the end", s: "hello", r: 'o', want: 4},
		{name: "rune not found", s: "hello", r: 'x', want: -1},
		{name: "empty string", s: "", r: 'a', want: -1},
		{name: "unicode rune", s: "hello мир", r: 'м', want: 6},
		{name: "first occurrence", s: "banana", r: 'a', want: 1},
	}
	for _, tt := range testCases {
		got := IndexRune(tt.s, tt.r)
		if got != tt.want {
			t.Errorf("%v   IndexRune()   s:%v   r:%v   got:%v   want:%v", tt.name, tt.s, tt.r, got, tt.want)
		}
	}
}

func TestTrimSpace(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{name: "no spaces", s: "hello", want: "hello"},
		{name: "spaces on both sides", s: "  hello  ", want: "hello"},
		{name: "spaces at the beginning", s: "  hello", want: "hello"},
		{name: "spaces at the end", s: "hello  ", want: "hello"},
		{name: "only spaces", s: "     ", want: ""},
		{name: "empty string", s: "", want: ""},
		{name: "one character", s: " a ", want: "a"},
		{name: "internal spaces", s: "hello world", want: "hello world"},
		{name: "tabs", s: "\t\thello\t", want: "hello"},
		{name: "newlines", s: "\nhello\n", want: "hello"},
	}

	for _, tt := range testCases {
		got := TrimSpace(tt.s)
		if got != tt.want {
			t.Errorf("%v   TrimSpace()  s:%v   got:%v   want:%v", tt.name, tt.s, got, tt.want)
		}
	}
}

func TestToLower(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{name: "lowercase", s: "hello", want: "hello"},
		{name: "uppercase", s: "HELLO", want: "hello"},
		{name: "mixed case", s: "HeLLo", want: "hello"},
		{name: "numbers", s: "Hello123", want: "hello123"},
		{name: "spaces", s: " Hello World ", want: " hello world "},
		{name: "empty string", s: "", want: ""},
		{name: "special characters", s: "!@#$%^&*", want: "!@#$%^&*"},
	}
	for _, tt := range testCases {
		got := ToLower(tt.s)
		if got != tt.want {
			t.Errorf("%v   ToLower()   s:%v   got:%v   want:%v", tt.name, tt.s, got, tt.want)
		}
	}
}

func TestRepeat(t *testing.T) {
	testCases := []struct {
		name  string
		s     string
		count int
		wantS string
		wantB bool
	}{
		{name: "repeat multiple times", s: "go", count: 3, wantS: "gogogo", wantB: true},
		{name: "repeat once", s: "go", count: 1, wantS: "go", wantB: true},
		{name: "repeat zero times", s: "go", count: 0, wantS: "", wantB: true},
		{name: "negative count", s: "go", count: -1, wantS: "", wantB: false},
		{name: "empty string", s: "", count: 5, wantS: "", wantB: true},
		{name: "single character", s: "a", count: 5, wantS: "aaaaa", wantB: true},
	}

	for _, tt := range testCases {
		gotS, gotB := Repeat(tt.s, tt.count)
		if gotS != tt.wantS || gotB != tt.wantB {
			t.Errorf("%v   Repeat()   s:%v   count:%v   got:%v; %v;   want:%v; %v", tt.name, tt.s, tt.count, gotS, gotB, tt.wantS, tt.wantB)
		}
	}
}
