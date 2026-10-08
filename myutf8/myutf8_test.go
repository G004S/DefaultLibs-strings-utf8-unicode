package myutf8

import (
	"bytes"
	"testing"
)

func TestValidRune(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // start of the range
		{name: "null rune", r: 0, want: true},
		{name: "one", r: 1, want: true},

		// normal symbols
		{name: "latin letter", r: 'a', want: true},
		{name: "space", r: ' ', want: true},
		{name: "cyrillic", r: 'ж', want: true},
		{name: "japanese hiragana", r: 'あ', want: true},
		{name: "emoji (valid, one rune)", r: '😀', want: true},

		// границы ASCII
		{name: "DEL", r: 127, want: true},
		{name: "first non-ASCII", r: 128, want: true},

		// border before surrogate zone
		{name: "just before surrogates (0xD7FF)", r: 0xD7FF, want: true},

		// surrogate zone 0xD800-0xDFFF
		{name: "first surrogate (0xD800)", r: 0xD800, want: false},
		{name: "high surrogate end (0xDBFF)", r: 0xDBFF, want: false},
		{name: "low surrogate start (0xDC00)", r: 0xDC00, want: false},
		{name: "middle of surrogates (0xDA00)", r: 0xDA00, want: false},
		{name: "last surrogate (0xDFFF)", r: 0xDFFF, want: false},

		// borderline after surrogate reange
		{name: "just after surrogates (0xE000)", r: 0xE000, want: true},

		// upper border of Unicode
		{name: "max rune (0x10FFFF)", r: 0x10FFFF, want: true},
		{name: "just above max (0x110000)", r: 0x110000, want: false},

		// incorrect values
		{name: "negative one", r: -1, want: false},
		{name: "min int32", r: -2147483648, want: false},
		{name: "max int32", r: 2147483647, want: false}}

	for _, tt := range testCases {
		gotB := ValidRune(tt.r)
		if gotB != tt.want {
			t.Errorf("%v   ValidRune()   r:%v   got:%v   want:%v", tt.name, tt.r, gotB, tt.want)
		}
	}
}

