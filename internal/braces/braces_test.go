package braces

import "testing"

func Test_compare(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  []rune
		want bool
	}{
		{
			name: "(){}[]",
			str:  []rune("(){}[]"),
			want: true,
		},
		{
			name: ")(}{][",
			str:  []rune(")(}{]["),
			want: false,
		},
		{
			name: "()}{][",
			str:  []rune("()}{]["),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compare(tt.str)
			if got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
