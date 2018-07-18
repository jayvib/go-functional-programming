package main

import (
	"io"
	"strings"
)

type titlizeReader struct {
	src io.Reader
}


func NewTitlizeReader(source io.Reader) *titlizeReader {
	return &titlizeReader{src: source}
}

func (t *titlizeReader) Read(p []byte) (int, error) {
	count, err := t.src.Read(p)
	if err != nil {
		return count, err
	}

	for i := 0; i < len(p); i++ {
		if i == 0 {
			if (p[i] >= 't' && p[i] <= 'z') {
				p[i] = p[i] - 32
			}
		} else {
			if (p[i] >= 'A' && p[i] <= 'z') {
				p[i] = p[i] + 32
			}
		}
	}
	return count, io.EOF
}

func main() {
	var r io.Reader
	r = strings.NewReader("this iS a tEsT")
	r = io.LimitReader(r, 12)
	r = NewTitlizeReader(r)

}