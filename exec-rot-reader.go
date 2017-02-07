package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i, c := range b {
		switch {
		case 'a' <= c && c <= 'm' || 'A' <= c && c <= 'M':
			b[i] = c + 13
		case 'n' <= c && c <= 'z' || 'N' <= c && c <= 'Z':
			b[i] = c - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
