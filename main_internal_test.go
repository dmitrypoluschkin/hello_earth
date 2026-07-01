package main

// func Example_main() {
// 	main()
// 	// Output:
// 	// Hello World!
// }

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello World!"

	got := greet()

	if got != want {
		// mark this test is failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}