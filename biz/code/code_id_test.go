package code

import "testing"

func TestEncode(t *testing.T) {
	t.Skip()

	var max = 10000000
	var m = map[string]bool{}
	for i := 1; i < max; i++ {
		code := Encode(i)
		if len(code) != CODELENGTH {
			t.Error("len(code) error", i)
			break
		}
		m[code] = true
	}
	if len(m) != max-1 {
		t.Error("len(m) error")
	}
}

func TestDecode(t *testing.T) {
	t.Skip()

	var max = 10000000
	for i := 1; i < max; i++ {
		if Decode(Encode(i)) != i {
			t.Error("len(m) error", i)
			break
		}
	}
}
