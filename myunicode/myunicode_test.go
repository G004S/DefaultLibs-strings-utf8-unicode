package myunicode

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // nums
		{name: "zero", r: '0', want: true},
		{name: "five", r: '5', want: true},
		{name: "nine", r: '9', want: true},

		// borderline num values
		{name: "slash before 0", r: '/', want: false},
		{name: "colon after 9", r: ':', want: false},

		// letters
		{name: "latin lower", r: 'a', want: false},
		{name: "latin upper", r: 'Z', want: false},
		{name: "cyrillic", r: 'ж', want: false},

		// other symbols
		{name: "space", r: ' ', want: false},
		{name: "minus", r: '-', want: false},
		{name: "dot", r: '.', want: false},
		{name: "newline", r: '\n', want: false},

		// borderline ASCII
		{name: "null rune", r: 0, want: false},
		{name: "max ASCII", r: 127, want: false},

		// Unicode nums. Not ASCII
		{name: "arabic-indic digit", r: '٣', want: false},
		{name: "fullwidth digit", r: '５', want: false},
		{name: "emoji", r: '😀', want: false}}

	for _, tt := range testCases {
		got := IsDigit(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsDigit()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsLetter(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ //upper
		{name: "upper A", r: 'A', want: true},
		{name: "upper M", r: 'M', want: true},
		{name: "upper Z", r: 'Z', want: true},

		// lower
		{name: "lower a", r: 'a', want: true},
		{name: "lower m", r: 'm', want: true},
		{name: "lower z", r: 'z', want: true},

		// borderline of the range
		{name: "at sign before A", r: '@', want: false},
		{name: "bracket after Z", r: '[', want: false},
		{name: "backtick before a", r: '`', want: false},
		{name: "brace after z", r: '{', want: false},

		// between range symbols
		{name: "underscore", r: '_', want: false},
		{name: "backslash", r: '\\', want: false},

		// nums
		{name: "digit zero", r: '0', want: false},
		{name: "digit nine", r: '9', want: false},

		// other symbols
		{name: "space", r: ' ', want: false},
		{name: "minus", r: '-', want: false},
		{name: "newline", r: '\n', want: false},

		// borderline values
		{name: "null rune", r: 0, want: false},
		{name: "max ASCII", r: 127, want: false},

		// non ASCII
		{name: "cyrillic lower", r: 'ж', want: false},
		{name: "cyrillic upper", r: 'Ж', want: false},
		{name: "latin with diacritic", r: 'é', want: false},
		{name: "japanese hiragana", r: 'あ', want: false},
		{name: "emoji", r: '😀', want: false}}

	for _, tt := range testCases {
		got := IsLetter(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsLetter()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsSpace(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // normal space
		{name: "space", r: ' ', want: true},

		// other "space" symbols from ASCII
		{name: "tab", r: '\t', want: false},
		{name: "newline", r: '\n', want: false},
		{name: "carriage return", r: '\r', want: false},
		{name: "vertical tab", r: '\v', want: false},
		{name: "form feed", r: '\f', want: false},

		// neighbour values
		{name: "exclamation after space", r: '!', want: false},
		{name: "unit separator before space", r: 31, want: false},

		// symbols
		{name: "latin letter", r: 'a', want: false},
		{name: "digit", r: '0', want: false},
		{name: "underscore", r: '_', want: false},
		{name: "cyrillic", r: 'ж', want: false},

		// Unicode spaces
		{name: "non-breaking space", r: '\u00A0', want: false},
		{name: "em space", r: '\u2003', want: false},
		{name: "ideographic space", r: '\u3000', want: false},

		// borderline values
		{name: "null rune", r: 0, want: false},
		{name: "max ASCII", r: 127, want: false},
		{name: "emoji", r: '😀', want: false}}

	for _, tt := range testCases {
		got := IsSpace(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsSpace()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsUpper(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // upper
		{name: "upper A", r: 'A', want: true},
		{name: "upper M", r: 'M', want: true},
		{name: "upper Z", r: 'Z', want: true},

		// neighbour values
		{name: "at sign before A", r: '@', want: false},
		{name: "bracket after Z", r: '[', want: false},

		// lower
		{name: "lower a", r: 'a', want: false},
		{name: "lower z", r: 'z', want: false},

		// nums and stuff
		{name: "digit", r: '0', want: false},
		{name: "space", r: ' ', want: false},
		{name: "underscore", r: '_', want: false},
		{name: "newline", r: '\n', want: false},

		// borderline values
		{name: "null rune", r: 0, want: false},
		{name: "max ASCII", r: 127, want: false},

		//non-ASCII
		{name: "cyrillic upper", r: 'Ж', want: false},
		{name: "latin upper with diacritic", r: 'É', want: false},
		{name: "emoji", r: '😀', want: false}}

	for _, tt := range testCases {
		got := IsUpper(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsUpper()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsLower(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // lowers
		{name: "lower a", r: 'a', want: true},
		{name: "lower m", r: 'm', want: true},
		{name: "lower z", r: 'z', want: true},

		// neighbour values
		{name: "backtick before a", r: '`', want: false},
		{name: "brace after z", r: '{', want: false},

		// upper
		{name: "upper A", r: 'A', want: false},
		{name: "upper Z", r: 'Z', want: false},

		// nums and stuff
		{name: "digit", r: '0', want: false},
		{name: "space", r: ' ', want: false},
		{name: "underscore", r: '_', want: false},
		{name: "newline", r: '\n', want: false},

		// borderline
		{name: "null rune", r: 0, want: false},
		{name: "max ASCII", r: 127, want: false},

		// non-ASCII
		{name: "cyrillic lower", r: 'ж', want: false},
		{name: "latin lower with diacritic", r: 'é', want: false},
		{name: "emoji", r: '😀', want: false}}

	for _, tt := range testCases {
		got := IsLower(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsLower()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsPunct(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // puncts
		{name: "dot", r: '.', want: true},
		{name: "comma", r: ',', want: true},
		{name: "exclamation", r: '!', want: true},
		{name: "question", r: '?', want: true},
		{name: "minus", r: '-', want: true},
		{name: "underscore", r: '_', want: true},
		{name: "quote", r: '"', want: true},
		{name: "backslash", r: '\\', want: true},
		{name: "tilde", r: '~', want: true},

		// neighbours
		{name: "slash before 0", r: '/', want: true},
		{name: "colon after 9", r: ':', want: true},
		{name: "at sign before A", r: '@', want: true},
		{name: "bracket after Z", r: '[', want: true},
		{name: "backtick before a", r: '`', want: true},
		{name: "brace after z", r: '{', want: true},

		// nums
		{name: "digit zero", r: '0', want: false},
		{name: "digit five", r: '5', want: false},
		{name: "digit nine", r: '9', want: false},

		// letters
		{name: "upper A", r: 'A', want: false},
		{name: "upper Z", r: 'Z', want: false},
		{name: "lower a", r: 'a', want: false},
		{name: "lower z", r: 'z', want: false},

		// space
		{name: "space", r: ' ', want: false},

		// neighbours of space
		{name: "exclamation right after space", r: 33, want: true},
		{name: "unit separator right before space", r: 31, want: true},

		// controls
		{name: "null rune", r: 0, want: true},
		{name: "tab", r: '\t', want: true},
		{name: "newline", r: '\n', want: true},
		{name: "carriage return", r: '\r', want: true},
		{name: "DEL", r: 127, want: true},

		// ASCII borderline
		{name: "last ASCII", r: 127, want: true},
		{name: "first non-ASCII", r: 128, want: false},

		// non-ASCII
		{name: "inverted exclamation", r: '¡', want: false},
		{name: "em dash", r: '—', want: false},
		{name: "cyrillic", r: 'ж', want: false},
		{name: "emoji", r: '😀', want: false},

		// negative
		{name: "negative rune", r: -1, want: true}}

	for _, tt := range testCases {
		got := IsPunct(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsPunct()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestIsControl(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want bool
	}{ // borderline of the range
		{name: "null rune", r: 0, want: true},
		{name: "start of heading", r: 1, want: true},
		{name: "unit separator (last in range)", r: 31, want: true},

		// controls
		{name: "tab", r: '\t', want: true},
		{name: "newline", r: '\n', want: true},
		{name: "vertical tab", r: '\v', want: true},
		{name: "form feed", r: '\f', want: true},
		{name: "carriage return", r: '\r', want: true},
		{name: "escape", r: 27, want: true},

		// DEL
		{name: "DEL", r: 127, want: true},

		// neighbour values
		{name: "space right after range", r: 32, want: false},
		{name: "tilde right before DEL", r: 126, want: false},
		{name: "first non-ASCII right after DEL", r: 128, want: false},

		// normal symbols
		{name: "latin letter", r: 'a', want: false},
		{name: "upper letter", r: 'Z', want: false},
		{name: "digit", r: '0', want: false},
		{name: "exclamation", r: '!', want: false},
		{name: "underscore", r: '_', want: false},

		// non ASCII
		{name: "C1 control (U+0080)", r: '\u0080', want: false},
		{name: "next line (U+0085)", r: '\u0085', want: false},
		{name: "non-breaking space", r: '\u00A0', want: false},
		{name: "cyrillic", r: 'ж', want: false},
		{name: "emoji", r: '😀', want: false},

		// negative
		{name: "negative rune", r: -1, want: false}}

	for _, tt := range testCases {
		got := IsControl(tt.r)
		if got != tt.want {
			t.Errorf("%v   IsControl()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestToUpper(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want rune
	}{ // lower
		{name: "lower a", r: 'a', want: 'A'},
		{name: "lower m", r: 'm', want: 'M'},
		{name: "lower z", r: 'z', want: 'Z'},

		// upper
		{name: "upper A", r: 'A', want: 'A'},
		{name: "upper M", r: 'M', want: 'M'},
		{name: "upper Z", r: 'Z', want: 'Z'},

		// borderline values
		{name: "backtick before a", r: '`', want: '`'},
		{name: "brace after z", r: '{', want: '{'},
		{name: "at sign before A", r: '@', want: '@'},
		{name: "bracket after Z", r: '[', want: '['},

		// nums and stuff
		{name: "digit", r: '5', want: '5'},
		{name: "space", r: ' ', want: ' '},
		{name: "underscore", r: '_', want: '_'},
		{name: "newline", r: '\n', want: '\n'},

		// borderline values
		{name: "null rune", r: 0, want: 0},
		{name: "DEL", r: 127, want: 127},
		{name: "negative rune", r: -1, want: -1},

		// non-ASCII
		{name: "cyrillic lower", r: 'ж', want: 'ж'},
		{name: "latin lower with diacritic", r: 'é', want: 'é'},
		{name: "emoji", r: '😀', want: '😀'}}

	for _, tt := range testCases {
		got := ToUpper(tt.r)
		if got != tt.want {
			t.Errorf("%v   ToUpper()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}

func TestToLower(t *testing.T) {
	testCases := []struct {
		name string
		r    rune
		want rune
	}{ // upper
		{name: "upper A", r: 'A', want: 'a'},
		{name: "upper M", r: 'M', want: 'm'},
		{name: "upper Z", r: 'Z', want: 'z'},

		// lower
		{name: "lower a", r: 'a', want: 'a'},
		{name: "lower m", r: 'm', want: 'm'},
		{name: "lower z", r: 'z', want: 'z'},

		// neighbour values
		{name: "at sign before A", r: '@', want: '@'},
		{name: "bracket after Z", r: '[', want: '['},
		{name: "backtick before a", r: '`', want: '`'},
		{name: "brace after z", r: '{', want: '{'},

		// nums and stuff
		{name: "digit", r: '5', want: '5'},
		{name: "space", r: ' ', want: ' '},
		{name: "underscore", r: '_', want: '_'},
		{name: "newline", r: '\n', want: '\n'},

		// borderline values
		{name: "null rune", r: 0, want: 0},
		{name: "DEL", r: 127, want: 127},
		{name: "negative rune", r: -1, want: -1},

		// non-ASCII
		{name: "cyrillic upper", r: 'Ж', want: 'Ж'},
		{name: "latin upper with diacritic", r: 'É', want: 'É'},
		{name: "emoji", r: '😀', want: '😀'}}
	for _, tt := range testCases {
		got := ToLower(tt.r)
		if got != tt.want {
			t.Errorf("%v   ToLower()   r:%v   got:%v   want:%v", tt.name, tt.r, got, tt.want)
		}
	}
}
