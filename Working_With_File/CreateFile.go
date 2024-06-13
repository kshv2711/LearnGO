package main

import (
	
	"log"
	"os"
)
func main(){
	// var newfile *os.File
	// fmt.Printf("%T\n",newfile)
	// var err error
	newfile, err:=os.Create("MyFile.txt")
	if err!=nil{
		log.Fatal(err)
	}
	// err=os.Truncate("MyFile.txt",0)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	newfile.Close()
	
}