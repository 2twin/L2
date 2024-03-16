package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	in := []string{
		"Столик",
		"ПЯТАК",
		"тяпка",
		"Алгоритм",
		"лисТОК",
		"ПяткА",
		"слиток",
	}

	want := map[string][]string{
		"столик": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"},
	}

	got := findAnagrams(in)
	
	assert.Equal(t, want, got)
}
