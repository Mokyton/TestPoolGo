package main

import (
	"reflect"
	"testing"
)

func Test_minCoins(t *testing.T) {
	type args struct {
		val   int
		coins []int
	}
	tests := []struct {
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCoins(tt.args.val, tt.args.coins); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCoins2(t *testing.T) {
	type args struct {
		val   int
		coins []int
	}
	tests := []struct {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCoins2(tt.args.val, tt.args.coins); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCoins2() = %v, want %v", got, tt.want)
			}
		})
	}
}
