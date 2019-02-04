package examplepkg

/*
Example of creating my own package and importing it.
Function that will be use outside of this code: go up one level and
check file 'my_addition.go'

Few things to take note of:

Uppercase - function names starts with an uppercase letter
is public and can be used when imported

Lowercase - is private and cannot be used when imported

Example: MyAdd function below is public and can be used if I import my_add package
*/

/*
MyAdd - takes an indefinite number of ints and add them all
nums - list of integers
sum - result after adding all nums
*/
func MyAdd(nums ...int) (sum int) {
	for i := range nums {
		sum += i
	}
	return
}
