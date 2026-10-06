package sl_test

import (
	"testing"

	sl "github.com/ppechkurov/sobesrazborgoogle/internal/slices"
)

func TestInfo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		slice []int
		want  int
		want2 int
	}{
		{
			name:  "[]int{1,2,3,4}",
			slice: []int{1, 2, 3, 4},
			want:  4,
			want2: 4,
		},
		{
			name:  "[]int{1,2,3,4}[:2]",
			slice: []int{1, 2, 3, 4}[:2],
			want:  2,
			want2: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, got2 := sl.Info(tt.slice)
			if got != tt.want {
				t.Errorf("Info() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("Info() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
