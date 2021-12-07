package day3

import (
	"reflect"
	"testing"
)

func Test_diagnose(t *testing.T) {
	type args struct {
		report []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy",
			args: args{
				report: []string{
					"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			},
			want: 198,
		},
		{
			name: "fromfile",
			args: args{
				report: readReport(),
			},
			want: 3847100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagnose(tt.args.report); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagnose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oxygenRating(t *testing.T) {
	type args struct {
		reports []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy",
			args: args{
				reports: []string{
					"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			},
			want: 230,
		},
		{
			name: "from File",
			args: args{
				reports: readReport(),
			},
			want: 4105235,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oxygenRating(tt.args.reports); got != tt.want {
				t.Errorf("oxygenRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitSlice(t *testing.T) {
	type args struct {
		mesurements  []int
		bitPossition int
		matchFunc    func(int, int) bool
	}
	tests := []struct {
		name          string
		args          args
		wantMatched   []int
		wantUnmatched []int
	}{
		{
			name: "Happy path",
			args: args{mesurements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 22, 55, 67, 68},
				bitPossition: 0,
				matchFunc:    matchOnSetBit},
			wantMatched:   []int{1, 3, 5, 7, 9, 55, 67},
			wantUnmatched: []int{2, 4, 6, 8, 10, 22, 68},
		},
		{
			name: "Happy path third bit",
			args: args{mesurements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 22, 55, 67, 68},
				bitPossition: 2,
				matchFunc:    matchOnSetBit},
			wantMatched:   []int{4, 5, 6, 7, 22, 55, 68},
			wantUnmatched: []int{1, 2, 3, 8, 9, 10, 67},
		},
		{
			name: "third bit is unset",
			args: args{mesurements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 22, 55, 67, 68},
				bitPossition: 2,
				matchFunc:    matchOnUnSetBit},
			wantMatched:   []int{1, 2, 3, 8, 9, 10, 67},
			wantUnmatched: []int{4, 5, 6, 7, 22, 55, 68},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatched, gotUnmatched := splitSlice(tt.args.mesurements, tt.args.bitPossition, tt.args.matchFunc)
			if !reflect.DeepEqual(gotMatched, tt.wantMatched) {
				t.Errorf("splitSlice() gotMatched = %v, want %v", gotMatched, tt.wantMatched)
			}
			if !reflect.DeepEqual(gotUnmatched, tt.wantUnmatched) {
				t.Errorf("splitSlice() gotUnmatched = %v, want %v", gotUnmatched, tt.wantUnmatched)
			}
		})
	}
}
