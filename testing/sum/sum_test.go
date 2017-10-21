package sum

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{3, 4, 7},
		{3, 5, 8},
		{3, 7, 10},
		{3, -4, -1},
	}

	for _, c := range tests {
		got := Sum(c.a, c.b)
		if got != c.want {
			t.Errorf("Sum(%d,%d) => %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
