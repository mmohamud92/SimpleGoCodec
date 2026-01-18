package codec

import "errors"

func RGB24ToYUV420(rgb []byte, w, h int) (y, u, v []byte, err error) {
	if w <= 0 {
		return nil, nil, nil, errors.New("width should be more than 0")
	}

	if h <= 0 {
		return nil, nil, nil, errors.New("height should be more than 0")
	}

	if len(rgb) != w*h*3 {
		return nil, nil, nil, errors.New("the length of rgb must be width * height * 3")
	}

	if w%2 != 0 || h%2 != 0 {
		return nil, nil, nil, errors.New(" the value of w or h is not even")
	}

	cw := w / 2
	ch := h / 2

	y = make([]byte, w*h)
	u = make([]byte, cw*ch)
	v = make([]byte, cw*ch)

	offsets := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}

	for by := range ch {
		for bx := range cw {
			px := bx * 2
			py := by * 2

			sumU := 0
			sumV := 0

			for _, off := range offsets {
				x := px + off[0]
				ypix := py + off[1]

				base := (ypix*w + x) * 3
				Y, U, V := RGBToYUV(rgb[base], rgb[base+1], rgb[base+2])

				y[ypix*w+x] = Y
				sumU += int(U)
				sumV += int(V)
			}

			u[by*cw+bx] = byte(roundClampToByte(float64(sumU) / 4.0))
			v[by*cw+bx] = byte(roundClampToByte(float64(sumV) / 4.0))
		}
	}

	return y, u, v, nil
}
