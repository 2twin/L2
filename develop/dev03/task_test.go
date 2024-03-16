package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lines = []string{
		"1 Dec 6 14:29 Afternoon",
		"2 Jun 10 2015 Search",
		"3 Oct 31 15:08 Voicechat",
		"4 Jan 13 11:42 Low",
		"5 Jan 11 12:33 Present",
	}
)

func TestSortEmptyInput(t *testing.T) {
	cs := &customSort{
		lines: []string{},
	}

	sorted := sortLines(cs)
	if len(sorted) > 0 {
		t.Errorf("expected 0 lines, got %d", len(sorted))
	}
}

func TestSortDefaultParams(t *testing.T) {
	cs := &customSort{
		lines: lines,
	}

	sorted := sortLines(cs)
	assert.Equal(t, lines, sorted)
}

func TestSortByMonth(t *testing.T) {
	cs := &customSort{
		lines:   lines,
		byMonth: true,
	}

	sorted := sortLines(cs)
	assert.Equal(t, []string{
		"4 Jan 13 11:42 Low",
		"5 Jan 11 12:33 Present",
		"2 Jun 10 2015 Search",
		"3 Oct 31 15:08 Voicechat",
		"1 Dec 6 14:29 Afternoon",
	}, sorted)
}

func TestSortByColumn(t *testing.T) {
	cs := &customSort{
		lines: lines,
		key:   5,
	}

	sorted := sortLines(cs)
	assert.Equal(t, []string{
		"1 Dec 6 14:29 Afternoon",
		"4 Jan 13 11:42 Low",
		"5 Jan 11 12:33 Present",
		"2 Jun 10 2015 Search",
		"3 Oct 31 15:08 Voicechat",
	}, sorted)
}

func TestSortNumeric(t *testing.T) {
	cs := &customSort{
		lines:   lines,
		key:     3,
		numeric: true,
	}

	sorted := sortLines(cs)
	assert.Equal(t, []string{
		"1 Dec 6 14:29 Afternoon",
		"2 Jun 10 2015 Search",
		"5 Jan 11 12:33 Present",
		"4 Jan 13 11:42 Low",
		"3 Oct 31 15:08 Voicechat",
	}, sorted)
}

func TestSortUnique(t *testing.T) {
	cs := &customSort{
		lines: []string{
			"golang",
			"rust",
			"golang",
			"python",
		},
		unique: true,
	}

	sorted := sortLines(cs)
	assert.Equal(t, []string{
		"golang",
		"python",
		"rust",
	}, sorted)
}

func TestSortReversed(t *testing.T) {
	cs := &customSort{
		lines:   lines,
		reverse: true,
	}

	sorted := sortLines(cs)
	assert.Equal(t, []string{
		"5 Jan 11 12:33 Present",
		"4 Jan 13 11:42 Low",
		"3 Oct 31 15:08 Voicechat",
		"2 Jun 10 2015 Search",
		"1 Dec 6 14:29 Afternoon",
	}, sorted)
}
