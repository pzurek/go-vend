package tools

import "testing"

func Test_PadLeft_1(t *testing.T) {
	const s, l, p, out = "10", 0, "0", "010"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadLeft_2(t *testing.T) {
	const s, l, p, out = "10", 2, "0", "10"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadLeft_3(t *testing.T) {
	const s, l, p, out = "ra", 3, "a", "ara"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadLeft_4(t *testing.T) {
	const s, l, p, out = "Test", 9, "x", "xxxxxTest"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadLeft_5(t *testing.T) {
	const s, l, p, out = "Test", 21, "xoxox", "xoxoxxoxoxxoxoxxoTest"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}

func Test_PadLeft_6(t *testing.T) {
	const s, l, p, out = "4", 5, "0", "00004"
	x := PadLeft(s, p, l)
	if x != out {
		t.Errorf("PadLeft(%v, %v, %v) = %v, want %v", s, l, p, x, out)
	}
}
