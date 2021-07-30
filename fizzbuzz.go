package main

import "strconv"

/*
	This functions takes 5 parameters and returns a list of strings with numbers from 1 to limit, where:
	- all multiples of int1 are replaced by str1
	- all multiples of int2 are replaced by str2
	- all multiples of int1 and int2 are replaced by str1str2
 */
func FizzBuzz(int1, int2, limit int, str1, str2 string) []string {
	results := make([]string, 0)

	// We iterate on every number until limit, starting from 1.
	for i := 1; i <= limit; i++ {
		switch {
		// If number is multiple of int1 and int2, we concatenate str1 and str2 and append the result to the list.
		case i % int1 == 0 && i % int2 == 0:
			results = append(results, str1+str2)
		// If number is multiple of int1, we append str1 to the list.
		case i % int1 == 0:
			results = append(results, str1)
		// If number is multiple of int2, we append str2 to the list.
		case i % int2 == 0:
			results = append(results, str2)
		// If number not a multiple of int1 or int2, we just append number to the list.
		default:
			results = append(results, strconv.FormatInt(int64(i), 10))
		}
	}

	return results
}
