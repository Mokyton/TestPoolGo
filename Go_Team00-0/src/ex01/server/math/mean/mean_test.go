package mean

import "testing"

func TestMeanCase1(t *testing.T) {
	got := Mean([]int{1, 5})
	want := float64(3)
	if got != want {
		t.Errorf("got %f; want %f", got, want)
	}
}

func TestMeanCase2(t *testing.T) {
	got := Mean([]int{42, 13, 31, 87, 24, 58, 76, 69})
	want := float64(50)
	if got != want {
		t.Errorf("got %f; want %f", got, want)
	}
}

func TestMeanCase3(t *testing.T) {
	got := Mean([]int{42, 13, 31, 87, 58, 76, 69, 230})
	want := 75.75
	if got != want {
		t.Errorf("got %f; want %f", got, want)
	}
}
