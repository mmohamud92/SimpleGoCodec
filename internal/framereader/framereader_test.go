package framereader

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

const bytesPerPixel = 3

func TestFrameReader_ReadsTwoFrames(t *testing.T) {
	// Arrange
	w, h := 2, 2
	frameSize := w * h * bytesPerPixel

	data := make([]byte, frameSize*2)
	for i := range data {
		data[i] = byte(i)
	}

	r := bytes.NewReader(data)

	fr, err := NewFrameReader(r, w, h)
	if err != nil {
		t.Fatalf("NewFrameReader: %v", err)
	}

	f1, err1 := fr.Next()
	f2, err2 := fr.Next()
	_, err3 := fr.Next()

	if err1 != nil {
		t.Fatalf("The error: %v", err1)
	}

	if err2 != nil {
		t.Fatalf("The error: %v", err2)
	}

	if len(f1) != frameSize {
		t.Errorf("len(f1) is %d, should be %d", len(f1), frameSize)
	}

	if !bytes.Equal(f1, data[:frameSize]) {
		t.Errorf("frame1 bytes mismatch: expected data[:%d]", frameSize)
	}

	if len(f2) != frameSize {
		t.Errorf("len(f2) is %d, should be %d", len(f2), frameSize)
	}

	if !bytes.Equal(f2, data[frameSize:]) {
		t.Errorf("frame2 bytes mismatch: expected data[%d:]", frameSize)
	}

	if !errors.Is(err3, io.EOF) {
		t.Errorf("This should've been the end of the file. Error: %v", err3)
	}
}

func TestFrameReader_EmptyInputReturnsEOF(t *testing.T) {
	r := bytes.NewReader(nil)

	fr, err := NewFrameReader(r, 2, 2)

	if err != nil {
		t.Fatalf("NewFrameReader: %v", err)
	}

	fr1, err1 := fr.Next()

	if fr1 != nil {
		t.Error("Expected nil frame on EOF")
	}

	if !errors.Is(err1, io.EOF) {
		t.Errorf("Got a %v error, should be an EOF error", err1)
	}
}

func TestFrameReader_InvalidDimensionsErrors(t *testing.T) {
	w, h := 2, 0
	r := bytes.NewReader(nil)

	_, err := NewFrameReader(r, w, h)

	if err == nil {
		t.Errorf("There should be an error of type %v, got nil", ErrInvalidDimensions)
	}
	if !errors.Is(err, ErrInvalidDimensions) {
		t.Errorf("Error should be %v", ErrInvalidDimensions)
	}
}

func TestFrameReader_TruncatedFrameErrorsCorrectly(t *testing.T) {
	w, h := 2, 2
	frameSize := w * h * bytesPerPixel
	data := make([]byte, frameSize*2-1)

	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}

	r := bytes.NewReader(data)

	fr, err := NewFrameReader(r, w, h)
	if err != nil {
		t.Fatalf("NewFrameReader: %v", err)
	}

	f1, err1 := fr.Next()

	if len(f1) != frameSize {
		t.Errorf("len(f1) is %d, should be %d", len(f1), frameSize)
	}

	if !bytes.Equal(f1, data[:frameSize]) {
		t.Errorf("frame1 bytes mismatch: expected data[:%d]", frameSize)
	}

	if err1 != nil {
		t.Errorf("There shouldn't be an error here: %v", err1)
	}

	_, err2 := fr.Next()

	if !errors.Is(err2, ErrTruncatedFrame) {
		t.Errorf("This should be an error of type %v", ErrTruncatedFrame)
	}

	if errors.Is(err2, io.ErrUnexpectedEOF) {
		t.Error("expected truncation got EOF")
	}
}
