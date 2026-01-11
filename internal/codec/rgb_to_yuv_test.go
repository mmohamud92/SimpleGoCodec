package codec

import "testing"

func TestRGBToYUV_GreyHasNeutralChroma(t *testing.T) {
	tests := []struct {
		name    string
		r, g, b uint8
		wantY   uint8
		wantU   uint8
		wantV   uint8
	}{
		{
			name:  "black",
			r:     0,
			g:     0,
			b:     0,
			wantY: 0,
			wantU: 128,
			wantV: 128,
		},
		{
			name:  "white",
			r:     255,
			g:     255,
			b:     255,
			wantY: 255,
			wantU: 128,
			wantV: 128,
		},
		{
			name:  "grey",
			r:     128,
			g:     128,
			b:     128,
			wantY: 128,
			wantU: 128,
			wantV: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotY, gotU, gotV := RGBToYUV(tt.r, tt.g, tt.b)

			if gotY != tt.wantY {
				t.Errorf("Y: got %d want %d", gotY, tt.wantY)

			}

			if gotU != tt.wantU {
				t.Errorf("U: got %d want %d", gotU, tt.wantU)
			}

			if gotV != tt.wantV {
				t.Errorf("V: got %d want %d", gotV, tt.wantV)
			}
		})
	}
}

func TestRGBToYUV_PureRedMovesChrome(t *testing.T) {
	// Pure red: lots of "red-ness" and no green/blue.
	// In YUV/YCbCr-style spaces:
	//   - V (Cr, red-difference) should move ABOVE 128 (more red).
	//   - U (Cb, blue-difference) should move BELOW 128 (not blue).
	// 128 is the neutral chroma centre because we add +128 when encoding U/V into uint8.
	y, u, v := RGBToYUV(255, 0, 0)

	if v <= 128 {
		t.Errorf("expected V > 128 for pure red, got %d", v)
	}
	if u >= 128 {
		t.Errorf("expected U < 128 for pure red, got %d", u)
	}

	if y == 0 || y == 255 {
		t.Errorf("expected Y not extreme for pure red, got %d", y)
	}
}
