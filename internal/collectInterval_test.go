package internal

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {
	type args struct {
		data []Interval
	}
	tests := []struct {
		name string
		args args
		want []ResultInterval
	}{
		{
			name: "sample 1",
			args: args{
				data: []Interval{
					{10, 50, "A"},
					{40, 100, "B"},
					{30, 120, "C"},
				},
			},
			want: []ResultInterval{
				{10, 30, []string{"A"}},
				{30, 40, []string{"A", "C"}},
				{40, 50, []string{"A", "B", "C"}},
				{50, 100, []string{"B", "C"}},
				{100, 120, []string{"C"}},
			},
		},
		{
			name: "sample 2",
			args: args{
				data: []Interval{
					{10, 50, "A"},
					{40, 100, "B"},
					{110, 120, "C"},
				},
			},
			want: []ResultInterval{
				{10, 40, []string{"A"}},
				{40, 50, []string{"A", "B"}},
				{50, 100, []string{"B"}},
				{110, 120, []string{"C"}},
			},
		},
		{
			name: "sample 3",
			args: args{
				data: []Interval{
					{10, 50, "A"},
					{40, 100, "B"},
					{20, 120, "C"},
				},
			},
			want: []ResultInterval{
				{10, 20, []string{"A"}},
				{20, 40, []string{"A", "C"}},
				{40, 50, []string{"A", "B", "C"}},
				{50, 100, []string{"B", "C"}},
				{100, 120, []string{"C"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect(%v) = %v, want %v", tt.args.data, got, tt.want)
			}
		})
	}
}
