package bitmap_test

import (
	"testing"

	"github.com/adityathebe/bitmap"
)

func TestSetAndUnset(t *testing.T) {
	bm := bitmap.NewBitmap(1000)
	testData := []int{23, 44, 124, 55}

	t.Run("setting", func(t *testing.T) {
		for i, td := range testData {
			bm.Set(td)
			if !bm.IsSet(td) {
				t.Fatalf("[%d] Expected bit [%d] to bet set\n", i, td)
			}
		}
	})

	t.Run("Unsetting", func(t *testing.T) {
		for i, td := range testData {
			bm.Unset(td)
			if bm.IsSet(td) {
				t.Fatalf("[%d] Expected bit [%d] to bet unset\n", i, td)
			}
		}
	})
}
