package ex00

import "testing"

func Test_getElement(t *testing.T) {
	type args struct {
		arr []int
		idx int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test_1",
			args:    args{arr: []int{0, 1, 2, 3, 4, 5}, idx: 2},
			want:    2,
			wantErr: false,
		},
		{
			name:    "test_2",
			args:    args{arr: []int{0, 1, 2, 3, 4, 5}, idx: -2},
			want:    0,
			wantErr: true,
		},
		{
			name:    "test_3",
			args:    args{arr: []int{0, 1, 2, 3, 4, 5}, idx: 6},
			want:    0,
			wantErr: true,
		},
		{
			name:    "test_4",
			args:    args{arr: []int{0, 1, 2, 3, 4, 5}, idx: 10},
			want:    0,
			wantErr: true,
		},
		{
			name:    "test_5",
			args:    args{arr: []int{}, idx: 2},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getElement(tt.args.arr, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}
