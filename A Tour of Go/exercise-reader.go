package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}
func main() {
	reader.Validate(MyReader{})
}
wBPM0Oo6PZxGCR2
