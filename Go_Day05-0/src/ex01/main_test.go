package main

import (
	"reflect"
	"testing"
)

func Test_unrollGarland(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "test_1",
			args: args{
				&TreeNode{HasToy: true,
					Left: &TreeNode{HasToy: true,
						Left:  &TreeNode{HasToy: true},
						Right: &TreeNode{HasToy: false}},
					Right: &TreeNode{HasToy: false,
						Left:  &TreeNode{HasToy: true},
						Right: &TreeNode{HasToy: true}}},
			},
			want: []bool{true, true, false, true, true, false, true},
		},
		{
			name: "test_2",
			args: args{
				&TreeNode{HasToy: true,
					Left: &TreeNode{HasToy: true,
						Left: &TreeNode{HasToy: false,
							Left:  &TreeNode{HasToy: true},
							Right: &TreeNode{HasToy: true}},
						Right: &TreeNode{HasToy: true,
							Left:  &TreeNode{HasToy: false},
							Right: &TreeNode{HasToy: false}}},
					Right: &TreeNode{HasToy: false,
						Left: &TreeNode{HasToy: false,
							Left:  &TreeNode{HasToy: false},
							Right: &TreeNode{HasToy: false}},
						Right: &TreeNode{HasToy: true,
							Left:  &TreeNode{HasToy: false},
							Right: &TreeNode{HasToy: false}},
					},
				},
			},
			want: []bool{
				true,
				true, false,
				true, false, true, false,
				true, true, false, false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unrollGarland(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unrollGarland() = %v, want %v", got, tt.want)
			}
		})
	}
}