func TestRuneLen(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want int
	}{ // 1byte
		{name: "null rune", r: 0, want: 1},
		{name: "latin letter", r: 'a', want: 1},
		{name: "space", r: ' ', want: 1},
		{name: "newline", r: '\n', want: 1},
		{name: "DEL (last 1-byte)", r: 127, want: 1},

		// 2byte
		{name: "first 2-byte (0x80)", r: 128, want: 2},
		{name: "inverted exclamation", r: '¡', want: 2},
		{name: "latin with diacritic", r: 'é', want: 2},
		{name: "cyrillic", r: 'ж', want: 2},
		{name: "last 2-byte (0x7FF)", r: 2047, want: 2},

		// 3byte
		{name: "first 3-byte (0x800)", r: 2048, want: 3},
		{name: "euro sign", r: '€', want: 3},
		{name: "japanese hiragana", r: 'あ', want: 3},
		{name: "just before surrogates (0xD7FF)", r: 0xD7FF, want: 3},
		{name: "just after surrogates (0xE000)", r: 0xE000, want: 3},
		{name: "replacement char (0xFFFD)", r: 0xFFFD, want: 3},
		{name: "last 3-byte (0xFFFF)", r: 65535, want: 3},

		//  4byte
		{name: "first 4-byte (0x10000)", r: 65536, want: 4},
		{name: "emoji", r: '😀', want: 4},
		{name: "max rune (0x10FFFF)", r: 1114111, want: 4},

		// surrogates
		{name: "first surrogate (0xD800)", r: 0xD800, want: -1},
		{name: "middle of surrogates (0xDA00)", r: 0xDA00, want: -1},
		{name: "last surrogate (0xDFFF)", r: 0xDFFF, want: -1},

		// non Unicode
		{name: "just above max (0x110000)", r: 0x110000, want: -1},
		{name: "max int32", r: 2147483647, want: -1},
		{name: "negative one", r: -1, want: -1},
		{name: "min int32", r: -2147483648, want: -1}}

	for _, tt := range testCases {
		got := RuneLen(tt.r)
		if got != tt.want {
			t.Errorf("%v   RuneLen()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}
func TestEncodeRune(t *testing.T) {
	testCases := []struct {
		name  string
		p     []byte
		r     rune
		wantI int
		wantB bool
		wantP []byte
	}{
		// 1 byte
		{name: "latin a", p: make([]byte, 1), r: 'a', wantI: 1, wantB: true, wantP: []byte{0x61}},
		{name: "null rune", p: []byte{0xFF}, r: 0, wantI: 1, wantB: true, wantP: []byte{0x00}},
		{name: "DEL (last 1-byte)", p: make([]byte, 1), r: 127, wantI: 1, wantB: true, wantP: []byte{0x7F}},

		// 2 bytes
		{name: "first 2-byte (0x80)", p: make([]byte, 2), r: 0x80, wantI: 2, wantB: true, wantP: []byte{0xC2, 0x80}},
		{name: "latin e acute", p: make([]byte, 2), r: 'é', wantI: 2, wantB: true, wantP: []byte{0xC3, 0xA9}},
		{name: "cyrillic zhe", p: make([]byte, 2), r: 'ж', wantI: 2, wantB: true, wantP: []byte{0xD0, 0xB6}},
		{name: "last 2-byte (0x7FF)", p: make([]byte, 2), r: 0x7FF, wantI: 2, wantB: true, wantP: []byte{0xDF, 0xBF}},

		// 3 bytes
		{name: "first 3-byte (0x800)", p: make([]byte, 3), r: 0x800, wantI: 3, wantB: true, wantP: []byte{0xE0, 0xA0, 0x80}},
		{name: "euro sign", p: make([]byte, 3), r: '€', wantI: 3, wantB: true, wantP: []byte{0xE2, 0x82, 0xAC}},
		{name: "hiragana a", p: make([]byte, 3), r: 'あ', wantI: 3, wantB: true, wantP: []byte{0xE3, 0x81, 0x82}},
		{name: "just before surrogates (0xD7FF)", p: make([]byte, 3), r: 0xD7FF, wantI: 3, wantB: true, wantP: []byte{0xED, 0x9F, 0xBF}},
		{name: "just after surrogates (0xE000)", p: make([]byte, 3), r: 0xE000, wantI: 3, wantB: true, wantP: []byte{0xEE, 0x80, 0x80}},
		{name: "replacement char (0xFFFD)", p: make([]byte, 3), r: 0xFFFD, wantI: 3, wantB: true, wantP: []byte{0xEF, 0xBF, 0xBD}},
		{name: "last 3-byte (0xFFFF)", p: make([]byte, 3), r: 0xFFFF, wantI: 3, wantB: true, wantP: []byte{0xEF, 0xBF, 0xBF}},

		// 4 bytes
		{name: "first 4-byte (0x10000)", p: make([]byte, 4), r: 0x10000, wantI: 4, wantB: true, wantP: []byte{0xF0, 0x90, 0x80, 0x80}},
		{name: "emoji", p: make([]byte, 4), r: '😀', wantI: 4, wantB: true, wantP: []byte{0xF0, 0x9F, 0x98, 0x80}},
		{name: "max rune (0x10FFFF)", p: make([]byte, 4), r: 0x10FFFF, wantI: 4, wantB: true, wantP: []byte{0xF4, 0x8F, 0xBF, 0xBF}},

		// buffer larger than needed: extra bytes stay untouched
		{name: "ascii in larger buffer", p: []byte{0xFF, 0xFF, 0xFF}, r: 'a', wantI: 1, wantB: true, wantP: []byte{0x61, 0xFF, 0xFF}},
		{name: "2-byte in larger buffer", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: 'ж', wantI: 2, wantB: true, wantP: []byte{0xD0, 0xB6, 0xFF, 0xFF}},
		{name: "3-byte in larger buffer", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: '€', wantI: 3, wantB: true, wantP: []byte{0xE2, 0x82, 0xAC, 0xFF}},

		// buffer too short: 0, false, buffer unchanged
		{name: "nil buffer", p: nil, r: 'a', wantI: 0, wantB: false, wantP: nil},
		{name: "empty buffer", p: []byte{}, r: 'a', wantI: 0, wantB: false, wantP: []byte{}},
		{name: "2-byte rune, 1-byte buffer", p: []byte{0xFF}, r: 'ж', wantI: 0, wantB: false, wantP: []byte{0xFF}},
		{name: "3-byte rune, 2-byte buffer", p: []byte{0xFF, 0xFF}, r: '€', wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF}},
		{name: "4-byte rune, 3-byte buffer", p: []byte{0xFF, 0xFF, 0xFF}, r: '😀', wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF}},

		// invalid runes: 0, false, buffer unchanged
		{name: "first surrogate (0xD800)", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: 0xD800, wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "last surrogate (0xDFFF)", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: 0xDFFF, wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "just above max (0x110000)", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: 0x110000, wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "negative one", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: -1, wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "max int32", p: []byte{0xFF, 0xFF, 0xFF, 0xFF}, r: 2147483647, wantI: 0, wantB: false, wantP: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "invalid rune, empty buffer", p: []byte{}, r: 0xD800, wantI: 0, wantB: false, wantP: []byte{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotI, gotB := EncodeRune(tt.p, tt.r)
			if gotI != tt.wantI || gotB != tt.wantB || !bytes.Equal(tt.p, tt.wantP) {
				t.Errorf("%v   EncodeRune()   r:%v   got:%v, %v, p:% x   want:%v, %v, p:% x",
					tt.name, tt.r, gotI, gotB, tt.p, tt.wantI, tt.wantB, tt.wantP)
			}
		})
	}
}

func TestAppendRune(t *testing.T) {
	testCases := []struct {
		name string
		p    []byte
		r    rune
		want []byte
	}{ // addition to nil slice and to empty slice
		{name: "nil + latin a", p: nil, r: 'a', want: []byte{0x61}},
		{name: "empty + latin a", p: []byte{}, r: 'a', want: []byte{0x61}},
		{name: "empty + cyrillic zhe", p: []byte{}, r: 'ж', want: []byte{0xD0, 0xB6}},
		{name: "empty + euro sign", p: []byte{}, r: '€', want: []byte{0xE2, 0x82, 0xAC}},
		{name: "empty + emoji", p: []byte{}, r: '😀', want: []byte{0xF0, 0x9F, 0x98, 0x80}},

		// length borders
		{name: "null rune", p: []byte{}, r: 0, want: []byte{0x00}},
		{name: "DEL (last 1-byte)", p: []byte{}, r: 127, want: []byte{0x7F}},
		{name: "first 2-byte (0x80)", p: []byte{}, r: 0x80, want: []byte{0xC2, 0x80}},
		{name: "last 2-byte (0x7FF)", p: []byte{}, r: 0x7FF, want: []byte{0xDF, 0xBF}},
		{name: "first 3-byte (0x800)", p: []byte{}, r: 0x800, want: []byte{0xE0, 0xA0, 0x80}},
		{name: "just before surrogates (0xD7FF)", p: []byte{}, r: 0xD7FF, want: []byte{0xED, 0x9F, 0xBF}},
		{name: "just after surrogates (0xE000)", p: []byte{}, r: 0xE000, want: []byte{0xEE, 0x80, 0x80}},
		{name: "last 3-byte (0xFFFF)", p: []byte{}, r: 0xFFFF, want: []byte{0xEF, 0xBF, 0xBF}},
		{name: "first 4-byte (0x10000)", p: []byte{}, r: 0x10000, want: []byte{0xF0, 0x90, 0x80, 0x80}},
		{name: "max rune (0x10FFFF)", p: []byte{}, r: 0x10FFFF, want: []byte{0xF4, 0x8F, 0xBF, 0xBF}},

		// addition to existing data
		{name: "ascii after ascii", p: []byte("hi"), r: '!', want: []byte("hi!")},
		{name: "cyrillic after ascii", p: []byte("hi"), r: 'ж', want: []byte{'h', 'i', 0xD0, 0xB6}},
		{name: "emoji after cyrillic", p: []byte("ж"), r: '😀', want: []byte{0xD0, 0xB6, 0xF0, 0x9F, 0x98, 0x80}},
		{name: "null rune after data", p: []byte("a"), r: 0, want: []byte{'a', 0x00}},

		// slice with spare capacity
		{name: "spare capacity", p: append(make([]byte, 0, 10), 'a'), r: '€', want: []byte{'a', 0xE2, 0x82, 0xAC}},

		// invalid rune. Nothing gets added
		{name: "first surrogate (0xD800), empty", p: []byte{}, r: 0xD800, want: []byte{}},
		{name: "last surrogate (0xDFFF), empty", p: []byte{}, r: 0xDFFF, want: []byte{}},
		{name: "above max (0x110000), empty", p: []byte{}, r: 0x110000, want: []byte{}},
		{name: "negative one, empty", p: []byte{}, r: -1, want: []byte{}},
		{name: "invalid rune, nil", p: nil, r: 0xD800, want: nil},
		{name: "surrogate keeps existing data", p: []byte("hi"), r: 0xD800, want: []byte("hi")},
		{name: "negative keeps existing data", p: []byte("hi"), r: -1, want: []byte("hi")},
		{name: "max int32 keeps existing data", p: []byte("hi"), r: 2147483647, want: []byte("hi")},
	}

	for _, tt := range testCases {
		got := AppendRune(tt.p, tt.r)
		if string(got) != string(tt.want) {
			t.Errorf("%v   AppendRune()   p:%v   r:%v   got:%v,  want:%v", tt.name, tt.p, tt.r, got, tt.want)
		}
	}
}
