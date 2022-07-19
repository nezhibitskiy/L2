package unpack

import (
	"bytes"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		isValid bool
		in      string
		out     string
	}{
		{
			true,
			"a4bc2d5e",
			"aaaabccddddde",
		},
		{
			true,
			"abcd",
			"abcd",
		},
		{
			true,
			"qwe\\4\\5",
			"qwe45",
		},
		{
			true,
			"qwe\\45",
			"qwe44444",
		},
		{
			false,
			"45",
			"",
		},
		{
			true,
			"",
			"",
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			if test.isValid == true {
				a, err := Unpack(test.in)
				if err != nil {
					t.Fatal()
				}
				inf := bytes.Equal([]byte(test.out), []byte(a))
				if !inf {
					t.Fatal()
				}
			} else {
				_, err := Unpack(test.in)
				if err == nil {
					t.Fatal()
				}
			}
		})
	}
}
