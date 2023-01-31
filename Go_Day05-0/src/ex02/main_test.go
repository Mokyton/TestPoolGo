package main

import (
	"ex02/myheap"
	"reflect"
	"testing"
)

func Test_getNCoolestPresents(t *testing.T) {
	type args struct {
		src []myheap.Present
		n   int
	}
	tests := []struct {
		name    string
		args    args
		want    []myheap.Present
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{n: 2, src: []myheap.Present{
				myheap.Present{5, 1},
				myheap.Present{4, 5},
				myheap.Present{3, 1},
				myheap.Present{5, 2},
			}},
			want:    []myheap.Present{myheap.Present{Value: 5, Size: 1}, {5, 2}},
			wantErr: false,
		},
		{
			name: "test_2",
			args: args{n: 2, src: []myheap.Present{
				myheap.Present{10, 1},
				myheap.Present{4, 5},
				myheap.Present{15, 1},
				myheap.Present{5, 2},
			}},
			want:    []myheap.Present{myheap.Present{Value: 15, Size: 1}, {10, 1}},
			wantErr: false,
		},
		{
			name: "test_3",
			args: args{n: -1, src: []myheap.Present{
				myheap.Present{10, 1},
				myheap.Present{4, 5},
				myheap.Present{15, 1},
				myheap.Present{5, 2},
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test_4",
			args: args{n: 5, src: []myheap.Present{
				myheap.Present{10, 1},
				myheap.Present{4, 5},
				myheap.Present{15, 1},
				myheap.Present{5, 2},
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNCoolestPresents(tt.args.src, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNCoolestPresents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNCoolestPresents() got = %v, want %v", got, tt.want)
			}
		})
	}
}
