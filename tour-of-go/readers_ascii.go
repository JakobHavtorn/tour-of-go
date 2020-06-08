package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
		b[x] = 'A'
	}
	return len(b), nil
}

func main() {
	// reader.Validate(MyReader{})
	b := []byte{1, 2, 3, 4}
	r := MyReader{}
	fmt.Println(r.Read(b))
	fmt.Println(b)
}
