package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello its hillheim")
	hashBroadcast(NewHashReader(payload))
}

type Hashreader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}
func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashBroadcast(r Hashreader) error {

	hash := r.hash()
	a, err := r.Read([]byte(" "))
	if err != nil {
		return err
	}
	fmt.Println(a)
	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of the bytes:", string(b))
	return nil
}
