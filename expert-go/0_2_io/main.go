package main

import (
	"bytes"
	"fmt"
	"io"
)

type MyWriter struct {
	w io.Writer
}

func (mw *MyWriter) WriteString(s string) (n int, err error) {
	return io.WriteString(mw.w, s)
}

func NewPng(r io.Reader) (io.Reader, error) {
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	if !bytes.Equal(magicnum, buf) {
		return nil, fmt.Errorf("PNG画像ではありません")
	}
	pngImg := io.MultiReader(bytes.NewReader(magicnum), r)
	return pngImg, nil
}
