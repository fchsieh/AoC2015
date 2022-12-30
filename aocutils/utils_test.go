package aocutils

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Error("Min(1, 2) != 1")
	}
	if Min(2, 1) != 1 {
		t.Error("Min(2, 1) != 1")
	}
}

func TestMinArray(t *testing.T) {
	arr := []int{1, 2, 3}
	if MinArray(arr) != 1 {
		t.Error("MinArray(arr) != 1")
	}
}

func TestMax(t *testing.T) {
	if Max(1, 2) != 2 {
		t.Error("Max(1, 2) != 2")
	}
	if Max(2, 1) != 2 {
		t.Error("Max(2, 1) != 2")
	}
}

func TestMaxArray(t *testing.T) {
	arr := []int{1, 2, 3}
	if MaxArray(arr) != 3 {
		t.Error("MaxArray(arr) != 3")
	}
}
