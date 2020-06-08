package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying
// the rot13 substitution cipher to all alphabetical characters.

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	isCapital := x >= 'A' && x <= 'Z'
	if !isCapital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	fmt.Println(x)
	x += 13
	if isCapital && x > 'Z' || !isCapital && x > 'z' {
		x -= 26
	}
	return x
}

func (r13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
