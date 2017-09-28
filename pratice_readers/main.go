package main

import (
	"strings"
	"fmt"
	"io"
	"golang.org/x/tour/reader"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n,err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
			p[i] -= 13
		}
	}
	return
}

type MyReader struct {}

func (r MyReader ) Read(s []byte) (int, error) {
	for i := range s {
		s[i] = 'A'
	}
	return len(s), nil
}

func main() {
	reader.Validate(MyReader{})

	r := strings.NewReader("Hello, Reader!")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rotReader := rot13Reader{s}
	io.Copy(os.Stdout, &rotReader)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
