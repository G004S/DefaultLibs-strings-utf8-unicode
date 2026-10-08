package mystrings

import "testing"

func TestIndex(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		substr string
		want   int
	}{}

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
	}{}

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
	}{}

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
	}{}

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
	}{}
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
	}{}

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
	}{}
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
	}{}

	for _, tt := range testCases {
		gotS, gotB := Repeat(tt.s, tt.count)
		if gotS != tt.wantS || gotB != tt.wantB {
			t.Errorf("%v   Repeat()   s:%v   count:%v   got:%v; %v;   want:%v; %v", tt.name, tt.s, tt.count, gotS, gotB, tt.wantS, tt.wantB)
		}
	}
}
