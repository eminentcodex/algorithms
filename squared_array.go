/*

  Write a function that takes in a non-empty array of integers that are sorted
  in ascending order and returns a new array of the same length with the squares
  of the original integers also sorted in ascending order.
  [1, 2, 3, 5, 6, 8, 9] => [1, 4, 9, 25, 36, 64, 81]
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20}
	out := SortedSquaredArray(nums)
	fmt.Println(out)
}

func SortedSquaredArray(array []int) []int {

	var output = make([]int, len(array))

	var start, end, counter int

	end = len(array) - 1

	counter = len(array) - 1

	for counter >=0 {
		if end == start {
			output[counter] = array[start] * array[start]
			break
		} else if math.Abs(float64(array[start])) > math.Abs(float64(array[end])) {
			output[counter] = array[start] * array[start]
			start++
			counter--
		} else if math.Abs(float64(array[end])) > math.Abs(float64(array[start])) {
			output[counter] = array[end] * array[end]
			end--
			counter--
		} else if math.Abs(float64(array[end])) == math.Abs(float64(array[start])) {
			output[counter] = array[end] * array[end]
			counter--
			output[counter] = array[end] * array[end]
			counter--
			start++
			end--
		}
	}

	return output
}

