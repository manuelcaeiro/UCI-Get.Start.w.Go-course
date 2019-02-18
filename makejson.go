/*Write a program which prompts the user to first enter a
name, and then enter an address. Your program should create
a map and add the name and address to the map using the
keys “name” and “address”, respectively. Your program
should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dir := make(map[string]string)
	var name string
	var addr string
	sreader := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter a name: ")
	if sreader.Scan() {
		name = sreader.Text()
	}
	dir["name"] = name

	fmt.Printf("Enter an address: ")
	if sreader.Scan() {
		addr = sreader.Text()
	}
	dir["address"] = addr

	dirlist, _ := json.Marshal(dir)
	fmt.Println(string(dirlist))
}

/*Activity: makejson.go, author: manuelcaeiro, date: January/2019*/
