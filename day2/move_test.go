package day2

import (
	"reflect"
	"testing"
)

func Test_move(t *testing.T) {
	type args struct {
		movment []movment
		start   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example",
			args{
				movment: []movment{
					NewMovement("forward", 5),
					NewMovement("down", 5),
					NewMovement("forward", 8),
					NewMovement("up", 3),
					NewMovement("down", 8),
					NewMovement("forward", 2),
				},
				start: []int{0, 0},
			},
			[]int{15, 10},
		},
		{
			"from file",
			args{
				movment: readMovements(),
				start:   []int{0, 0},
			},
			[]int{15, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.start, tt.args.movment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveWithAim(t *testing.T) {
	type args struct {
		start   []int
		movment []movment
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example",
			args{
				movment: []movment{
					NewMovement("forward", 5),
					NewMovement("down", 5),
					NewMovement("forward", 8),
					NewMovement("up", 3),
					NewMovement("down", 8),
					NewMovement("forward", 2),
				},
				start: []int{0, 0, 0},
			},
			[]int{15, 60, 10},
		},
		{
			"from file",
			args{
				movment: readMovements(),
				start:   []int{0, 0, 0},
			},
			[]int{1970, 1000556, 916},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveWithAim(tt.args.start, tt.args.movment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveWithAim() = %v, want %v", got, tt.want)
			}
		})
	}
}
