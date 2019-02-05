package pack

/*
Tests package pack.math function
*/

// Import testing
import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		// Prints the failed test
		t.Log("Failed to add one and two")
		// Signals that test have failed
		t.Fail()
	}
}
