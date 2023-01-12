package mode

import (
	"fmt"
	"testing"
)

func TestMode(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 3, 5}, 1},
		{[]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 4}, 1},
		{[]int{0, 1, 2, 3, 4, 5}, 0},
		{[]int{12, 25, 36, 14, 12, 12}, 12},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case(%v)", test.arr)
		t.Run(name, func(t *testing.T) {
			got := Mode(test.arr)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
