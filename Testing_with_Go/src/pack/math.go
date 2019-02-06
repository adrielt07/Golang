package pack

/*
Use as an example for basic testing.
Basic testing should be 'Ok'

Code coverage is not 100% on purpose.
*/

/*
Add - adds two integers
sum - Returns a single integers (sum)
*/
func Add(nums ...int) (sum int) {

	// Simple example to test code coverage
	if len(nums) == 0 {
		println("No arguments provided")
		return 0
	}

	for _, i := range nums {
		sum += i
	}
	return
}

func Subtract(nums ...int) (difference int) {
	for idx, i := range nums {
		if idx == 0 {
			difference = i
		} else {
			difference -= i
		}
	}
	return
}
