package main

import (
	"fmt"
	"strconv"
)

func permute(digits string) []string {
	if len(digits) == 1 {
		return []string{digits}
	}
	perms := []string{}
	for i, digit := range digits {
		rest := digits[:i] + digits[i+1:]
		for _, perm := range permute(rest) {
			perms = append(perms, string(digit)+perm)
		}
	}
	return perms
}

func generateNumbers(number int) ([]int, error) {
	strNumber := strconv.Itoa(number)
	perms := permute(strNumber)
	result := []int{}
	for _, perm := range perms {
		parsedNum, err := strconv.Atoi(perm)
		if err != nil {
			return nil, err
		}
		result = append(result, parsedNum)
	}
	return result, nil
}

func main() {
	number := 531
	results, err := generateNumbers(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All possible numbers with the same length as", number)
	fmt.Println(results)
}

/*
Write a function that takes a positive integer and returns the next smaller positive
integer containing the same digits.

For example:

nextSmaller(21) == 12
nextSmaller(531) == 513
nextSmaller(2071) == 2017
Return -1 (for Haskell: return Nothing, for Rust: return None), when there is no smaller
number that contains the same digits. Also return -1 when the next smaller number with
the same digits would require the leading digit to be zero.

nextSmaller(9) == -1
nextSmaller(111) == -1
nextSmaller(135) == -1
nextSmaller(1027) == -1 // 0721 is out since we don't write numbers with leading zeros
some tests will include very large numbers.
test data only employs positive integers.
The function you write for this challenge is the inverse of this kata: "Next bigger number
with the same digits."
*/
