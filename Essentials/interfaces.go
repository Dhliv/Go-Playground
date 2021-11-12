package Essentials

import (
	"bytes"
	"fmt"
)

type consoleWriter struct {
}

func (cw consoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type writer interface {
	write([]byte) (int, error)
}

type closer interface {
	close() error
}

type writerCloser interface {
	writer
	closer
}

type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *bufferedWriterCloser) write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *bufferedWriterCloser) close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func newBufferedWriterCloser() *bufferedWriterCloser {
	return &bufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func interfaces() {
	var w writer = consoleWriter{}
	fmt.Println("algo")
	w.write([]byte("Hello GO!"))

	var wc writerCloser = newBufferedWriterCloser()
	wc.write([]byte("askjasdkjnasdkjaskdjsadasdasdasddsaasdsaasddsaasd"))
	wc.close()
}
