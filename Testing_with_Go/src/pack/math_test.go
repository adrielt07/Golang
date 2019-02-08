package pack

/*
 Tests package pack.math function

 Parallel:
 TestCanAddNumbers and TestCanSubtractNumbers run in parallel.
 Both tests sleeps for 1 second. That means it will take 2 seconds for the test to finish. (comment out t.Parallel to see)
 However, if ran in Parallel the test will finish in 1 sec
*/

// Imports
import (
	"fmt"
	"os"
	"testing"
	"time"
)

/*
Table driven test
tdt_table - a slice that contains an anonymouse struct with two fields:
in - slice of integers used for calculation
out - the expected output from in
*/
var tdt_table = []struct {
	function string
	in       []int
	out      int
}{
	{"add", []int{1, 2}, 3},
	{"add", []int{1, 2, 3, 4}, 10},
	{"subtract", []int{1, 2, 3}, -4},
	{"subtract", []int{5, 3}, 2},
}

var subtract_arr []int

/*
 TestMain - is similar to setup and teardown
 When the test runner finds a function with the name 'TestMain'
 it will give control of executing the test to this function

 As a result: it will be similar to setup and teardown
*/
func TestMain(m *testing.M) {
	// Run with -v flag to see the messages
	println("Setting up variables for Subtract")
	subtract_arr = []int{1, 2, 3}
	fmt.Printf("subtract_arr = %v\n", subtract_arr)

	result := m.Run()

	println("Tests are done")
	os.Exit(result)
}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()                // Run this test in Parallel with TestCanSubtractNumbers
	time.Sleep(1 * time.Second) // Sleep for 1 second

	if testing.Short() {
		t.Skip("Skipping long test")
	}

	// Used Table Driven Test
	for _, entry := range tdt_table {
		if entry.function == "add" {
			result := Add(entry.in...)
			if result != entry.out {
				t.Error("Failed to add numbers")
			}
		}
	}
}

// Test Subtract function
func TestCanSubtractNumbers(t *testing.T) {
	t.Parallel()                // Run this test in Parrallel with TestCanAddNumbers
	time.Sleep(1 * time.Second) // Sleep for 1 second

	// Use Table Driven Test
	for _, entry := range tdt_table {
		if entry.function == "subtract" {
			result := Subtract(entry.in...)
			if result != entry.out {
				t.Error("Failed to add numbers")
			}
		}
	}

	// Shows the usage of TestMain - changed the value of subtract_arr
	result := Subtract(subtract_arr...)
	if result != -4 {
		t.Error("Failed to Subtract Numbers")
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
