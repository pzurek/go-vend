package tools

import "testing"

func Test_PadRight_1(t *testing.T) {
	const s, l, p, out = "10", 0, "0", "10"
	x := PadRight(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadRight_2(t *testing.T) {
	const s, l, p, out = "10", 2, "0", "10"
	x := PadRight(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadRight_3(t *testing.T) {
	const s, l, p, out = "ra", 3, "a", "raa"
	x := PadRight(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadRight_4(t *testing.T) {
	const s, l, p, out = "Test", 9, "x", "Testxxxxx"
	x := PadRight(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadRight_5(t *testing.T) {
	const s, l, p, out = "Test", 21, "xoxox", "Testxoxoxxoxoxxoxoxxo"
	x := PadRight(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}
