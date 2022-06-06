package genh

import (
	"io"
	"os"
)

type TypeEncoder interface {
	Encode(v any) error
}

type TypeDecoder interface {
	Decode(v any) error
}

func EncodeFile(fp string, v any, fn func(w io.Writer) TypeEncoder) (err error) {
	var f *os.File
	if f, err = os.Create(fp); err != nil {
		return
	}
	defer f.Close()
	return Encode(f, v, fn)
}

func Encode(w io.Writer, v any, fn func(w io.Writer) TypeEncoder) (err error) {
	enc := fn(w)
	err = enc.Encode(v)
	return
}

func DecodeFile[T any](fp string, fn func(r io.Reader) TypeDecoder) (v T, err error) {
	var f *os.File
	if f, err = os.Open(fp); err != nil {
		return
	}
	defer f.Close()
	return Decode[T](f, fn)
}

func Decode[T any](r io.Reader, fn func(r io.Reader) TypeDecoder) (v T, err error) {
	dec := fn(r)
	err = dec.Decode(&v)
	return
}
