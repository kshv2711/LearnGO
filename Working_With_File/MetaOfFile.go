package main

import (
	"fmt"
	"log"
	"os"
)
func main(){
	FileStat,err:=os.Stat("Test")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("FileName:",FileStat.Name())
	fmt.Println("Size:",FileStat.Size())
	fmt.Println("Permissions:",FileStat.Mode())
	fmt.Println("Last Modified:",FileStat.ModTime())
	fmt.Println("Is Diredctory:",FileStat.IsDir())
	
	
}