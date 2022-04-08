/*
Given two non-empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.

A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same order as they appear in the array. For instance, the numbers [1, 3, 4] form a subsequence of the array [1, 2, 3, 4], and so do the numbers [2, 4]. Note that a single number in an array and the array itself are both valid subsequences of the array.
*/

package main

func IsValidSubsequence(array []int, sequence []int) bool {
	look := 0
	reference := 0
	sLen := len(sequence)
	
	if len(array) == 0 || len(sequence) == 0 {
		return false
	}
	
	if len(array) < len(sequence) {
		return false
	}
	
	for reference < len(sequence) && look < len(array) {
		if sequence[reference] == array[look] {
			sLen--
			reference++
			look++
		} else {
			look++
		}
	}
	
	if sLen == 0 {
		return true
	}
	
	return false
}
