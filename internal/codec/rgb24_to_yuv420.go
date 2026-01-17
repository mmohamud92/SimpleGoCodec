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

	return nil, nil, nil, nil
}
