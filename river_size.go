/*
You're given a two-dimensional array (a matrix) of potentially unequal height and width containing only 0s and 1s. Each 0 represents land, and each 1 represents part of a river. A river consists of any number of 1s that are either horizontally or vertically adjacent (but not diagonally adjacent). The number of adjacent 1s forming a river determine its size.

Note that a river can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line; it can be L-shaped, for example.

Write a function that returns an array of the sizes of all rivers represented in the input matrix. The sizes don't need to be in any particular order.
*/
package main

func RiverSizes(matrix [][]int) []int {
	riverLenght := []int{}
	
	for row:= 0; row < len(matrix); row++{
		colLength := len(matrix[row])
		for col:=0; col<colLength;col++ {
			if matrix[row][col] == 0 {
				continue
			}
			l := explore(row, col, matrix)
			if l != 0 {
				riverLenght = append(riverLenght, l)	
			}
		}
	}
	
	return riverLenght
}

func explore(i int, j int, matrix [][]int) int {
	
	if i >= len(matrix) || i < 0 {
		return 0
	} else if j >= len(matrix[i]) || j < 0 {
		return 0
	}
	
	if matrix[i][j] == 0 {
		matrix[i][j]=0
		return 0
	}
	
	matrix[i][j]=0 // mark as visited
	
	top := explore(i-1, j, matrix)
	bottom := explore(i+1, j, matrix)
	left := explore(i, j-1, matrix)
	right:= explore(i, j+1, matrix)
	
	return 1 + top + bottom + left + right
}
