package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = (b[i]-'a'+13)%26 + 'a'
		} else if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = (b[i]-'A'+13)%26 + 'A'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
