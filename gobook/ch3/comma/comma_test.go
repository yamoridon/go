package comma

import "testing"

func Test(t *testing.T) {
	if comma("12345") != "12,345" {
		t.Error(`comma("12345") != "12,345"`)
	}
}
