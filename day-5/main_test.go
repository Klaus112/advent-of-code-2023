package main

import "testing"

func Test_converter_toOutput(t *testing.T) {
	type fields struct {
		sourceRangeStart      int
		destinationRangeStart int
		rangeLength           int
	}
	tests := map[string]struct {
		in     int
		fields fields
		want   int
	}{
		"in range minimum": {
			in: 95,
			fields: fields{
				sourceRangeStart:      95,
				destinationRangeStart: 112,
				rangeLength:           5,
			},
			want: 112,
		},
		"in range max": {
			in: 99,
			fields: fields{
				sourceRangeStart:      98,
				destinationRangeStart: 112,
				rangeLength:           2,
			},
			want: 113,
		},
		"out of range min": {
			in: 100,
			fields: fields{
				sourceRangeStart:      98,
				destinationRangeStart: 112,
				rangeLength:           2,
			},
			want: 0,
		},
		"out of range max": {
			in: 97,
			fields: fields{
				sourceRangeStart:      98,
				destinationRangeStart: 112,
				rangeLength:           2,
			},
			want: 0,
		},
		"seed 77 = soil 81": {
			in: 79,
			fields: fields{
				sourceRangeStart:      50,
				destinationRangeStart: 52,
				rangeLength:           48,
			},
			want: 81,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := converter{
				sourceRangeStart:      tt.fields.sourceRangeStart,
				destinationRangeStart: tt.fields.destinationRangeStart,
				rangeLength:           tt.fields.rangeLength,
			}
			if got := c.toOutput(tt.in); got != tt.want {
				t.Errorf("converter.toOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
