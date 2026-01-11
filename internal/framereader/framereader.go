package framereader

import (
	"errors"
	"fmt"
	"io"
)

type FrameReader struct {
	r         io.Reader
	frameSize int
}

var ErrTruncatedFrame = errors.New("truncated frame")

var ErrInvalidDimensions = errors.New("invalid dimensions")

func NewFrameReader(r io.Reader, w, h int) (*FrameReader, error) {
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("%w: %dx%d", ErrInvalidDimensions, w, h)
	}

	fr := &FrameReader{
		r:         r,
		frameSize: w * h * 3,
	}

	return fr, nil
}

func (fr *FrameReader) Next() ([]byte, error) {
	buf := make([]byte, fr.frameSize)
	_, err := io.ReadFull(fr.r, buf)
	if err == nil {
		return buf, nil
	}

	if errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, ErrTruncatedFrame
	}

	return nil, err
}
