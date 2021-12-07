package day1

import (
	"testing"
)

func Test_increased(t *testing.T) {
	type args struct {
		mesurements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"happy",
			args{
				mesurements: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			7,
		},
		{
			"happy",
			args{
				mesurements: mesurements(),
			},
			1162,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasedSliding(tt.args.mesurements, 1); got != tt.want {
				t.Errorf("increased() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increasedSliding(t *testing.T) {
	type args struct {
		mesurements []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"happy",
			args{
				mesurements: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			5,
		},
		{
			"happy full",
			args{
				mesurements: mesurements(),
			},
			1190,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasedSliding(tt.args.mesurements, 3); got != tt.want {
				t.Errorf("increasedSliding() = %v, want %v", got, tt.want)
			}
		})
	}
}
