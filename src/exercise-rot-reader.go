package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if ('a'<=b && b<='m') || ('A'<=b && b<='M') {
		return b + 13
	}
	
	if ('N'<=b && b<='Z') || ('n'<=b && b<='z') {
		return b - 13
	}
	
	return b
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	//io.Copy(os.Stdout, s)
	io.Copy(os.Stdout, &r)
}

