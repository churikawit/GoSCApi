package scapi

import (
	"testing"
	"fmt"
)

func TestBin2Str(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"5A", "3541"}, // "5A" -> [0x35][0x41] -> "3541"
	}

	for _, c := range cases {
		got := Bin2Str(c.in)
		if got != c.want {
			t.Errorf("Bin2Str(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStr2byte(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"A000000054480001", "A000000054480001"},
	}

	for _, c := range cases {
		got := str2byte(c.in)
		got2 := fmt.Sprintf("%X", got)
		if got2 != c.want {
			t.Errorf("str2byte(%q) == %q, want %q", c.in, got, c.want)
		}
	}
} 
