package day1

import "testing"

func Test_distance(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				file: "example.txt",
			},
			want: 11,
		},
		{
			name: "example",
			args: args{
				file: "input.txt",
			},
			want: 1590491,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.file); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
