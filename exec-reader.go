package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error)  {
	count := 0
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
		count ++
	}

	return count, nil
}

func main() {
	reader.Validate(MyReader{})
}

