package tasks

import (
	"testing"
)

func TestReturnRandomSlice(t *testing.T) {

	if len(returnRandomSlice(2)) != 2 {
		t.Error()
	}
}

func TestFactorial(t *testing.T) {

	if factorial(1) != 1 {
		t.Error()
	}
}
