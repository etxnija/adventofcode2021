package day4

import (
	"reflect"
	"testing"
)

func Test_processRow(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name    string
		args    args
		want    result
		wantErr bool
	}{
		{
			name: "example row",
			args: args{
				row: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			},
			want: result{
				winning: []int{41, 48, 83, 86, 17},
				my:      []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processRow(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("processRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
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
			name: "example",
			args: args{
				in: "example.txt",
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "input",
			args: args{
				in: "input.txt",
			},
			want:    26346,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
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
			name: "example",
			args: args{
				in: "example.txt",
			},
			want:    13,
			wantErr: false,
		},
		// {
		// 	name: "input",
		// 	args: args{
		// 		in: "input.txt",
		// 	},
		// 	want:    26346,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solve2(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
