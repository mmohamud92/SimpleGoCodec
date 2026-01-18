package codec

import (
	"reflect"
	"testing"
)

func TestRLEEncode_Basic(t *testing.T) {
	data := []byte{5, 5, 5, 2, 2, 9}

	got := RLEEncode(data)
	want := []Run{{5, 3}, {2, 2}, {9, 1}}

	assertRunsEqual(t, got, want)
}

func TestRLEEncode_Empty(t *testing.T) {
	var data []byte

	got := RLEEncode(data)
	want := []Run{}
	assertRunsEqual(t, got, want)
}

func TestRLEEncode_Single(t *testing.T) {
	data := []byte{7}

	got := RLEEncode(data)
	want := []Run{{7, 1}}

	assertRunsEqual(t, got, want)
}

func TestRLEEncode_AllSame(t *testing.T) {
	data := []byte{5, 5, 5, 5, 5, 5}

	got := RLEEncode(data)
	want := []Run{{5, 6}}

	assertRunsEqual(t, got, want)
}

func assertRunsEqual(t *testing.T, got, want []Run) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
