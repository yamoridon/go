package treesort

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var values = []int{1, 3, 2, 5, 4, 6, 7, 9, 8}
	Sort(values)
	var s = fmt.Sprintf("%v", values)
	if s != "[1 2 3 4 5 6 7 8 9]" {
		t.Error(`Sort is broken`)
	}
}\