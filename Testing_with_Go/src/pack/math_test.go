package pack

/*
 Tests package pack.math function

 Parallel:
 TestCanAddNumbers and TestCanSubtractNumbers run in parallel.
 Both tests sleeps for 1 second. That means it will take 2 seconds for the test to finish. (comment out t.Parallel to see)
 However, if ran in Parallel the test will finish in 1 sec
*/

// Import testing
import (
	"testing"
	"time"
)

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()                // Run this test in Parallel with TestCanSubtractNumbers
	time.Sleep(1 * time.Second) // Sleep for 1 second

	if testing.Short() {
		t.Skip("Skipping long test")
	}

	result := Add(1, 2)
	if result != 3 {
		// Prints the failed test
		t.Log("Failed to add one and two")
		// Signals that test have failed
		t.Fail()
	}

	result = Add(1, 2, 3, 4)
	if result != 10 {
		t.Error("Failed to add more than two numebrs")
	}
}

// Test Subtract function
func TestCanSubtractNumbers(t *testing.T) {
	t.Parallel()                // Run this test in Parrallel with TestCanAddNumbers
	time.Sleep(1 * time.Second) // Sleep for 1 second
	result := Subtract(10, 5)
	if result != 5 {
		t.Error("Failed to subtract two numbers")
	}

	result = Subtract(10, 5, 5)
	if result != 0 {
		t.Error("Failed to subtract more than 2 numbers")
	}
}

// This is an example of how to skip a certain test that may not be ready yet
func TestNumbersCanFly(t *testing.T) {
	t.Skip("Example of how to skip in testing...")
}

func TestCanMultiplyNumbers(t *testing.T) {
	// Running test with -v flag will skip the test (Odd test)
	if testing.Verbose() {
		t.Skip("Not implemented yet")
	}
}
