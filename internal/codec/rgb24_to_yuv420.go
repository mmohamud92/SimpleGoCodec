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

	for by := range ch {
		for bx := range cw {
			px := bx * 2
			py := by * 2

			sumU := 0
			sumV := 0

			// pixel (px, py)
			base := (py*w + px) * 3
			Y, U, V := RGBToYUV(rgb[base], rgb[base+1], rgb[base+2])
			y[py*w+px] = Y
			sumU += int(U)
			sumV += int(V)

			// pixel (px + 1, py)
			base = (py*w + px + 1) * 3
			Y, U, V = RGBToYUV(rgb[base], rgb[base+1], rgb[base+2])
			y[py*w+px+1] = Y
			sumU += int(U)
			sumV += int(V)

			// pixel (px, py + 1)
			base = ((py+1)*w + px) * 3
			Y, U, V = RGBToYUV(rgb[base], rgb[base+1], rgb[base+2])
			y[(py+1)*w+px] = Y
			sumU += int(U)
			sumV += int(V)

			// pixel (px + 1, py + 1)
			base = ((py+1)*w + px + 1) * 3
			Y, U, V = RGBToYUV(rgb[base], rgb[base+1], rgb[base+2])
			y[(py+1)*w+px+1] = Y
			sumU += int(U)
			sumV += int(V)

			u[by*cw+bx] = byte(roundClampToByte(float64(sumU) / 4.0))
			v[by*cw+bx] = byte(roundClampToByte(float64(sumV) / 4.0))
		}
	}

	return y, u, v, nil
}
