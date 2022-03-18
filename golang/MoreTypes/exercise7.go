package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		// lowercase
		ch = (ch - 'a' + 13) % 26
		ch += 'a'
		return ch
	} else {
		// uppercase
		ch = (ch - 'A' + 13) % 26
		ch += 'A'
		return ch
	}
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	ss, err1 := rot.r.Read(b)
	if err1 != nil {
		return 0, err1
	}
	for i := 0; i < ss; i++ {
		b[i] = rot13(b[i])
	}
	return ss, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
