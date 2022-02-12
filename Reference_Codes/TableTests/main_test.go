package main
import (
	"testing"
)
func TestSum(t *testing.T) {
	type test struct {
		input []int
		answer int
	}
	cases := []test{
		test{[]int{1,2,3}, 6},
		test{[]int{10,11,12,13}, 46},
		test{[]int{1,2,3,4,5,6,7,8,9,10}, 55},
		test{[]int{-1,0,1}, 0},
	}
	for _, c := range cases {
		actual := Sum(c.input...)
		if actual != c.answer {
			t.Errorf("Sum(%v) == %d, expected %d", c.input, actual, c.answer)
		}
	}
}