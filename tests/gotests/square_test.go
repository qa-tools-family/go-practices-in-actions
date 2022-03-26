package math

import "testing"

func TestSquare(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{m: 1}, 1},
		{"case2", args{m: 2}, 4},
		{"case3", args{m: -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.m); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}
