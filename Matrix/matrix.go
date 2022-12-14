package matrix

import (
	"errors"
	"fmt"
)

/*
Method: Display
Parameters: One 2D float slice
Return: void
How it works: This method's purpose is to automate the tedious task of writing
two nested for loops every time you want to print out a 2D slice. I find myself
doing this a lot while debugging, so I wrote a simple method to make it faster.
*/
func Display(array [][]float64) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

/*
Method: CheckDims
Parameters: Two 2D float slices
Return: a boolean
How it works: checks if two matricies can be multiplied together by finding if the number of
columns in the first matrix are equal to the number of rows in the second.
*/
func CheckDims(array1 [][]float64, array2 [][]float64) bool {
	return (len(array1[0]) == len(array2))
}

/*
Method: Dot
Parameters: Two 2D float slices
Return: a 2D float slice and an error
How it works: This method starts by checking if the two matrcies have the correct
dimensions to multiply. I am currently creating a dummy matrix for the return to work.
I wish there was some way where I didn't have to return the dummy array.
Since the dimensions not being able to multiply is an error, I am using the error package
to send the user an error.

Next, if the matricies can multiple we create a third array of the correct dimensions which we will return
Then we take the dot product of the two matricies assigning the answers to the corresponding
indexes in our result array. Lastly we return the matrix and a nil for the error.
*/
func Dot(array1 [][]float64, array2 [][]float64) ([][]float64, error) {
	if !CheckDims(array1, array2) {
		array3 := make([][]float64, len(array1))
		for i := range array3 {
			array3[i] = make([]float64, len(array2[0]))
		}
		return array3, errors.New("these dimensions didn't work")
	}
	array3 := make([][]float64, len(array1))
	for i := range array3 {
		array3[i] = make([]float64, len(array2[0]))
	}
	var total float64 = 0
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2[0]); j++ {
			for k := 0; k < len(array2); k++ {
				total += array1[i][k] * array2[k][j]
			}
			array3[i][j] = total
			total = 0

		}
	}
	return array3, nil
}

func Transpose(array1 [][]float64) [][]float64 {
	array2 := make([][]float64, len(array1[0]))
	for i := range array2 {
		array2[i] = make([]float64, len(array1))
	}
	for i := 0; i < len(array1[0]); i++ {
		for j := 0; j < len(array1); j++ {
			array2[i][j] = array1[j][i]
		}
	}
	return array2
}

func Sum(array [][]float64) float64 {
	var sum float64 = 0.0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			sum += array[i][j]
		}
	}
	return sum
}
