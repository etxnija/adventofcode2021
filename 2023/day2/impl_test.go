package day2

import (
	"reflect"
	"testing"
)

func Test_color_GetColor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    color
		wantErr bool
	}{
		{
			name: "green",
			args: args{
				s: "green",
			},
			want:    green,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetColor(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("color.GetColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("color.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pulls(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []map[color]int
	}{
		{
			name: "happy",
			args: args{
				s: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: []map[color]int{
				{
					blue: 3,
					red:  4,
				},
				{
					red:   1,
					green: 2,
					blue:  6,
				},
				{
					green: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pulls(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pulls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_possible(t *testing.T) {
	type args struct {
		in     string
		limits map[color]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example1",
			args: args{
				in: "example.txt",
				limits: map[color]int{
					blue:  14,
					green: 13,
					red:   12,
				},
			},
			want:    8,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				in: "input.txt",
				limits: map[color]int{
					blue:  14,
					green: 13,
					red:   12,
				},
			},
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := possible(tt.args.in, tt.args.limits)
			if (err != nil) != tt.wantErr {
				t.Errorf("possible() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("possible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example1",
			args: args{
				in: "example.txt",
			},
			want:    2286,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				in: "input.txt",
			},
			want:    66909,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := min(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
