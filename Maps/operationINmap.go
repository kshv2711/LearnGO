package main

import (
	"fmt"
	"strings"
)
func main()  {
	myMap:=map[int]string{
		1:"keshav",
		2:"aastha",
		3:"sunil",
		4:"anil",
	}
	fmt.Println(myMap)
	fmt.Println(strings.Repeat("*",20))
	// add key and value in  map
	myMap[5]="Ankita"
	fmt.Println(myMap)
	fmt.Println(strings.Repeat("*",20))
	// delete key vaue in map
	delete(myMap,4)
	fmt.Println(myMap)
	fmt.Println(strings.Repeat("*",20))
	// update key value in map
	myMap[3]="pranav"
	fmt.Println(myMap)
}