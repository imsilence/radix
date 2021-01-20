package radix

import "testing"

func TestMin(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		if 1 != min(1, 2) {
			t.Error("min(1, 2) != 1")
		}
	})
	t.Run("case 2", func(t *testing.T) {
		if 2 != min(3, 2) {
			t.Error("min(3, 2) != 2")
		}
	})
}
