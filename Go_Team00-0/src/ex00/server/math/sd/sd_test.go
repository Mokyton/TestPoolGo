package sd

import (
	"fmt"
	"testing"
)

func TestSD(t *testing.T) {
	var tests = []struct {
		arr  []int
		want float32
	}{
		{[]int{46, 69, 32, 60, 52, 41}, 12.151817},
		{[]int{1, -3, 5, -9, 23, -478, 990, 1892, 43, 21, -7, 89, 92, 784, -223}, 564.551453},
	}
	for _, test := range tests {
		name := fmt.Sprintf("case(%v)", test.arr)
		t.Run(name, func(t *testing.T) {
			got := SD(test.arr)
			if got != test.want {
				t.Errorf("got %f, want %f", got, test.want)
			}
		})
	}
}
