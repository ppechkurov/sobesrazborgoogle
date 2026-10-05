package sums

import (
	"testing"
)

func Test_getIndices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
		want2  int
		ok     bool
	}{
		{
			name:   "1 2 3 4=0 1 true",
			nums:   []int{1, 2, 3, 4},
			target: 3,
			want:   0,
			want2:  1,
			ok:     true,
		},
		{
			name:   "3 4 1 1=2 3 true",
			nums:   []int{3, 4, 1, 1},
			target: 2,
			want:   2,
			want2:  3,
			ok:     true,
		},
		{
			name:   "3 4 1 1=-1 -1 false",
			nums:   []int{},
			target: 2,
			want:   -1,
			want2:  -1,
			ok:     false,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got2, ok := getIndices(tt.nums, tt.target)
			if ok != tt.ok {
				t.Errorf("success expected")
				t.FailNow()
			}

			if got != tt.want && got2 != tt.want2 {
				t.Errorf("getIndices() = %v, %v; want %v, %v", got, got2, tt.want, tt.want2)
			}
		})
	}
}
