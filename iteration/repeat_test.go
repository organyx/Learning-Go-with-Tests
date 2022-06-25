package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := RepeatWithCount("b", 3)
	expected := "bbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
