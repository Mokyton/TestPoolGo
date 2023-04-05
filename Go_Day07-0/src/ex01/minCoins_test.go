package main

import "testing"

type args struct {
	val   int
	coins []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "test_1",
		args: args{val: 13, coins: []int{1, 5, 10}},
		want: []int{10, 1, 1, 1},
	},
	{
		name: "test_2",
		args: args{val: 15, coins: []int{2, 3, 5}},
		want: []int{5, 5, 5},
	},
	{
		name: "test_3",
		args: args{val: 21, coins: []int{1, 5, 10}},
		want: []int{10, 10, 1},
	},
	{
		name: "test_4",
		args: args{val: 21, coins: []int{10, 5, 1}},
		want: []int{10, 10, 1},
	},
	{
		name: "test_5",
		args: args{val: 10, coins: []int{}},
		want: []int{},
	},
}

func BenchmarkMinCoins(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				minCoins(test.args.val, test.args.coins)
			}
		})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				minCoins(test.args.val, test.args.coins)
			}
		})
	}
}
