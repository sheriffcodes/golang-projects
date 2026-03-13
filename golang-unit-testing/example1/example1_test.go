package example1

import(
	"testing"
)

func TestCheckDivisibility(t *testing.T) {
	input := 5
	want := "Five!"
	got := CheckDivisibility(input)
	t.Logf("Running test for input:%d", input)
	if want != got {
		//t.Error("Incorrect Response")
		t.Errorf("Incorrect Response [input: %d] [want: %s] [got: %s]", input, want, got)
	}
}



