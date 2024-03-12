package main

import (
	"testing"
)

type Result struct {
	String string
	Error  error
}

func TestUnzip(t *testing.T) {
	cases := map[string]Result{
		"a4bc2d5e": {"aaaabccddddde", nil},
		"abcd":     {"abcd", nil},
		"45":       {"", ErrIncorrectString},
		"":         {"", ErrIncorrectString},
		"5oijsd":   {"", ErrIncorrectString},
		`qwe\4\5`:  {"qwe45", nil},
		`qwe\45`:   {"qwe44444", nil},
		`qwe\\5`:   {`qwe\\\\\`, nil},
		"$µ2€№3":   {"$µµ€№№№", nil},
	}

	for in, out := range cases {
		got, err := Unzip(in)

		if got != out.String && err != out.Error {
			t.Errorf("Want: %s, got: %s", got, out)
		}
	}
}
