package codec

import (
	"bytes"
	"testing"
)

func TestRGB24ToYUV420_2x2Block(t *testing.T) {
	rgb := []byte{
		255, 0, 0,
		0, 255, 0,
		0, 0, 255,
		255, 255, 255,
	}

	w, h := 2, 2

	wantY := []byte{76, 150, 29, 255}
	wantU := []byte{128}
	wantV := []byte{121}

	gotY, gotU, gotV, err := RGB24ToYUV420(rgb, w, h)
	if err != nil {
		t.Fatalf("RGB24ToYUV420: %v", err)
	}

	if len(gotY) != w*h {
		t.Errorf("The length of Y is wrong: %v", gotY)
	}

	if !bytes.Equal(gotY, wantY) {
		t.Errorf("Y mismatch: got %v want %v", gotY, wantY)
	}

	if len(gotU) != 1 {
		t.Errorf("The length of U is wrong: %v", gotU)
	}

	if !bytes.Equal(gotU, wantU) {
		t.Errorf("U mismatch: got %v want %v", gotU, wantU)
	}

	if len(gotV) != 1 {
		t.Errorf("The length of V is wrong: %v", gotV)
	}

	if !bytes.Equal(gotV, wantV) {
		t.Errorf("V mismatch: got %v want %v", gotV, wantV)
	}
}

func TestRGB24ToYUV420_4x4TwoBlocks(t *testing.T) {
	rgb := []byte{
		// y=0
		255, 0, 0, 0, 255, 0, 0, 255, 255, 255, 0, 255,
		// y=1
		0, 0, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0,
		// y=2
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// y=3
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	w, h := 4, 4

	gotY, gotU, gotV, err := RGB24ToYUV420(rgb, w, h)

	if err != nil {
		t.Fatalf("RGB24ToYUV420: %v", err)
	}

	if len(gotY) != w*h {
		t.Errorf("The length of Y is wrong: %v", gotY)
	}

	if len(gotU) != (w/2)*(h/2) {
		t.Errorf("The length of U is wrong: %v", gotU)
	}

	if len(gotV) != (w/2)*(h/2) {
		t.Errorf("The length of V is wrong: %v", gotV)
	}

	wantU := []byte{128, 128, 128, 128}
	wantV := []byte{121, 134, 128, 128}

	if !bytes.Equal(gotU, wantU) {
		t.Errorf("U mismatch: got %v want %v", gotU, wantU)
	}

	if !bytes.Equal(gotV, wantV) {
		t.Errorf("V mismatch: got %v want %v", gotV, wantV)
	}
}
