package binarydist

import (
	"github.com/klauspost/compress/zstd"
	"io"
)

func newZstdWriter(w io.Writer) (io.WriteCloser, error) {
	zstdWriter, err := zstd.NewWriter(w)
	if err != nil {
		return nil, err
	}
	return &zstdWriteCloser{Encoder: zstdWriter}, nil
}

type zstdWriteCloser struct {
	*zstd.Encoder
}

func (z *zstdWriteCloser) Close() error {
	return z.Encoder.Close()
}
