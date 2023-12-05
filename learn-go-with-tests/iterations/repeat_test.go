package iteration

import "testing"
import "fmt"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(6, "a")
	}
}

func ExampleRepeat() {
	repeated := Repeat(3, "0")
	fmt.Println(repeated)
	// Output: 000
}

func TestRepeat(t *testing.T) {
	repeated := Repeat(10, "a")
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
