package codec

import (
	"bytes"
	"testing"
)

func TestRGB24ToYUV444_2x2(t *testing.T) {
	w, h := 2, 2
	// RGB24 packed buffer in row-major order:
	// (0,0), (1,0), (0,1), (1,1) — each pixel is 3 bytes: R,G,B.
	rgb := []byte{
		0, 0, 0, // top left
		255, 255, 255, // top right

		128, 128, 128, // bottom left
		0, 0, 0, // bottom right
	}

	wantY := []byte{0, 255, 128, 0}
	wantU := []byte{128, 128, 128, 128}
	wantV := []byte{128, 128, 128, 128}

	gotY, gotU, gotV, err := RGB24ToYUV444(rgb, w, h)
	if err != nil {
		t.Fatalf("RGB24ToYUV444: %v", err)
	}

	if len(gotY) != w*h {
		t.Errorf("The length of Y is wrong: %v", gotY)
	}

	if !bytes.Equal(gotY, wantY) {
		t.Errorf("Y mismatch: got %v want %v", len(gotU), w*h)
	}

	if len(gotU) != w*h {
		t.Errorf("The length of U is wrong: %v", gotU)
	}

	if !bytes.Equal(gotU, wantU) {
		t.Errorf("U mismatch: got %v want %v", len(gotU), w*h)
	}

	if len(gotV) != w*h {
		t.Errorf("The length of V is wrong: %v", gotV)
	}

	if !bytes.Equal(gotV, wantV) {
		t.Errorf("V mismatch: got %v want %v", len(gotV), w*h)
	}
}

func TestRGB24ToYUV444_InvalidLengthErrors(t *testing.T) {
	w, h := 2, 2
	rgb := make([]byte, w*h*3-1)

	_, _, _, err := RGB24ToYUV444(rgb, w, h)

	if err == nil {
		t.Error("Expected an error here")
	}
}
