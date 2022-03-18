package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {

	if len(b) <= 0 {
		return 0, io.EOF
	}

	b[0] = 'A'

	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
