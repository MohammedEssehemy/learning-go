package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rd.r.Read(bytes)
	for i := 0; i < len(bytes); i++ {
		if (bytes[i] >= 'a' && bytes[i] <= 'm') || (bytes[i] >= 'A' && bytes[i] <= 'M') {
			bytes[i] += 13
		} else if (bytes[i] >= 'n' && bytes[i] <= 'z') || (bytes[i] >= 'N' && bytes[i] <= 'Z') {
			bytes[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
