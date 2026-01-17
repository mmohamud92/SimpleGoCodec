package codec

import "errors"

func RGB24ToYUV444(rgb []byte, w, h int) (y, u, v []byte, err error) {
	if w <= 0 {
		return nil, nil, nil, errors.New("width should be more than 0")
	}

	if h <= 0 {
		return nil, nil, nil, errors.New("height should be more than 0")
	}

	if len(rgb) != w*h*3 {
		return nil, nil, nil, errors.New("the length of rgb must be width * height * 3")
	}

	nPixels := w * h
	y = make([]byte, nPixels)
	u = make([]byte, nPixels)
	v = make([]byte, nPixels)

	for i := range nPixels {
		base := i * 3
		r := rgb[base]
		g := rgb[base+1]
		b := rgb[base+2]

		Y, U, V := RGBToYUV(r, g, b)

		y[i] = Y
		u[i] = U
		v[i] = V
	}

	return y, u, v, nil
}
