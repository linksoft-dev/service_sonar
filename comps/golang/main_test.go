package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	// Check the response body is what we expect.
	expected := "I do stuff!"
	if doStuff() != expected {
		t.Errorf("doStuff(): got %v want %v",
			doStuff(), expected)
	}
	expected = "I do stuff!"
	if doStuff2() != expected {
		t.Errorf("doStuff(): got %v want %v",
			doStuff2(), expected)
	}

}
