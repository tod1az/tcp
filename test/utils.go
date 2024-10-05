package test

import (
	"bytes"
	"os"
)

func PipeStdout() (*os.File, *os.File) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	return r, w
}

func GetOutput(r, w *os.File) string {
	var buf bytes.Buffer
	w.Close()
	buf.ReadFrom(r)
	return buf.String()
}
