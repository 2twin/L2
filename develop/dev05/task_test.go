package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lines = []string{
		"Asking Alexandria",
		"Bring Me The Horizon",
		"AC/DC, Metallica",
		"Bullet For My Valentine",
		"The Weeknd",
	}
)

func TestGrepEmptyParams(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		result: make(map[int]string),
	}

	target := "AC/DC"
	grep(cfg, target)

	want := map[int]string{
		2: "AC/DC, Metallica",
	}
	assert.Equal(t, want, cfg.result)
}

func TestGrepAfter(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		after:  6,
		result: make(map[int]string),
	}

	target := "Horizon"
	grep(cfg, target)

	want := map[int]string{
		1: "Bring Me The Horizon",
		2: "AC/DC, Metallica",
		3: "Bullet For My Valentine",
		4: "The Weeknd",
	}
	assert.Equal(t, want, cfg.result)
}

func TestGrepBefore(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		before: 1,
		result: make(map[int]string),
	}

	target := "Valentine"
	grep(cfg, target)

	want := map[int]string{
		2: "AC/DC, Metallica",
		3: "Bullet For My Valentine",
	}
	assert.Equal(t, want, cfg.result)
}

func TestGrepContext(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		ctx:    1,
		result: make(map[int]string),
	}

	target := "Metallica"
	grep(cfg, target)

	want := map[int]string{
		1: "Bring Me The Horizon",
		2: "AC/DC, Metallica",
		3: "Bullet For My Valentine",
	}
	assert.Equal(t, want, cfg.result)
}

func TestGrepCount(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		result: make(map[int]string),
	}

	target := "The"
	grep(cfg, target)

	want := 2
	assert.Equal(t, want, cfg.counter)
}

func TestGrepIgnoreCase(t *testing.T) {
	cfg := &grepConfig{
		lines:      lines,
		ignoreCase: true,
		result:     make(map[int]string),
	}

	target := "bring"
	grep(cfg, target)

	want := map[int]string{
		1: "Bring Me The Horizon",
	}
	assert.Equal(t, want, cfg.result)
}

func TestGrepInvert(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		invert: true,
		result: make(map[int]string),
	}

	target := "The"
	grep(cfg, target)

	want := 2
	assert.Equal(t, want, cfg.counter)
}

func TestGrepFixed(t *testing.T) {
	cfg := &grepConfig{
		lines:  lines,
		fixed:  true,
		result: make(map[int]string),
	}

	target := "Bullet"
	grep(cfg, target)

	want := map[int]string{}
	assert.Equal(t, want, cfg.result)
}

func TestGrepLineNum(t *testing.T) {
	cfg := &grepConfig{
		lines:      lines,
		lineNum:    true,
		ctx:        2,
		ignoreCase: true,
		result:     make(map[int]string),
	}

	target := "ac/dc"
	grep(cfg, target)

	want := map[int]string{
		0: "1-Asking Alexandria",
		1: "2-Bring Me The Horizon",
		2: "3:AC/DC, Metallica",
		3: "4-Bullet For My Valentine",
		4: "5-The Weeknd",
	}
	assert.Equal(t, want, cfg.result)
}
