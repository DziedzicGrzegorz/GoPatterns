package _726

import (
	"testing"
)

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				formula: "H2O",
			},
			want: "H2O",
		},
		{
			name: "Test 2",
			args: args{
				formula: "Mg(OH)2",
			},
			want: "H2MgO2",
		},
		{
			name: "Test 3",
			args: args{
				formula: "K4(ON(SO3)2)2",
			},
			want: "K4N2O14S4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countOfAtoms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countOfAtoms("K4(ON(SO3)2)2")
	}
}
