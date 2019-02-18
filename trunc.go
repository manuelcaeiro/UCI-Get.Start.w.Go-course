/*Write a program which prompts the user to enter a floating
point number and prints the integer which is a truncated
version of the floating point number that was entered.
Truncation is the process of removing the digits to the
right of the decimal place.*/

package main

import "fmt"

func main() {
	var fltNum float64

	fmt.Println("Enter a float number: ")
	fmt.Scan(&fltNum)

	intNum := int(fltNum)
	fmt.Println("The whole part of that number is", intNum)
}

/*Activity: trunc.go, author: manuelcaeiro, date: December/2018*/
