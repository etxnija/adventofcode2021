package day6

import "testing"

func Test_simulate(t *testing.T) {
	type args struct {
		scol []int
		days int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Example",
		args: args{scol: []int{3, 4, 3, 1, 2}, days: 80},
		want: 5934,
	},
		{
			name: "input",
			args: args{scol: fish(), days: 256},
			want: 8,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateWithCount(tt.args.scol, tt.args.days); got != tt.want {
				t.Errorf("simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
