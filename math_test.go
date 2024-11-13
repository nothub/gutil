package gutil

import "testing"

func TestClamp(t *testing.T) {
	type args struct {
		number float64
		min    float64
		max    float64
	}
	tests := []struct {
		name string
		args args
		want float64
		err  bool
	}{
		{name: "number within range", args: args{number: 42, min: 40, max: 50}, want: 42},
		{name: "number equal to min", args: args{number: 40, min: 40, max: 50}, want: 40},
		{name: "number equal to max", args: args{number: 50, min: 40, max: 50}, want: 50},
		{name: "number below min", args: args{number: 39, min: 40, max: 50}, want: 40},
		{name: "number above max", args: args{number: 51, min: 40, max: 50}, want: 50},
		{name: "negative number within range", args: args{number: -10, min: -20, max: -5}, want: -10},
		{name: "negative number below min", args: args{number: -25, min: -20, max: -5}, want: -20},
		{name: "negative number above max", args: args{number: 0, min: -20, max: -5}, want: -5},
		{name: "min and max are equal", args: args{number: 10, min: 10, max: 10}, want: 10},
		{name: "number below equal min and max", args: args{number: 5, min: 10, max: 10}, want: 10},
		{name: "number above equal min and max", args: args{number: 15, min: 10, max: 10}, want: 10},
		{name: "number in large range", args: args{number: 500, min: 100, max: 1000}, want: 500},
		{name: "number below large range", args: args{number: 50, min: 100, max: 1000}, want: 100},
		{name: "number above large range", args: args{number: 1500, min: 100, max: 1000}, want: 1000},

		// edge case: min equals max
		{name: "min equals max, number below", args: args{number: 5, min: 10, max: 10}, want: 10, err: false},
		{name: "min equals max, number above", args: args{number: 15, min: 10, max: 10}, want: 10, err: false},
		{name: "min equals max, number equal", args: args{number: 10, min: 10, max: 10}, want: 10, err: false},

		// edge case: min greater than max
		{name: "min greater than max, number within range", args: args{number: 12, min: 15, max: 10}, want: 0, err: true},
		{name: "min greater than max, number below min", args: args{number: 8, min: 15, max: 10}, want: 0, err: true},
		{name: "min greater than max, number above max", args: args{number: 20, min: 15, max: 10}, want: 0, err: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Clamp(tt.args.number, tt.args.min, tt.args.max)

			if tt.err {
				// should error
				if err == nil {
					// but did not error
					t.Fatalf("expected error: %v", err)
				}
				return
			}

			if !tt.err {
				// should not error
				if err != nil {
					// but did error
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if got != tt.want {
				t.Fatalf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
