package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	sum := MySum(2, 4)
	if sum != 6 {
		t.Errorf("Expect 6 but received: %v", sum)
	}
}
