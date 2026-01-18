package codec

type Run struct {
	Value byte
	Count int
}

func RLEEncode(data []byte) []Run {
	runs := make([]Run, 0, len(data))

	for i := 0; i < len(data); i++ {
		count := 1

		for i+1 < len(data) && data[i] == data[i+1] {
			count++
			i++
		}

		runs = append(runs, Run{data[i], count})
	}

	return runs
}
