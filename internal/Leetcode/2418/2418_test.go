package _2418

import (
	"reflect"
	"testing"
)

func Test_sortPeople(t *testing.T) {
	type args struct {
		names   []string
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				names:   []string{"Mary", "John", "Emma"},
				heights: []int{180, 165, 170},
			},
			want: []string{"Mary", "Emma", "John"},
		},
		{
			name: "Test case 2",
			args: args{
				names:   []string{"Alice", "Bob", "Bob"},
				heights: []int{155, 185, 150},
			},
			want: []string{"Bob", "Alice", "Bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeople(tt.args.names, tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_sortPeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	}
}
