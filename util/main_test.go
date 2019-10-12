package util

import (
	"testing"
)

func TestCompareInt(t *testing.T) {
	t.Run("max", func(t *testing.T) {
		res := CompareInt("max", 2, 1)
		if res != 2 {
			t.Fatalf("Fail to compare. Expected: 2 but got %d", res)
		}
	})

	t.Run("min", func(t *testing.T) {
		res := CompareInt("min", 2, 1)
		if res != 1 {
			t.Fatalf("Fail to compare. Expected: 1 but got %d", res)
		}
	})

	t.Run("unknown mode", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != "Unknown mode" {
				t.Errorf("got %v\nwant %v", err, "Unknown mode")
			}
		}()

		_ = CompareInt("hoge", 2, 1)
	})
}
