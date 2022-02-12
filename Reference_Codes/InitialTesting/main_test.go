package main
import (
	"testing"
)
func TestSum(t *testing.T) {
	v := Sum(1, 2, 3, 4, 5)
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}