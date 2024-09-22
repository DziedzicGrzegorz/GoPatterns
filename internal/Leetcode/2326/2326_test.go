package _2326

import (
	"reflect"
	"testing"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				m: 3,
				n: 3,
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
										Next: &ListNode{
											Val: 7,
											Next: &ListNode{
												Val: 8,
												Next: &ListNode{
													Val: 9,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSpiralMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spiralMatrix(3, 3, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 7,
									Next: &ListNode{
										Val: 8,
										Next: &ListNode{
											Val: 9,
										},
									},
								},
							},
						},
					},
				},
			},
		})
	}
}
