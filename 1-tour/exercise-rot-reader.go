package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/**
 * This basically compares the byte letter code with the relevant runes
 * If it's between A and M we return the number + 13, from N-Z we return the number - 13
 *
 * We also do the same for a-m and n-z. Finally, we have a fallthrough that returns b
 */
func rot13(b byte) byte {
	// I'm aware this can be chained with an OR to make it two lines,
	// but I think it's more readable this way
	if b >= 'A' && b <= 'M' {
		return b + 13
	} else if b >= 'N' && b <= 'Z' {
		return b - 13
	} else if b >= 'a' && b <= 'm' {
		return b + 13
	} else if b >= 'n' && b <= 'z' {
		return b - 13
	}
	return b
}

func (rot rot13Reader) Read(data []byte) (n int, err error) {
	// As per the Docs, n here represents the amount of bytes read.
	// In this case with the string "Lbh penpxrq gur pbqr!" it's 21
	// Err should always be nil
	n, err = rot.r.Read(data)

	// We overwrite the data as long as our bytes are within the first
	// 21, after that we get a lot of 0's. Not sure why.
	for i := 0; i < n; i++ {
		// We replace each byte/character with their rot13 equivalent
		data[i] = rot13(data[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	// You cracked the code!
	io.Copy(os.Stdout, &r)
}
