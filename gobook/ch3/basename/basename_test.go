package basename

import "testing"

func Test(t *testing.T) {
	if basename("a") != "a" {
		t.Error(`basename("a") != "a"`)
	}
	if basename("a.go") != "a" {
		t.Error(`basename("a.go") != "a"`)
	}
	if basename("a/b/c.go") != "c" {
		t.Error(`basename("a/b/c.go") != "c"`)
	}
	if basename("a/b.c.go") != "b.c" {
		t.Error(`basename("a/b.c.go") != "b.c"`)
	}
}
