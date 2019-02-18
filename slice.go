/*Write a program which prompts the user to enter integers
and stores the integers in a sorted slice. The program
should be written as a loop. Before entering the loop, the
program should create an empty integer slice of size
(length) 3. During each pass through the loop, the program
prompts the user to enter an integer to be added to the
slice. The program adds the integer to the slice, sorts
the slice, and prints the contents of the slice in sorted
order. The slice must grow in size to accommodate any
number of integers which the user decides to enter. The
program should only quit (exiting the loop) when the user
enters the character ‘X’ instead of an integer.*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Ready to make a sorted list of integers.")
	fmt.Println("Type X (upper case) to quit.")
	intsList := make([]int, 0, 3)
	var input string
	for {
		fmt.Println("Enter an integer:  ")
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			intInput, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error: not an integer")
				continue
			}
			intsList = append(intsList, intInput)
			sort.Ints(intsList)
			fmt.Println(intsList)
		}
	}
}

/*Activity: slice.go, author: manuelcaeiro, date: December/2018*/
