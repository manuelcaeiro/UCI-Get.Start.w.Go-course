/* Write a program which reads information from a file and
represents it in a slice of structs. Assume that there is a
text file which contains a series of names. Each line of
the text file has a first name and a last name, in that
order, separated by a single space on the line.
Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name. Each
field will be a string of size 20 (characters).
Your program should prompt the user for the name of the
text file. Your program will successively read each line of
the text file and create a struct which contains the first
and last names found in the file. Each struct created will
be added to a slice, and after all lines have been read
from the file, your program will have a slice containing
one struct for each line in the file. After reading all
lines from the file, your program should iterate through
your slice of structs and print the first and last names
found in each struct. */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FixLenStr(length int, str string) string {
	// align left...
	strSize := fmt.Sprintf("%%-%d.%ds", length, length)
	// or align right
	//strSize := fmt.Sprintf("%%-%d.%ds", length, length)
	return fmt.Sprintf(strSize, str)
}

type fullNames struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Printf("Enter the name of the file to read: ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	structs := make([]fullNames, 0)
	read := bufio.NewScanner(file)
	for read.Scan() {
		line := fullNames{}
		line.fname = strings.Split(read.Text(), " ")[0]
		line.lname = strings.Split(read.Text(), " ")[1]
		structs = append(structs, line)
	}
	//fmt.Println(structs)
	for _, item := range structs {
		fmt.Println(FixLenStr(20, item.fname), FixLenStr(20, item.lname))
		//fmt.Printf("%-20s%-20s\n", item.fname, item.lname)
	}

}

/*Activity: read.go, author: manuelcaeiro, date: January/2019*/
