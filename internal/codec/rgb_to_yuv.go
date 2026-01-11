package codec

import "math"

func RGBToYUV(r, g, b uint8) (y, u, v uint8) {
	// Convert 8-bit channel values into float so we can apply the conversion matrix.
	// (The coefficients are fractional, so float math is the clearest starting point.)
	rf := float64(r)
	gf := float64(g)
	bf := float64(b)

	// BT.601-ish (full-range) RGB -> Y'CbCr-style transform.
	//
	// Y' (luma/brightness) is a weighted sum of R/G/B:
	//   - green contributes most to perceived brightness, then red, then blue.
	//   - coefficients sum to ~1, so greys (R=G=B) keep the same level (e.g. 128 -> ~128).
	yf := 0.299*rf + 0.587*gf + 0.114*bf

	// U and V are chroma ("colour difference") channels. They can be negative/positive, but we store
	// them in uint8 [0..255], so we add +128 to centre "neutral chroma" at 128.
	//
	// For greys (R=G=B), the difference terms cancel and U≈128, V≈128.
	uf := -0.14713*rf - 0.28886*gf + 0.436*bf + 128
	vf := 0.615*rf - 0.51499*gf - 0.10001*bf + 128

	// Round to nearest and clamp into [0..255] so we can safely return uint8s.
	return roundClampToByte(yf), roundClampToByte(uf), roundClampToByte(vf)
}

func roundClampToByte(val float64) uint8 {
	n := min(max(int(math.Round(val)), 0), 255)
	return uint8(n)
}
