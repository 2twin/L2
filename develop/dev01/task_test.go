package main

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	_, gotErr := GetTime()
	if gotErr != nil {
		t.Errorf("error running GetTime: %s", gotErr.Error())
	}
}