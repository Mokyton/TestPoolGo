package median

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	var tests = []struct {
		a    []int
		want float64
	}{
		{[]int{1, 4, 2, 5, 0}, 2},
		{[]int{10, 40, 20, 50}, 30},
		{[]int{10, 15, 16, 18, 20}, 16},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case(%v)", test.a)
		t.Run(name, func(t *testing.T) {
			got := Median(test.a)
			if got != test.want {
				t.Errorf("got %f, want %f", got, test.want)
			}
		})
	}
}
