package main

import "testing"

func Test_areToysBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_1",
			args: args{
				root: &TreeNode{HasToy: false,
					Left:  &TreeNode{HasToy: false, Left: &TreeNode{HasToy: false, Right: &TreeNode{HasToy: true}}},
					Right: &TreeNode{HasToy: true}}},
			want: true,
		},
		{
			name: "test_2",
			args: args{
				root: &TreeNode{HasToy: true,
					Left: &TreeNode{HasToy: true,
						Left:  &TreeNode{HasToy: true},
						Right: &TreeNode{HasToy: false}},
					Right: &TreeNode{HasToy: false,
						Left:  &TreeNode{HasToy: true},
						Right: &TreeNode{HasToy: true}},
				}},
			want: true,
		},
		{
			name: "test_3",
			args: args{
				root: &TreeNode{HasToy: true,
					Right: &TreeNode{HasToy: false},
					Left:  &TreeNode{HasToy: true}},
			},
			want: false,
		},
		{
			name: "test_4",
			args: args{
				root: &TreeNode{HasToy: false,
					Right: &TreeNode{HasToy: false,
						Right: &TreeNode{HasToy: true}},
					Left: &TreeNode{HasToy: true,
						Right: &TreeNode{HasToy: true}}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areToysBalanced(tt.args.root); got != tt.want {
				t.Errorf("areToysBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
