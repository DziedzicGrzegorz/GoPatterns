package samplehash

import (
	"testing"
)

func TestMakeHash(t *testing.T) {
	type args struct {
		value interface{}
	}
	type testCase struct {
		name string
		args args
		want uint64
	}

	tests := []testCase{
		{
			name: "string",
			args: args{value: "Gopher"},
			want: MakeHash("Gopher"),
		},
		{
			name: "int",
			args: args{value: 12345},
			want: MakeHash(12345),
		},
		{
			name: "slice",
			args: args{value: []int{1, 2, 3, 4, 5}},
			want: MakeHash([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.args.value.(type) {
			case string:
				if got := MakeHash(v); got != tt.want {
					t.Errorf("MakeHash() = %v, want %v", got, tt.want)
				}
			case int:
				if got := MakeHash(v); got != tt.want {
					t.Errorf("MakeHash() = %v, want %v", got, tt.want)
				}
			case []int:
				if got := MakeHash(v); got != tt.want {
					t.Errorf("MakeHash() = %v, want %v", got, tt.want)
				}
			default:
				t.Errorf("Unsupported type: %T", v)
			}
		})
	}
}
