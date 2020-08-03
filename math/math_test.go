package math

import "testing"

func TestMath_Add(t *testing.T) {
	testCases := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{4, 5, 9},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			m := &Math{}
			got := m.Add(tc.a, tc.b)
			if tc.c != got {
				t.Errorf("Expected %v, got %v", tc.c, got)
			}
		})
	}
}
