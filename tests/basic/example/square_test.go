package math

import "testing"

func TestSquare(t *testing.T) {
	input := 3
	expect := 9
	output := Square(input)
	if output != expect {
		t.Errorf("Square Calculate Fail, input: %d, output: %d, expect: %d", input, output, expect)
	}
	// Logf only print in -v mode when case success
	t.Logf("input: %d, output: %d, expect: %d", input, output, expect)
}

func TestSquareWithFatal(t *testing.T) {
	input := 3
	expect := 9
	output := Square(input)
	if output != expect {
		// Fatalf will break function, if execute there, t.Logf will skip
		t.Fatalf("Square Calculate Fail, input: %d, output: %d, expect: %d", input, output, expect)
	}
	// Logf only print in -v mode when case success
	t.Logf("input: %d, output: %d, expect: %d", input, output, expect)
}