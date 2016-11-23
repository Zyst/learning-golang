package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// How do we 'rot?'
//
// Follow up: bytes can just be +'d to achieve a similar effect, ie: b + 13
//
// Also need to find a way to "Roll around for letters like Z -> A + 12"
func rot13(b byte) byte {
	return b
}

// Figure out how to do this tomorrow brain is dead today
func (read rot13Reader) Read(data []byte) (n int, err error) {
	return 0, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
