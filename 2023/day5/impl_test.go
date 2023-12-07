package day5

import (
	"testing"
)

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
			want:    35,
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

func Test_result_convertToDest(t *testing.T) {
	type fields struct {
		seeds  []int
		From   string
		To     string
		ranges []ranges
	}
	type args struct {
		source int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "in map 98 ",
			fields: fields{
				ranges: []ranges{
					{
						DestinationRange: 50,
						SourceRange:      98,
						RangeLength:      2,
					},
					{
						DestinationRange: 52,
						SourceRange:      50,
						RangeLength:      48,
					},
				},
			},
			args: args{
				source: 98,
			},
			want: 50,
		},
		{
			name: "in map 99 ",
			fields: fields{
				ranges: []ranges{
					{
						DestinationRange: 50,
						SourceRange:      98,
						RangeLength:      2,
					},
					{
						DestinationRange: 52,
						SourceRange:      50,
						RangeLength:      48,
					},
				},
			},
			args: args{
				source: 99,
			},
			want: 51,
		},
		{
			name: "in map 53 ",
			fields: fields{
				ranges: []ranges{
					{
						DestinationRange: 50,
						SourceRange:      96,
						RangeLength:      2,
					},
					{
						DestinationRange: 52,
						SourceRange:      50,
						RangeLength:      48,
					},
				},
			},
			args: args{
				source: 53,
			},
			want: 55,
		},
		{
			name: "not in map 10 ",
			fields: fields{
				ranges: []ranges{
					{
						DestinationRange: 50,
						SourceRange:      96,
						RangeLength:      2,
					},
					{
						DestinationRange: 52,
						SourceRange:      50,
						RangeLength:      48,
					},
				},
			},
			args: args{
				source: 10,
			},
			want: 10,
		},
		{
			name: "in map 79 ",
			fields: fields{
				ranges: []ranges{
					{
						DestinationRange: 50,
						SourceRange:      96,
						RangeLength:      2,
					},
					{
						DestinationRange: 52,
						SourceRange:      50,
						RangeLength:      48,
					},
				},
			},
			args: args{
				source: 79,
			},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := result{
				seeds:  tt.fields.seeds,
				From:   tt.fields.From,
				To:     tt.fields.To,
				ranges: tt.fields.ranges,
			}
			if got := r.convertToDest(tt.args.source); got != tt.want {
				t.Errorf("result.convertToDest() = %v, want %v", got, tt.want)
			}
		})
	}
}
