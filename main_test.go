package main

import "testing"

func TestHello(t *testing.T) {
	hello := Hello()
	const expectedHello = "Hello world!"

	if hello != expectedHello {
		t.Errorf(`Expected: "%s"; actual: "%s"`, expectedHello, hello)
	}
}

