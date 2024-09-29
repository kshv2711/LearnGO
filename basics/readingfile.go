package main

import (
	"bufio"
	"fmt"
	_ "go/scanner"
	"os"
)
func main(){
	var f *os.File
	f=os.Stdin
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(">",scanner.Text())
	}
}
/*First, there is a call to bufio.NewScanner() using standard input (os.Stdin) as its
parameter. This call returns a bufio.Scanner variable, which is used with the Scan()
function for reading from os.Stdin line by line.*/